package govcert

type PickupReq struct {
	PickupID string
}

func (p *PickupReq) Request() (*Request, error) {
	params := make(map[string]RequestField)
	params["format"] = RequestField{"json"}
	params["pickup-id"] = RequestField{p.PickupID}

	req := &Request{
		Action:  "pickup",
		Mparams: params,
	}
	req.Params(params, "no-prompt")

	return req, nil
}
