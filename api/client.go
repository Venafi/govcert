package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"bytes"

	"github.com/davecgh/go-spew/spew"
	"github.com/opencredo/govcert"
)

type client struct {
	url    *url.URL
	APIKey string
}

func NewClient(apiurl, apikey string) (govcert.Client, error) {
	u, err := url.Parse(apiurl)
	if err != nil {
		return nil, err
	}
	c := &client{
		url:    u,
		APIKey: apikey,
	}
	return c, nil
}

func (c *client) Do(req *govcert.Request) (govcert.Response, error) {
	// c.url.Path = req.Action
	fmt.Println("---- API CLIENT DO CALLED ------")
	u := c.url
	u.Path = req.Action

	p := req.Mparams
	p["k"] = []string{c.APIKey}
	var r *http.Request
	var err error
	switch req.Method {
	case "POST":
		json, e := json.Marshal(p)
		if e != nil {
			return nil, e
		}
		buf := bytes.NewBuffer(json)
		spew.Dump(buf.String())
		r, err = http.NewRequest(req.Method, u.String(), buf)
		if err != nil {
			return nil, err
		}
	case "GET":
		q := u.Query()
		buildQuery(&q, p)
		spew.Dump(u.String())
		r, err = http.NewRequest(req.Method, u.String(), nil)
		if err != nil {
			return nil, err
		}
	}
	client := &http.Client{}
	apiresp, err := client.Do(r)
	resp := govcert.ResponseFromAPI(apiresp)

	return resp, err
}

func (c *client) Retryable(req *govcert.Request, wait time.Duration, maxWait time.Duration) (resp govcert.Response, err error) {
	for {
		resp, err = c.Do(req)
		spew.Dump(resp, err)
		return
	}
}

func buildQuery(q *url.Values, param map[string][]string) {
	for k, val := range param {
		for _, v := range val {
			q.Add(k, v)
		}
	}
}
