// Package service provides functionality related to HTTP transport and logging.
package service

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"cypt/internal/dddcore"
	"cypt/internal/logger/entity"
	"cypt/internal/logger/entity/events"
)

// HTTPBodyDecoder represents a function signature for decoding a byte slice into a string, with the possibility of returning an error.
type HTTPBodyDecoder func([]byte) (string, error)

// LogHTTPTransport is a custom HTTP transport that logs the HTTP requests and responses.
type LogHTTPTransport struct {
	core     http.RoundTripper
	eventBus dddcore.EventBus
	decoder  HTTPBodyDecoder
}

// RoundTrip executes a single HTTP transaction and logs the request and response information.
func (t *LogHTTPTransport) RoundTrip(req *http.Request) (res *http.Response, err error) {
	var reqBody, resBody []byte

	log := entity.HTTPRequestLog{
		At:        time.Now(),
		Method:    req.Method,
		Origin:    req.URL.String(),
		Host:      req.Host,
		ReqHeader: req.Header,
	}

	if reqBody, req.Body, err = drainBody(req.Body); err == nil {
		log.ReqBody = t.decodeBody(reqBody)
	}

	res, err = t.core.RoundTrip(req)
	if err != nil {
		log.Error = err
		return nil, err
	}

	if resBody, res.Body, err = drainBody(res.Body); err == nil {
		log.ResBody = t.decodeBody(resBody)
	}

	log.StatusCode = res.StatusCode
	log.Latency = time.Since(log.At)
	log.ResHeader = res.Header

	ev := events.NewHTTPRequestDoneEvent(&log)
	_ = t.eventBus.Post(ev)

	return res, nil
}

// NewLogHTTPTransport creates a new LogHTTPTransport instance with the specified event bus.
func NewLogHTTPTransport(eb dddcore.EventBus) http.RoundTripper {
	t := NewDefaultHTTPTransport()
	logTransport := &LogHTTPTransport{core: t, eventBus: eb}

	return logTransport
}

// NewLogHTTPTransportWithDecoder creates a new LogHTTPTransport instance with the specified event bus and body decoder.
func NewLogHTTPTransportWithDecoder(eb dddcore.EventBus, decoder HTTPBodyDecoder) http.RoundTripper {
	t := NewDefaultHTTPTransport()
	logTransport := &LogHTTPTransport{core: t, eventBus: eb, decoder: decoder}

	return logTransport
}

// NewDefaultHTTPTransport creates a new HTTP transport with custom settings.
func NewDefaultHTTPTransport() *http.Transport {
	transport := &http.Transport{
		Proxy: NoProxyAllowed,
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 60 * time.Second,
		}).Dial,
		DisableKeepAlives:     true,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   60 * time.Second,
		ResponseHeaderTimeout: 60 * time.Second,
	}

	return transport
}

// NoProxyAllowed is a proxy function that disallows any proxy.
func NoProxyAllowed(request *http.Request) (*url.URL, error) {
	return nil, nil
}

// NewHTTPClient creates a new HTTP client with the specified event bus.
func NewHTTPClient(eb dddcore.EventBus) *http.Client {
	tripper := NewLogHTTPTransport(eb)
	client := &http.Client{
		Transport: tripper,
		Timeout:   time.Second * 120,
	}

	return client
}

// NewHTTPClientWithDecoder creates a new HTTP client with the specified event bus and body decoder.
func NewHTTPClientWithDecoder(eb dddcore.EventBus, decoder HTTPBodyDecoder) *http.Client {
	tripper := NewLogHTTPTransportWithDecoder(eb, decoder)
	client := &http.Client{
		Transport: tripper,
		Timeout:   time.Second * 120,
	}

	return client
}

// drainBody reads and drains the body, returning the body content and a new ReadCloser.
// @see net/http/httputil/dump.go
func drainBody(b io.ReadCloser) ([]byte, io.ReadCloser, error) {
	if b == nil || b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return nil, http.NoBody, nil
	}

	var buf bytes.Buffer

	if _, err := buf.ReadFrom(b); err != nil {
		return nil, b, err
	}

	if err := b.Close(); err != nil {
		return nil, b, err
	}

	bufBytes := buf.Bytes()
	return bufBytes, io.NopCloser(bytes.NewReader(bufBytes)), nil
}

// decodeBody decodes the given byte slice using the decoder function of the LogHTTPTransport.
// If no decoder is set, it returns the byte slice as a string.
// If decoding encounters an error, it constructs an error message with the original byte slice and returns it.
func (t *LogHTTPTransport) decodeBody(bufBytes []byte) string {
	if t.decoder == nil {
		return string(bufBytes)
	}

	body, err := t.decoder(bufBytes)

	if err != nil {
		var builder strings.Builder
		builder.WriteString("Decode Error:")
		builder.WriteString(err.Error())
		builder.WriteString("\n")
		builder.WriteString("Origin:")
		builder.WriteString(string(bufBytes))

		body = builder.String()
	}

	return body
}
