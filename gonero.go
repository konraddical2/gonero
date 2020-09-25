package gonero

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gabstv/httpdigest"
)

// RPCConfig holds the configuration of a monero RPC client.
type RPCConfig struct {
	Protocol      string
	Host          string
	Port          uint
	CustomHeaders map[string]string
	Transport     http.RoundTripper
}

// Address returns a formatted address string for the RPC server
func (cfg RPCConfig) Address() string {
	p := cfg.Protocol
	if p == "" {
		p = "http"
	}
	return fmt.Sprintf("%s://%s:%d", p, cfg.Host, cfg.Port)
}

// NewRPCConfig creates a new RPCConfig struct
// with a Transport for authentication if needed
func NewRPCConfig(protocol, host string, port uint, username, password, caFile string) (*RPCConfig, error) {

	cfg := &RPCConfig{
		Protocol: protocol,
		Host:     host,
		Port:     port,
	}
	// Add TLS configuration
	if cfg.Protocol == "https" {
		t, err := tlsTransport(caFile)
		if err != nil {
			return nil, err
		}
		cfg.Transport = t
	}
	// Use http digest auth for RPC username:password
	if username != "" {
		t := httpdigest.New(username, password)
		// chain previous transport
		t.Transport = cfg.Transport
		// set updated transport
		cfg.Transport = t
	}
	return cfg, nil
}

func tlsTransport(caFile string) (*http.Transport, error) {
	var tlsConfig *tls.Config

	if caFile != "" {
		// Load CA cert
		caCert, err := ioutil.ReadFile(caFile)
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		// Setup HTTPS client
		tlsConfig = &tls.Config{
			RootCAs: caCertPool,
		}
		tlsConfig.BuildNameToCertificate()
	} else {
		tlsConfig = &tls.Config{InsecureSkipVerify: true}

	}
	return &http.Transport{TLSClientConfig: tlsConfig}, nil
}
