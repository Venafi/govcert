package govcert

import (
	"fmt"
	"strings"
)

type PickupReq struct {
	PickupID string
}

func (p *PickupReq) Request() (*request, error) {
	params := make(map[string]RequestField)
	params["format"] = RequestField{"json"}
	if len(strings.TrimSpace(p.PickupID)) == 0 {
		return nil, fmt.Errorf("Cannot make pickup request without pickupID")
	}
	params["pickup-id"] = RequestField{p.PickupID}

	req := &request{
		Action:  "pickup",
		Mparams: params,
	}
	req.Params(params, "no-prompt")

	return req, nil
}

func (PickupReq) RequiresAPI() bool {
	return true
}
