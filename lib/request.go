package lib

import (
	"bytes"
	"crypto/tls"
	. "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"strings"
	"time"
)

func (request *Request) Exec() error {
	start := time.Now()
	var dnsTime, connectionTime, tlsHandshakeTime, firstByteTime, totalTime time.Duration

	httpReq, err := http.NewRequest(strings.ToUpper(request.Method), request.Url, bytes.NewReader([]byte(request.Body)))
	if err != nil {
		return Errorf("unable to construct request, %s\n", err)
	}

	httpReq.Header.Set("user-agent", Sprintf("RestBeast-%s", request.ExecutionContext.Version))
	for key, value := range request.Headers {
		httpReq.Header.Set(key, value)
	}

	if request.ExecutionContext.Debug {
		log.Printf("request method: %s", request.Method)
		log.Printf("request url: %s", request.Url)

		for k, v := range httpReq.Header {
			log.Printf("request header: %s=%s", k, v)
		}

		log.Printf("request body: %s", request.Body)
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

	if request.ExecutionContext.Debug {
		log.Printf("response status: %d %s", res.StatusCode, res.Status)
		for k, v := range res.Header {
			log.Printf("response header: %s=%s", k, v)
		}

		log.Printf("response body: %s", data)
	}

	totalTime = time.Since(start)
	timing := RequestTiming{
		Dns:       dnsTime,
		Conn:      connectionTime,
		Tls:       tlsHandshakeTime,
		FirstByte: firstByteTime,
		Total:     totalTime,
	}

	request.Response = &Response{
		Method:     request.Method,
		Url:        request.Url,
		StatusCode: res.StatusCode,
		Proto:      res.Proto,
		Body:       data,
		Headers:    res.Header,
		Timing:     timing,
		Request:    request,
	}

	return nil
}
