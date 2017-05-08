package govcert

import (
	"fmt"
	"regexp"
)

type CSRReq struct {
	CommonName         string
	OrganizationalUnit []string
	OrganizationName   string
	Country            string
	State              string
	Locality           string
	KeyPassword        string
	SanDNS             []string
	SanEmail           []string
	SanIP              []string
}

func (c *CSRReq) Request() (*request, error) {
	params := make(map[string]RequestField)
	if c.CommonName == "" {
		return nil, fmt.Errorf("CommonName must be specified")
	}
	params["cn"] = RequestField{c.CommonName}
	params["ou"] = RequestField(c.OrganizationalUnit)
	params["o"] = RequestField{c.OrganizationName}
	params["c"] = RequestField{c.Country}
	params["st"] = RequestField{c.State}
	params["l"] = RequestField{c.Locality}
	params["san-email"] = RequestField(c.SanEmail)
	params["san-dns"] = RequestField(c.SanDNS)
	params["san-ip"] = RequestField(c.SanIP)
	params["key-password"] = RequestField{c.KeyPassword}
	req := &request{
		Action:  "gencsr",
		Mparams: params,
	}
	req.Params(params, "no-prompt")
	return req, nil
}

func (CSRReq) RequiresAuth() bool {
	return false
}

type CSRResp struct {
	PrivateKey string
	CSR        string
}

func (r response) ParseCSR() (*CSRResp, error) {

	re := regexp.MustCompile("(?s)-----BEGIN RSA PRIVATE KEY-----.*-----END RSA PRIVATE KEY-----")
	requestMatches := re.FindStringSubmatch(r.stdOut.String())
	if len(requestMatches) < 1 {
		return nil, fmt.Errorf("PrivateKEY not found")
	}
	re = regexp.MustCompile("(?s)-----BEGIN CERTIFICATE REQUEST-----.*-----END CERTIFICATE REQUEST-----")
	csrMatches := re.FindStringSubmatch(r.stdOut.String())
	if len(csrMatches) < 1 {
		return nil, fmt.Errorf("CSR not found")
	}
	return &CSRResp{
		PrivateKey: requestMatches[0],
		CSR:        csrMatches[0],
	}, nil
}
