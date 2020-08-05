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

type RequestTiming struct {
	Dns       time.Duration
	Conn      time.Duration
	Tls       time.Duration
	FirstByte time.Duration
	Total     time.Duration
}

type Response struct {
	StatusCode int
	Proto      string
	Body       []byte
	Headers    http.Header
	Timing     RequestTiming
}

func DoRequest(request Request, execCtx *ExecutionContext) (response Response) {
	start := time.Now()
	var dnsTime, connectionTime, tlsHandshakeTime, firstByteTime, totalTime time.Duration

	req, _ := http.NewRequest(strings.ToUpper(request.Method), request.Url, bytes.NewReader([]byte(request.Body)))

	ctx := *execCtx
	req.Header.Set("user-agent", Sprintf("RestBeast-%s", ctx.Version))
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	if ctx.Debug {
		log.Printf("method: %s", request.Method)
		log.Printf("url: %s", request.Url)

		for k, v := range req.Header {
			log.Printf("header: %s=%s", k, v)
		}

		log.Printf("body: %s", request.Body)
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
		log.Fatal(err)
	}

	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)

	totalTime = time.Since(start)
	timing := RequestTiming{
		Dns:       dnsTime,
		Conn:      connectionTime,
		Tls:       tlsHandshakeTime,
		FirstByte: firstByteTime,
		Total:     totalTime,
	}

	return Response{
		StatusCode: res.StatusCode,
		Proto:      res.Proto,
		Body:       data,
		Headers:    res.Header,
		Timing:     timing,
	}
}
