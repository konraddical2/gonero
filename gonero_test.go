package gonero

import (
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRPCConfigAddress(t *testing.T) {
	type test struct {
		input    RPCConfig
		expected string
	}

	tests := []test{
		{input: RPCConfig{Protocol: "http", Host: "localhost", Port: 18081}, expected: "http://localhost:18081"},
		{input: RPCConfig{Protocol: "https", Host: "127.0.0.1", Port: 38081}, expected: "https://127.0.0.1:38081"},
		{input: RPCConfig{Protocol: "", Host: "localhost", Port: 18081}, expected: "http://localhost:18081"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expected, tc.input.Address(), "Address() formats properly")
	}
}

func TestNewRPCConfig(t *testing.T) {
	type test struct {
		protocol        string
		host            string
		port            uint
		username        string
		password        string
		caFile          string
		expectedAddress string
	}

	tests := []test{
		// unauthenticated with no ssl
		{"", "localhost", 18081, "", "", "", "http://localhost:18081"},
		{"http", "localhost", 18081, "", "", "", "http://localhost:18081"},
		{"http", "127.0.0.1", 38081, "", "", "", "http://127.0.0.1:38081"},
		// unauthenticated with ssl
		{"https", "localhost", 18081, "", "", "", "https://localhost:18081"},
		{"https", "127.0.0.1", 38081, "", "", "", "https://127.0.0.1:38081"},
		// authenticated with usename, no ssl
		{"", "localhost", 18081, "foo", "", "", "http://localhost:18081"},
		{"http", "localhost", 18081, "bar", "", "", "http://localhost:18081"},
		{"http", "127.0.0.1", 38081, "qux", "", "", "http://127.0.0.1:38081"},
		// authenticated with usename, ssl
		{"https", "localhost", 18081, "bar", "", "", "https://localhost:18081"},
		{"https", "127.0.0.1", 38081, "qux", "", "", "https://127.0.0.1:38081"},
		// authenticated with usename and password, no ssl
		{"", "localhost", 18081, "foo", "password", "", "http://localhost:18081"},
		{"http", "localhost", 18081, "bar", "password", "", "http://localhost:18081"},
		{"http", "127.0.0.1", 38081, "qux", "password", "", "http://127.0.0.1:38081"},
		// authenticated with usename and password, ssl
		{"https", "localhost", 18081, "bar", "password", "", "https://localhost:18081"},
		{"https", "127.0.0.1", 38081, "qux", "password", "", "https://127.0.0.1:38081"},
		// authenticated with usename and password, ssl with cert
		{"https", "localhost", 18081, "bar", "password", "tests/cert.tls", "https://localhost:18081"},
		{"https", "127.0.0.1", 38081, "qux", "password", "tests/cert.tls", "https://127.0.0.1:38081"},
	}

	for _, tc := range tests {
		cfg, _ := NewRPCConfig(tc.protocol, tc.host, tc.port, tc.username, tc.password, tc.caFile)
		assert.Equal(t, tc.expectedAddress, cfg.Address())
		//if cfg.Protocol == "https" {
		// TODO test ssl Transport
		//}
		//if tc.username != "" {
		// TODO test digest auth
		//}
	}

	failTests := []test{
		// authenticated with usename and password, ssl with incorrect path to cert
		{"https", "localhost", 18081, "bar", "password", "Foobar.tls", "https://localhost:18081"},
		{"https", "127.0.0.1", 38081, "qux", "password", "blah_bal.tls", "https://127.0.0.1:38081"},
	}

	for _, tc := range failTests {
		_, err := NewRPCConfig(tc.protocol, tc.host, tc.port, tc.username, tc.password, tc.caFile)
		assert.Error(t, err)
		assert.Equal(t, &os.PathError{Op: "open", Path: tc.caFile, Err: syscall.Errno(2)}, err)
	}
}
