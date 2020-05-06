package climacell

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testingHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewTLSServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return cli, s.Close
}

func testingAssertAuthValid(t *testing.T, r *http.Request) {
	assert.Equal(t, "abc123", r.URL.Query().Get("apikey"))
}

func testingMockResponseHandler(t *testing.T, filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		testingAssertAuthValid(t, r)

		data, _ := ioutil.ReadFile(filename)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
		w.Write([]byte(data))
	}
}

func TestClientGetEndpoint(t *testing.T) {
	cli := NewClient("abc123")

	e := cli.getEndpoint("v3", "weather/realtime")
	assert.Equal(t, "https://api.climacell.co/v3/weather/realtime", e)

	e = cli.getEndpoint("v0", "weather/realtime")
	assert.Equal(t, "https://api.climacell.co/v3/weather/realtime", e)
}

func TestClientNewGetRequest(t *testing.T) {
	cli := NewClient("abc123")
	req := &RealtimeRequest{}
	r, err := cli.newGetRequest("v3", "weather/realtime", req)

	assert.Nil(t, err)
	assert.Equal(t, "GET", r.Method)
	assert.Equal(t, "application/json", r.Header.Get("Accept"))
	assert.Equal(t, UserAgentString, r.Header.Get("User-Agent"))
}

func TestClientNewPostRequest(t *testing.T) {
	cli := NewClient("abc123")
	r, err := cli.newPostRequest("v3", "weather/realtime", []byte("{\"foo\": \"bar\"}"))

	assert.Nil(t, err)
	assert.Equal(t, "POST", r.Method)
	assert.Equal(t, "application/json", r.Header.Get("Accept"))
	assert.Equal(t, UserAgentString, r.Header.Get("User-Agent"))
}

func TestClientSetAuth(t *testing.T) {
	cli := NewClient("abc123")
	req := &RealtimeRequest{}
	r, err := cli.newGetRequest("v3", "weather/realtime", req)
	cli.setAuth(r)

	assert.Nil(t, err)
	assert.Equal(t, "GET", r.Method)
	assert.Equal(t, "application/json", r.Header.Get("Accept"))
	assert.Equal(t, UserAgentString, r.Header.Get("User-Agent"))
	assert.Equal(t, "abc123", r.URL.Query().Get("apikey"))
	assert.Equal(t, "0", r.URL.Query().Get("lat"))
	assert.Equal(t, "0", r.URL.Query().Get("lon"))
}
