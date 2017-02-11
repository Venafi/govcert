package govcert

import (
	"bytes"
	"fmt"
	"regexp"
)

type response struct {
	errOut *bytes.Buffer
	stdOut *bytes.Buffer
}

type Response interface {
	Body() (string, error)
	RequestID() (string, error)
}

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
