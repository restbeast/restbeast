package lib

import (
	"bytes"
	"crypto/tls"
	"errors"
	. "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"strings"
	"time"
)

type RequestTiming struct {
	Dns       time.Duration
	Conn      time.Duration
	Tls       time.Duration
	FirstByte time.Duration
	Total     time.Duration
}

type Response struct {
	Method     string
	Url        string
	StatusCode int
	Proto      string
	Body       []byte
	Headers    http.Header
	Timing     RequestTiming
	Request    *Request
}

func DoRequest(request Request, execCtx *ExecutionContext) (*Response, error) {
	start := time.Now()
	var dnsTime, connectionTime, tlsHandshakeTime, firstByteTime, totalTime time.Duration

	req, err := http.NewRequest(strings.ToUpper(request.Method), request.Url, bytes.NewReader([]byte(request.Body)))
	if err != nil {
		return nil, errors.New(Sprintf("unable to construct request, %s\n", err))
	}

	ctx := *execCtx
	req.Header.Set("user-agent", Sprintf("RestBeast-%s", ctx.Version))
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	if ctx.Debug {
		log.Printf("request method: %s", request.Method)
		log.Printf("request url: %s", request.Url)

		for k, v := range req.Header {
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

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if ctx.Debug {
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

	return &Response{
		Method:     request.Method,
		Url:        request.Url,
		StatusCode: res.StatusCode,
		Proto:      res.Proto,
		Body:       data,
		Headers:    res.Header,
		Timing:     timing,
		Request:    &request,
	}, nil
}
