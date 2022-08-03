package http

import (
	"fmt"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockNoopHandler struct{}

func (mockNoopHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func freePort(t *testing.T) int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	require.Nil(t, err)

	l, err := net.ListenTCP("tcp", addr)
	require.Nil(t, err)
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}

func Test_AnHTTPServerCanBeStarted(t *testing.T) {
	for _, test := range []struct {
		name      string
		address   string
		handler   http.Handler
		errString string
	}{
		{
			name:      "missing port",
			address:   "-1",
			errString: "listen tcp: address -1: missing port in address",
		},
		{
			name:      "invalid port",
			address:   ":-1",
			errString: "listen tcp: address -1: invalid port",
		},
		{
			name:    "valid; happy path",
			address: fmt.Sprintf("localhost:%d", freePort(t)),
			handler: mockNoopHandler{},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			server := NewServer(test.address, test.handler)
			err := server.Start()

			if test.errString != "" {
				require.NotNil(t, err)
				require.Equal(t, test.errString, err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func Test_ARunningHTTPServerCanBeStopped(t *testing.T) {
	server := NewServer(fmt.Sprintf("localhost:%d", freePort(t)), &mockNoopHandler{})
	require.Nil(t, server.Start())
	require.Nil(t, server.Stop())
}

func Test_AnHTTPServerReturnsAUsefulErrorWhenItsAddressIsInUse(t *testing.T) {
	address := fmt.Sprintf("localhost:%d", freePort(t))

	s0 := NewServer(address, &mockNoopHandler{})
	require.Nil(t, s0.Start())
	defer s0.Stop()

	s1 := NewServer(address, &mockNoopHandler{})
	err := s1.Start()
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "address already in use")
}
