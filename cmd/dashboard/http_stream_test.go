package main

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/http2"
)

func TestHTTP2StreamOutlivesInitialHeaderDeadline(t *testing.T) {
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor != 2 {
			t.Errorf("stream used HTTP/%d", r.ProtoMajor)
		}
		_, _ = io.WriteString(w, "connected\n")
		w.(http.Flusher).Flush()
		select {
		case <-time.After(200 * time.Millisecond):
			_, _ = io.WriteString(w, "still connected\n")
		case <-r.Context().Done():
		}
	}))
	server.Config.ReadHeaderTimeout = 30 * time.Millisecond
	server.Config.ConnState = clearHeaderReadDeadline
	server.Config.Protocols = new(http.Protocols)
	server.Config.Protocols.SetHTTP1(true)
	server.Config.Protocols.SetUnencryptedHTTP2(true)
	server.Start()
	defer server.Close()
	transport := &http2.Transport{AllowHTTP: true, DialTLSContext: func(ctx context.Context, network, address string, _ *tls.Config) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, network, address)
	}}
	defer transport.CloseIdleConnections()
	client := &http.Client{Transport: transport, Timeout: 3 * time.Second}
	response, err := client.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil || string(body) != "connected\nstill connected\n" {
		t.Fatalf("HTTP/2 stream expired after initial headers: body=%q err=%v", body, err)
	}
	if !strings.Contains(response.Proto, "2") {
		t.Fatal("request did not negotiate HTTP/2")
	}
}
