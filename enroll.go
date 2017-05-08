package govcert

import (
	"fmt"
	"strconv"
	"strings"
)

type EnrollReq struct {
	CommonName  string
	KeyCurve    string
	Sans        SAN
	Format      string
	Zone        string
	KeySize     int
	KeyType     string
	Chain       string
	KeyPassword string
	//APIKey      string
}

type SAN struct {
	DNS   []string
	IP    []string
	Email []string
}

func (e *EnrollReq) Request() (*request, error) {
	params := make(map[string]RequestField)
	params["format"] = RequestField{"json"}
	if e.CommonName == "" {
		return nil, fmt.Errorf("Common name must be specified")
	}
	if e.Zone == "" {
		e.Zone = "Default"
	}
	params["z"] = RequestField{e.Zone}
	if strings.EqualFold(e.KeyType, "ecdsa") {
		params["key-type"] = RequestField{"ecdsa"}
		params["key-curve"] = RequestField{e.KeyCurve}
	} else {
		params["key-type"] = RequestField{"rsa"}
		params["key-size"] = RequestField{strconv.Itoa(e.KeySize)}
	}
	params["cn"] = RequestField{e.CommonName}
	if e.Chain != "" {
		params["chain"] = RequestField{e.Chain}
	}
	if e.KeyPassword != "" {
		params["key-password"] = RequestField{e.KeyPassword}
	}

	params["san-dns"] = RequestField(e.Sans.DNS)
	params["san-email"] = RequestField(e.Sans.Email)
	params["san-ip"] = RequestField(e.Sans.IP)
	req := &request{
		Action: "enroll",
	}
	req.Params(params, "no-pickup", "no-prompt")

	return req, nil
}

func (EnrollReq) RequiresAuth() bool {
	return true
}
