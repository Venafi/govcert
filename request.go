package govcert

// Request is the call that will be sent to the Venafi SaaS
type Request struct {
	Action  string
	Method  string
	params  []string
	Mparams map[string][]string
}

// func NewRequest(c) *Request {
// 	return &Request{
// 		client: c,
// 	}
// }

// Help returns help text for the client or action if set
// func (r *Request) Help() (string, error) {
// 	r.help = true
// 	r.params = []string{"-h"}
// 	return r.Do()
// }

// func (r *Request) APIKey(key string) {
// 	r.client.APIKey = key
// }

// Do builds the request and captures output
// func (r *Request) Do() (string, error) {
// 	if r.hasAction() {
// 		r.client.cmd.Args = append(r.client.cmd.Args, r.Action)
// 	}
// 	if r.hasAPIKey() && !inSlice(r.params, "-k") {
// 		r.params = append(r.params, "-k", r.apiKey)
// 	}
// 	r.client.cmd.Stdout = r.client.output
// 	r.client.cmd.Stderr = r.client.errOutput
// 	r.client.cmd.Args = append(r.client.cmd.Args, r.params...)
// 	if err := r.client.cmd.Run(); err != nil && !r.help {
// 		// httpErr := parseHTTPError()
// 		// // httpErr :=
// 		// fmt.Println("ERROR", err)
// 		// spew.Dump(r)
// 		return r.client.parseError(), err
// 	}
// 	// r.client.cmd.Args = append(r.client.cmd.Args, )
// 	// r.client.cmd.
// 	return r.client.parseOutput(), nil
// }

// func (r *Request) ParamMap()

// Params accepts command parameters in multiple formats
func (r *Request) Params(p ...interface{}) {
	for _, param := range p {
		switch v := param.(type) {
		case string:
			r.params = append(r.params, "-"+v)
		case map[string][]string:
			r.parseMap(v)
		}
	}
}

func (r *Request) parseMap(m map[string][]string) {
	for p, vals := range m {
		for _, v := range vals {
			r.params = append(r.params, "-"+p, v)
		}
	}
}

func (r *Request) hasAction() bool {
	return len(r.Action) > 0
}

func inSlice(s []string, p string) bool {
	for _, v := range s {
		if v == p {
			return true
		}
	}
	return false
}

// func (r *Request) hasAPIKey() bool {
// 	return len(r.apiKey) > 0
// }

// func (r *Request) paramSet() ()
