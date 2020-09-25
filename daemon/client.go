package daemon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/rpc/v2/json2"
	"github.com/konraddical2/gonero"
)

// New returns a new monerod daemon RPC client.
func New(cfg *gonero.RPCConfig) Client {
	cl := &client{
		addr:    cfg.Address(),
		headers: cfg.CustomHeaders,
		httpcl:  http.DefaultClient,
	}
	if cfg.Transport == nil {
		cl.httpcl = http.DefaultClient
	} else {
		cl.httpcl = &http.Client{
			Transport: cfg.Transport,
		}
	}
	return cl
}

type client struct {
	httpcl  *http.Client
	addr    string
	headers map[string]string
}

// Helper function for JSON RPC Methods
func (c *client) do(method string, in, out interface{}) error {
	payload, err := json2.EncodeClientRequest(method, in)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, c.addr+"/json_rpc", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	if c.headers != nil {
		for k, v := range c.headers {
			req.Header.Set(k, v)
		}
	}
	resp, err := c.httpcl.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// TODO preprocess body?

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	// in theory this is only done to catch
	// any monero related errors if
	// we are not expecting any data back
	if out == nil {
		v := &json2.EmptyResponse{}
		return GetRPCError(json2.DecodeClientResponse(resp.Body, v))

	}

	err = json2.DecodeClientResponse(resp.Body, out)
	if jsonErr, ok := err.(*json.SyntaxError); ok {
		// TODO handle binary data
		fmt.Printf("%s\n", body)
		fmt.Printf("%s\n", jsonErr)
		err = fmt.Errorf("Binary data in response")
	}
	return GetRPCError(err)
}

// Helper function for Other RPC Methods
func (c *client) doSlash(method string, in, out interface{}) error {
	payload, err := json.Marshal(in)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, c.addr+method, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	if c.headers != nil {
		for k, v := range c.headers {
			req.Header.Set(k, v)
		}
	}
	resp, err := c.httpcl.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status %v", resp.StatusCode)
	}
	defer resp.Body.Close()

	if out == nil {
		return nil
		//v := &json2.EmptyResponse{}
		//return json2.DecodeClientResponse(resp.Body, v)
	}
	return json.NewDecoder(resp.Body).Decode(&out)
}
