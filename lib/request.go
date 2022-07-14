package lib

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	. "fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

func (request *Request) SetUrl(urlToSet string) {
	var encodedParams string
	if request.Params != nil {
		params := url.Values{}

		for k, v := range request.Params {
			params.Add(k, v)
		}

		encodedParams = params.Encode()
	}

	fullUrl := urlToSet
	if encodedParams != "" {
		if strings.Contains(fullUrl, "?") {
			fullUrl += "&" + encodedParams
		} else {
			fullUrl += "?" + encodedParams
		}
	}

	request.Url = urlToSet
	request.FullUrl = fullUrl
}

func (request *Request) SetMethod(method string) {
	if method != "" {
		request.Method = strings.ToUpper(method)
	} else {
		request.Method = "GET"
	}
}

func (request *Request) Exec() error {
	start := time.Now()
	var dnsTime, connectionTime, tlsHandshakeTime, firstByteTime, totalTime time.Duration

	httpReq, err := http.NewRequest(strings.ToUpper(request.Method), request.FullUrl, request.Body)
	if err != nil {
		return Errorf("unable to construct request, %s\n", err)
	}

	vRemovedVersion := strings.Replace(request.ExecutionContext.Version, "v", "", 1)
	httpReq.Header.Set("user-agent", Sprintf("RestBeast/%s", vRemovedVersion))
	request.Headers.ToRequest(httpReq)

	if request.ExecutionContext.Debug {
		log.Printf("request method: %s", request.Method)
		log.Printf("request url: %s", request.FullUrl)

		for k, v := range httpReq.Header {
			log.Printf("request header: %s=%s", k, v)
		}

		if jsonMarshalled, err := json.MarshalIndent(request.Body, "", "  "); err == nil {
			log.Printf("request body: %s", string(jsonMarshalled))
		}
	}

	trace := &httptrace.ClientTrace{
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			dnsTime = time.Since(start)
		},

		ConnectDone: func(network, addr string, err error) {
			connectionTime = time.Since(start)
		},

		TLSHandshakeDone: func(tls.ConnectionState, error) {
			tlsHandshakeTime = time.Since(start)
		},

		GotFirstResponseByte: func() {
			firstByteTime = time.Since(start)
		},
	}

	httpReq = httpReq.WithContext(httptrace.WithClientTrace(httpReq.Context(), trace))
	res, err := request.RoundTripper.RoundTrip(httpReq)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	totalTime = time.Since(start)
	timing := RequestTiming{
		Dns:       dnsTime,
		Conn:      connectionTime,
		Tls:       tlsHandshakeTime,
		FirstByte: firstByteTime,
		Total:     totalTime,
	}

	buf := &bytes.Buffer{}
	var bodySize int64
	if request.Body != nil {
		bodySize, _ = io.Copy(buf, request.Body)
	}

	headers := &Headers{}
	headers.FromResponse(res.Header)

	if request.ExecutionContext.Debug {
		log.Printf("response status: %d %s", res.StatusCode, res.Status)
		log.Printf("response DNS Resolve time: %d ms", timing.Dns.Milliseconds())
		log.Printf("response Connection time: %d ms", timing.Conn.Milliseconds())
		if timing.Tls > 0 {
			log.Printf("response TLS Handshake time: %d ms", timing.Tls.Milliseconds())
		}
		log.Printf("response Firstbyte time: %d ms", timing.FirstByte.Milliseconds())
		log.Printf("response Total time: %d ms", timing.Total.Milliseconds())
		if bodySize > 0 {
			log.Printf("response Bytes Sent: %s", humanize.Bytes(uint64(bodySize)))
		}
		log.Printf("response Bytes Received: %s", humanize.Bytes(uint64(len(data))))

		headers.OrderedCallBack(
			func(k, v string) {
				log.Printf("response header: %s=%s", k, v)
			},
		)

		if jsonMarshalled, err := json.MarshalIndent(data, "", "  "); err == nil {
			log.Printf("response body: %s", string(jsonMarshalled))
		}
	}

	request.Response = &Response{
		Method:        request.Method,
		Url:           request.Url,
		StatusCode:    res.StatusCode,
		Proto:         res.Proto,
		Body:          data,
		Headers:       headers,
		Timing:        timing,
		Request:       request,
		BytesSend:     uint64(bodySize),
		BytesReceived: uint64(len(data)),
	}

	return nil
}
