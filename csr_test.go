package govcert

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"
)

func TestCSRReq_Request(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name    string
		fields  fields
		want    func(r *request) error
		wantErr bool
	}{
		{
			"Fails without CommonName",
			fields{
				Country: "UK",
				State:   "London",
			},
			func(r *request) error {
				return nil
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CSRReq{
				CommonName:         tt.fields.CommonName,
				OrganizationalUnit: tt.fields.OrganizationalUnit,
				OrganizationName:   tt.fields.OrganizationName,
				Country:            tt.fields.Country,
				State:              tt.fields.State,
				Locality:           tt.fields.Locality,
				KeyPassword:        tt.fields.KeyPassword,
				SanDNS:             tt.fields.SanDNS,
				SanEmail:           tt.fields.SanEmail,
				SanIP:              tt.fields.SanIP,
			}
			got, err := c.Request()
			if (err != nil) != tt.wantErr {
				t.Errorf("CSRReq.Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err = tt.want(got); err != nil {
				t.Errorf("Error running test. Wanted nil, got %v", err)
			}
		})
	}
}

func TestCSRReq_RequiresAPI(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"Will never require API key",
			fields{
				CommonName: "Mycert",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CSRReq{
				CommonName:         tt.fields.CommonName,
				OrganizationalUnit: tt.fields.OrganizationalUnit,
				OrganizationName:   tt.fields.OrganizationName,
				Country:            tt.fields.Country,
				State:              tt.fields.State,
				Locality:           tt.fields.Locality,
				KeyPassword:        tt.fields.KeyPassword,
				SanDNS:             tt.fields.SanDNS,
				SanEmail:           tt.fields.SanEmail,
				SanIP:              tt.fields.SanIP,
			}
			if got := c.RequiresAPI(); got != tt.want {
				t.Errorf("CSRReq.RequiresAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ParseCSR(t *testing.T) {
	type fields struct {
		errOut  *bytes.Buffer
		stdOut  *bytes.Buffer
		apiResp *http.Response
	}
	tests := []struct {
		name    string
		fields  fields
		want    *CSRResp
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := response{
				errOut: tt.fields.errOut,
				stdOut: tt.fields.stdOut,
			}
			got, err := r.ParseCSR()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ParseCSR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ParseCSR() = %v, want %v", got, tt.want)
			}
		})
	}
}
