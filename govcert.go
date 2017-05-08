package govcert

import (
	"fmt"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
        "os"
        "io"
	"path/filepath"
)

var vcertCmd *exec.Cmd

type Client interface {
	Do(Requestor) (Response, error)
	APIKey() string
	URL() string
	Username() string
	Password() string
}

// Client represents the base command and input/output handling
type client struct {
	cmd     *exec.Cmd
	cmdPath string
	apiKey  string
	url	string
	tppUser	string
	tppPass	string
}

// NewClient returns a client that wraps the temporary VCert binary
func NewClient(path, apikey string, url string) *client {
	return &client{
		cmd:    exec.Command(path),
		apiKey: apikey,
		url:	url,
	}
}

func NewClientTPP(path, username string, password string, url string) *client {
	return &client{
		cmd:    exec.Command(path),
		tppUser: username,
		tppPass: password,
		url: 	url,
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

func (c *client) Do(r Requestor) (Response, error) {
	cmd := *c.cmd
	resp := NewResponse()
	req, err := r.Request()
	if err != nil {
		return nil, err
	}
	if !req.hasAction() {
		return nil, fmt.Errorf("No action called")
	}
	if r.RequiresAuth() {
		if c.hasAPIKey() {
			req.params = append(req.params, "-k", c.apiKey)
			if c.hasURL() {
				if !strings.HasSuffix(c.url, "/") {
					c.url += "/"  // work around vcert bug requiring url to end with a slash character
				}
				req.params = append(req.params, "-venafi-saas-url", c.url)
			}
		} else if c.hasUsername() && c.hasPassword() {
			if !c.hasURL() {
				return nil, fmt.Errorf("Missing required WebSDK URL")
			}
			req.params = append(req.params, "-tpp-url", c.url)
			req.params = append(req.params, "-tpp-user", c.tppUser)
			req.params = append(req.params, "-tpp-password", c.tppPass)
		} else {
			return nil, fmt.Errorf("Insufficient auth parameters")
		}
	}
	// if !inSlice(req.params, "-k") {
	// 	req.params = append(req.params, "-k", c.APIKey)
	// }
	// if req.hasAPIKey() && !inSlice(req.params, "-k") {
	// 	req.params = append(req.params, "-k", req.apiKey)
	// }
	cmd.Stdout = resp.stdOut
	cmd.Stderr = resp.errOut

	cmd.Args = append(cmd.Args, req.Action)
	cmd.Args = append(cmd.Args, req.params...)

        s := strings.Join(cmd.Args," ")
        // If the debug flag file is present, write command args
        // to named file in debug flag file
	var fname string
        user,_ := user.Current()
        flagfile := filepath.Join(user.HomeDir,"debug.flag")
        freader,_ := os.Open(flagfile)
        fmt.Fscan(freader,&fname)
	fmt.Printf(fname)
        WriteStringToFile(fname, s)
	err = cmd.Run()

	if err != nil {
		body,_ := resp.Body()
		return nil, fmt.Errorf(err.Error() + "\n" + body)
	}
	return resp, err
}

// func (c *client) Retryable(req *Request, waittime, maxwait time.Duration) (Response, error) {
// 	return nil, nil
// }

func (c *client) parse(out []byte) string {
	if re, err := regexp.Compile("^[ .]+"); err == nil {
		out = re.ReplaceAll(out, []byte{})
	}
	return strings.Replace(string(out), c.cmd.Path, "", -1)
}

// fun

func (c *client) APIKey() string {
	return c.apiKey
}

func (c *client) hasAPIKey() bool {
	return len(c.apiKey) > 0
}

func (c *client) URL() string {
	return c.url
}

func (c *client) hasURL() bool {
	return len(c.url) > 0
}

func (c *client) Username() string {
	return c.tppUser
}

func (c *client) hasUsername() bool {
	return len(c.tppUser) > 0
}

func (c *client) Password() string {
	return c.tppPass
}

func (c *client) hasPassword() bool {
	return len(c.tppPass) > 0
}

// Copied from https://siongui.github.io/2016/04/05/go-write-string-to-file/

func WriteStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}

	return nil
}
