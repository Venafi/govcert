package govcert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
)

type response struct {
	errOut  *bytes.Buffer
	stdOut  *bytes.Buffer
	apiResp *http.Response
}

type Response interface {
	Body() (string, error)
	RequestID() (string, error)
	JSONBody() (map[string]interface{}, error)
	Unmarshal(interface{}) error
	// Location() (*url.URL, error)
}

func ResponseFromAPI(resp *http.Response) *response {
	// spew.Dump(resp.Body)
	r := &response{
		apiResp: resp,
		stdOut:  new(bytes.Buffer),
	}
	body, _ := ioutil.ReadAll(resp.Body)
	r.stdOut.Read(body)
	spew.Dump(r)
	return r
}

func NewResponse() *response {
	return &response{
		stdOut: new(bytes.Buffer),
		errOut: new(bytes.Buffer),
	}
}

func (r *response) Location() (*url.URL, error) {
	return nil, nil
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

func (r *response) JSONBody() (j map[string]interface{}, err error) {
	// Clean output
	// myStr := `^[.]+`
	// testByte := []byte(myStr)
	// spew.Dump(testByte)
	regBytes := []byte{0x5e, 0x5b, 0x2e, 0x1a, 0x0a, 0x5d, 0x2b}
	re, err := regexp.Compile(string(regBytes))
	if err != nil {
		return
	}
	fmt.Println("--- JSON BODY ---")
	spew.Dump(r.stdOut.Bytes())
	out := re.ReplaceAll(r.stdOut.Bytes(), []byte{})
	fmt.Println("--- PURE OUT ---")
	spew.Dump(out)
	err = json.Unmarshal(out, &j)
	return
}

func (r *response) Unmarshal(d interface{}) error {
	fmt.Printf("---- RAW BODY ----\n%s\n---- RAW BODY END ----\n", spew.Sdump(r.apiResp.Body))
	body, _ := ioutil.ReadAll(r.apiResp.Body)
	fmt.Printf("---- JSON Body ----\n%s\n", body)
	err := json.Unmarshal(body, d)
	return err
}
