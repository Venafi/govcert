package govcert

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

type CSRResp struct {
	PrivateKey string
	CSR        string
}

func (c *client) CSR(req *CSRReq) (*CSRResp, error) {
	resp := &CSRResp{}
	return resp, nil
}
