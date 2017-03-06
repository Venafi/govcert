package govcert

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNewResponse(t *testing.T) {
	tests := []struct {
		name string
		want *response
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_Body(t *testing.T) {
	type fields struct {
		errOut *bytes.Buffer
		stdOut *bytes.Buffer
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &response{
				errOut: tt.fields.errOut,
				stdOut: tt.fields.stdOut,
			}
			got, err := r.Body()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.Body() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.Body() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_RequestID(t *testing.T) {
	type fields struct {
		errOut *bytes.Buffer
		stdOut *bytes.Buffer
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &response{
				errOut: tt.fields.errOut,
				stdOut: tt.fields.stdOut,
			}
			got, err := r.RequestID()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.RequestID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.RequestID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_CompletedID(t *testing.T) {
	type fields struct {
		errOut *bytes.Buffer
		stdOut *bytes.Buffer
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &response{
				errOut: tt.fields.errOut,
				stdOut: tt.fields.stdOut,
			}
			got, err := r.CompletedID()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.CompletedID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.CompletedID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_JSONBody(t *testing.T) {
	type fields struct {
		errOut *bytes.Buffer
		stdOut *bytes.Buffer
	}
	tests := []struct {
		name    string
		fields  fields
		wantJ   map[string]interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &response{
				errOut: tt.fields.errOut,
				stdOut: tt.fields.stdOut,
			}
			gotJ, err := r.JSONBody()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.JSONBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotJ, tt.wantJ) {
				t.Errorf("response.JSONBody() = %v, want %v", gotJ, tt.wantJ)
			}
		})
	}
}

func Test_response_Pending(t *testing.T) {
	type fields struct {
		errOut *bytes.Buffer
		stdOut *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &response{
				errOut: tt.fields.errOut,
				stdOut: tt.fields.stdOut,
			}
			if got := r.Pending(); got != tt.want {
				t.Errorf("response.Pending() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_Bytes(t *testing.T) {
	type fields struct {
		errOut *bytes.Buffer
		stdOut *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &response{
				errOut: tt.fields.errOut,
				stdOut: tt.fields.stdOut,
			}
			if got := r.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
