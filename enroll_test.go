package govcert

import (
	"fmt"
	"testing"
)

func TestEnrollReq_Request(t *testing.T) {
	type fields struct {
		CommonName  string
		KeyCurve    string
		Sans        SAN
		Format      string
		Zone        string
		KeySize     int
		KeyType     string
		Chain       string
		KeyPassword string
		APIKey      string
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
				Format:  "json",
				KeyType: "RSA",
			},
			func(r *request) error {
				return nil
			},
			true,
		},
		{
			"Sets key curve and not key size if ecdsa",
			fields{
				CommonName: "Tester",
				KeyType:    "ECDSA",
				KeyCurve:   "P224",
				KeySize:    2048,
			},
			func(r *request) error {
				foundCurve := false
				foundSize := false
				for _, v := range r.params {
					if v == "-key-curve" {
						foundCurve = true
					}
					if v == "-key-size" {
						foundSize = true
					}
				}
				if !foundCurve {
					return fmt.Errorf("key-curve not found in: %v", r.params)
				}
				if foundSize {
					return fmt.Errorf("key-size should not be set for ecdsa. %v", r.params)
				}

				return nil
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EnrollReq{
				CommonName:  tt.fields.CommonName,
				KeyCurve:    tt.fields.KeyCurve,
				Sans:        tt.fields.Sans,
				Format:      tt.fields.Format,
				Zone:        tt.fields.Zone,
				KeySize:     tt.fields.KeySize,
				KeyType:     tt.fields.KeyType,
				Chain:       tt.fields.Chain,
				KeyPassword: tt.fields.KeyPassword,
				APIKey:      tt.fields.APIKey,
			}
			got, err := e.Request()
			if (err != nil) != tt.wantErr {
				t.Errorf("EnrollReq.Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err = tt.want(got); err != nil {
				t.Errorf("Error running test. Wanted nil, got %v", err)
			}
		})
	}
}

func TestEnrollReq_RequiresAPI(t *testing.T) {
	type fields struct {
		CommonName  string
		KeyCurve    string
		Sans        SAN
		Format      string
		Zone        string
		KeySize     int
		KeyType     string
		Chain       string
		KeyPassword string
		APIKey      string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"Will always require API key",
			fields{
				CommonName: "Mycert",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EnrollReq{
				CommonName:  tt.fields.CommonName,
				KeyCurve:    tt.fields.KeyCurve,
				Sans:        tt.fields.Sans,
				Format:      tt.fields.Format,
				Zone:        tt.fields.Zone,
				KeySize:     tt.fields.KeySize,
				KeyType:     tt.fields.KeyType,
				Chain:       tt.fields.Chain,
				KeyPassword: tt.fields.KeyPassword,
				APIKey:      tt.fields.APIKey,
			}
			if got := e.RequiresAPI(); got != tt.want {
				t.Errorf("EnrollReq.RequiresAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
