package govcert

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"
)

var vcertCmd *exec.Cmd

// Client represents the base command and input/output handling
type Client struct {
	cmd       *exec.Cmd
	output    *bytes.Buffer
	errOutput *bytes.Buffer
}

// NewClient returns a client that wraps the temporary VCert binary
func NewClient(path string) *Client {
	return &Client{
		cmd:       exec.Command(path),
		output:    new(bytes.Buffer),
		errOutput: new(bytes.Buffer),
	}
}

// NewAuthorisedRequest prepares requests that require an api key
func (c *Client) NewAuthorisedRequest(apikey string) *Request {
	return &Request{
		client: c,
		apiKey: apikey,
	}
}

// NewRequest prepares requests that don't require an api key such as
// registration or returning help
func (c *Client) NewRequest() *Request {
	return &Request{
		client: c,
	}
}

func (c *Client) parse(out []byte) string {
	if re, err := regexp.Compile("^[ .]+"); err == nil {
		out = re.ReplaceAll(out, []byte{})
	}
	return strings.Replace(string(out), c.cmd.Path, "", -1)
}

func (c *Client) parseOutput() string {
	return c.parse(c.output.Bytes())
}

func (c *Client) parseError() string {
	return c.parse(c.errOutput.Bytes())
}

func (c *Client) httpError() HTTPError {
	return newHTTPError(c.errOutput.Bytes())
}
