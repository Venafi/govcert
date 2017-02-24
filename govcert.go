package govcert

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var vcertCmd *exec.Cmd

type Client interface {
	Do(*Request) (Response, error)
	Retryable(*Request, time.Duration, time.Duration) (Response, error)
	CSR(*CSRReq) (*CSRResp, error)
}

// Client represents the base command and input/output handling
type client struct {
	cmd     *exec.Cmd
	cmdPath string
	APIKey  string
	// uuid    uuid.UUID
	// output    *bytes.Buffer
	// errOutput *bytes.Buffer
}

// NewClient returns a client that wraps the temporary VCert binary
func NewClient(path, apikey string) *client {
	return &client{
		cmd:     exec.Command(path),
		cmdPath: path,
		APIKey:  apikey,
		// output:    new(bytes.Buffer),
		// errOutput: new(bytes.Buffer),
	}
}

// func (c *client) GenUUID() {
// 	c.uuid = uuid.NewV4()
// }

// NewAuthorisedRequest prepares requests that require an api key
// func (c *Client) NewAuthorisedRequest(apikey string) *Request {
// 	c.cmd = exec.Command(c.cmdPath)
// 	return &Request{
// 		apiKey: apikey,
// 		params: []string{},
// 	}
// }

// NewRequest prepares requests that don't require an api key such as
// registration or returning help
// func (c *Client) NewRequest() *Request {
// 	c.cmd = exec.Command(c.cmdPath)
// 	return &Request{
// 		params: []string{},
// 	}
// }

func (c *client) Do(req *Request) (Response, error) {
	cmd := *c.cmd
	resp := NewResponse()
	if !req.hasAction() {
		return nil, fmt.Errorf("No action called")
	}
	// if req.hasAPIKey() && !inSlice(req.params, "-k") {
	// 	req.params = append(req.params, "-k", req.apiKey)
	// }
	cmd.Stdout = resp.stdOut
	cmd.Stderr = resp.errOut

	cmd.Args = append(cmd.Args, req.Action)
	cmd.Args = append(cmd.Args, req.params...)
	err := cmd.Run()
	return resp, err
}

func (c *client) Retryable(req *Request, waittime, maxwait time.Duration) (Response, error) {
	return nil, nil
}

func (c *client) parse(out []byte) string {
	if re, err := regexp.Compile("^[ .]+"); err == nil {
		out = re.ReplaceAll(out, []byte{})
	}
	return strings.Replace(string(out), c.cmd.Path, "", -1)
}

// fun

func (c *client) hasAPIKey() bool {
	return len(c.APIKey) > 0
}
