package govcert

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var vcertCmd *exec.Cmd

// Client represents the base command and input/output handling
type Client struct {
	cmd       *exec.Cmd
	cmdPath   string
	apiKey    string
	output    *bytes.Buffer
	errOutput *bytes.Buffer
}

// NewClient returns a client that wraps the temporary VCert binary
func NewClient(path, key string) *Client {
	return &Client{
		cmd:       exec.Command(path),
		cmdPath:   path,
		apiKey:    key,
		output:    new(bytes.Buffer),
		errOutput: new(bytes.Buffer),
	}
}

// NewAuthorisedRequest prepares requests that require an api key
func (c *Client) NewAuthorisedRequest(apikey string) *Request {
	c.cmd = exec.Command(c.cmdPath)
	return &Request{
		apiKey: apikey,
		params: []string{},
	}
}

// NewRequest prepares requests that don't require an api key such as
// registration or returning help
func (c *Client) NewRequest() *Request {
	c.cmd = exec.Command(c.cmdPath)
	return &Request{
		params: []string{},
	}
}

func (c *Client) Do(req *Request) (Response, error) {
	cmd := *c.cmd
	resp := NewResponse()
	if !req.hasAction() {
		return nil, fmt.Errorf("No action called")
	}
	if c.hasAPIKey() && !inSlice(req.params, "-k") {
		req.params = append(req.params, "-k", c.apiKey)
	}
	cmd.Stdout = resp.stdOut
	cmd.Stderr = resp.errOut

	cmd.Args = append(cmd.Args, req.Action)
	cmd.Args = append(cmd.Args, req.params...)
	err := cmd.Run()
	return resp, err
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

func (c *Client) hasAPIKey() bool {
	return len(c.apiKey) > 0
}
