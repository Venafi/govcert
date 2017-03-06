package govcert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
)

type response struct {
	errOut *bytes.Buffer
	stdOut *bytes.Buffer
	// apiResp *http.Response
}

type Response interface {
	Body() (string, error)
	RequestID() (string, error)
	JSONBody() (map[string]interface{}, error)
	// Unmarshal(interface{}) error
	Pending() bool
	CompletedID() (string, error)
	ParseCSR() (*CSRResp, error)
	Bytes() []byte
	// Location() (*url.URL, error)
}

// func ResponseFromAPI(resp *http.Response) *response {
// 	// spew.Dump(resp.Body)
// 	r := &response{
// 		apiResp: resp,
// 		stdOut:  new(bytes.Buffer),
// 	}
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	r.stdOut.Read(body)
// 	spew.Dump(r)
// 	return r
// }

func NewResponse() *response {
	return &response{
		stdOut: new(bytes.Buffer),
		errOut: new(bytes.Buffer),
	}
}

func (r *response) Body() (string, error) {
	if len(r.errOut.Bytes()) == 0 && len(r.stdOut.Bytes()) == 0 {
		return "", fmt.Errorf("Response body is empty")
	}
	return r.stdOut.String() + r.errOut.String(), nil
}

func (r *response) RequestID() (string, error) {
	// Certificate issuance pending, you may request the retrieval of the Certificate later using the Request ID: 262d3ff0-efa9-11e6-9be2-891dab33d0eb
	re := regexp.MustCompile("Certificate issuance pending.*Request ID: ([a-z0-9-]+)")
	requestMatches := re.FindStringSubmatch(r.errOut.String())
	if len(requestMatches) == 2 {
		return requestMatches[1], nil
	}
	return "", fmt.Errorf("No pending Request ID was found")
}

func (r *response) CompletedID() (string, error) {
	re := regexp.MustCompile("retrieved request for ([a-z0-9-]+)")
	requestMatches := re.FindStringSubmatch(r.errOut.String())
	if len(requestMatches) == 2 {
		return requestMatches[1], nil
	}
	return "", fmt.Errorf("No pending Request ID was found")
}

func (r *response) JSONBody() (j map[string]interface{}, err error) {
	// Clean output
	regBytes := []byte{0x5e, 0x5b, 0x2e, 0x1a, 0x0a, 0x5d, 0x2b}
	re, err := regexp.Compile(string(regBytes))
	if err != nil {
		return
	}
	out := re.ReplaceAll(r.stdOut.Bytes(), []byte{})
	err = json.Unmarshal(out, &j)
	return
}

// func (r *response) Unmarshal(d interface{}) error {
// 	body, _ := ioutil.ReadAll(r.apiResp.Body)
// 	err := json.Unmarshal(body, d)
// 	return err
// }

func (r *response) Pending() bool {
	re, _ := regexp.Compile("Certificate issuance pending")
	return re.Match(r.errOut.Bytes())
}

func (r *response) Bytes() []byte {
	return r.stdOut.Bytes()
}
