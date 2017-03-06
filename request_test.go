package govcert

import "testing"

func Test_request_Params(t *testing.T) {
	type fields struct {
		Action  string
		Method  string
		params  []string
		Mparams map[string]RequestField
	}
	type args struct {
		p []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &request{
				Action:  tt.fields.Action,
				Method:  tt.fields.Method,
				params:  tt.fields.params,
				Mparams: tt.fields.Mparams,
			}
			r.Params(tt.args.p...)
		})
	}
}

func Test_request_parseMap(t *testing.T) {
	type fields struct {
		Action  string
		Method  string
		params  []string
		Mparams map[string]RequestField
	}
	type args struct {
		m map[string]RequestField
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &request{
				Action:  tt.fields.Action,
				Method:  tt.fields.Method,
				params:  tt.fields.params,
				Mparams: tt.fields.Mparams,
			}
			r.parseMap(tt.args.m)
		})
	}
}

func Test_request_hasAction(t *testing.T) {
	type fields struct {
		Action  string
		Method  string
		params  []string
		Mparams map[string]RequestField
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"When acion is set hasAction should return true",
			fields{
				Action: "Something",
			},
			true,
		},
		{
			"When not set, has action should be false",
			fields{
				Method: "blah",
			},
			false,
		},
		{
			"When action consists only of whitespace hasAction should be false",
			fields{
				Action: "            ",
			},
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &request{
				Action:  tt.fields.Action,
				Method:  tt.fields.Method,
				params:  tt.fields.params,
				Mparams: tt.fields.Mparams,
			}
			if got := r.hasAction(); got != tt.want {
				t.Errorf("request.hasAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inSlice(t *testing.T) {
	type args struct {
		s []string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inSlice(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("inSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
