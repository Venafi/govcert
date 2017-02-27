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
}

type SAN struct {
	DNS   []string
	IP    []string
	Email []string
}

func (e *EnrollReq) Request() (*Request, error) {
	params := make(map[string][]string)
	params["format"] = []string{"json"}
	if strings.EqualFold(e.KeyType, "ecdsa") {
		params["key-type"] = []string{"ecdsa"}
		params["key-curve"] = []string{e.KeyCurve}
	} else {
		params["key-type"] = []string{"rsa"}
		params["key-size"] = []string{strconv.Itoa(e.KeySize)}
	}
	if e.CommonName == "" {
		return nil, fmt.Errorf("Common name must be specified")
	}
	params["cn"] = []string{e.CommonName}
	params["z"] = []string{e.Zone}
	if e.Chain != "" {
		params["chain"] = []string{e.Chain}
	}
	if e.KeyPassword != "" {
		params["key-password"] = []string{e.KeyPassword}
	}

	params["san-dns"] = e.Sans.DNS
	params["san-email"] = e.Sans.Email
	params["san-ip"] = e.Sans.IP
	req := &Request{
		Action:  "enroll",
		Mparams: params,
	}
	req.Params(params, "no-pickup", "no-prompt")

	return req, nil
}

// func (req *EnrollReq) paramMap() (map[string][]string, error) {
// 	params := make(map[string][]string)

// 	params["format"] = []string{"json"}
// 	// values with default
// 	if req.KeySize < 2048 {
// 		return params, fmt.Errorf("Key size must be 2048 or greater")
// 	}
// 	if req.CommonName == "" {
// 		return params, fmt.Errorf("Common name must be specified")
// 	}
// 	params["key-size"] = []string{strconv.Itoa(req.KeySize)}
// 	params["cn"] = []string{req.CommonName}
// 	if req.Zone != "" {
// 		params["z"] = []string{req.Zone}
// 	}
// 	if req.KeyType != "" {
// 		params["key-type"] = []string{req.KeyType}
// 	}
// 	if req.KeyType != "" {
// 		params["key-type"] = []string{req.KeyType}
// 	}
// 	if req.Chain != "" {
// 		params["chain"] = []string{req.Chain}
// 	}
// 	if req.Chain != "" {
// 		params["chain"] = []string{req.Chain}
// 	}

// 	params["san-dns"] = req.Sans.DNS
// 	params["san-email"] = req.Sans.Email
// 	params["san-ip"] = req.Sans.IP
// 	return params, nil
// }
