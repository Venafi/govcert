package govcert

import (
	"os/exec"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		path   string
		apikey string
	}
	tests := []struct {
		name string
		args args
		want *client
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.path, tt.args.apikey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Do(t *testing.T) {
	type fields struct {
		cmd     *exec.Cmd
		cmdPath string
		apiKey  string
	}
	type args struct {
		r Requestor
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Response
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cmd:     tt.fields.cmd,
				cmdPath: tt.fields.cmdPath,
				apiKey:  tt.fields.apiKey,
			}
			got, err := c.Do(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_parse(t *testing.T) {
	type fields struct {
		cmd     *exec.Cmd
		cmdPath string
		apiKey  string
	}
	type args struct {
		out []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cmd:     tt.fields.cmd,
				cmdPath: tt.fields.cmdPath,
				apiKey:  tt.fields.apiKey,
			}
			if got := c.parse(tt.args.out); got != tt.want {
				t.Errorf("client.parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_APIKey(t *testing.T) {
	type fields struct {
		cmd     *exec.Cmd
		cmdPath string
		apiKey  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cmd:     tt.fields.cmd,
				cmdPath: tt.fields.cmdPath,
				apiKey:  tt.fields.apiKey,
			}
			if got := c.APIKey(); got != tt.want {
				t.Errorf("client.APIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_hasAPIKey(t *testing.T) {
	type fields struct {
		cmd     *exec.Cmd
		cmdPath string
		apiKey  string
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
			c := &client{
				cmd:     tt.fields.cmd,
				cmdPath: tt.fields.cmdPath,
				apiKey:  tt.fields.apiKey,
			}
			if got := c.hasAPIKey(); got != tt.want {
				t.Errorf("client.hasAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
