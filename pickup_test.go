package govcert

import "testing"

func TestPickupReq_Request(t *testing.T) {
	type fields struct {
		PickupID string
	}
	tests := []struct {
		name    string
		fields  fields
		want    func(r *request) error
		wantErr bool
	}{
		{
			"Fails without PickupID",
			fields{},
			func(r *request) error {
				return nil
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PickupReq{
				PickupID: tt.fields.PickupID,
			}
			got, err := p.Request()
			if (err != nil) != tt.wantErr {
				t.Errorf("PickupReq.Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err = tt.want(got); err != nil {
				t.Errorf("Error running test. Wanted nil, got %v", err)
			}
		})
	}
}

func TestPickupReq_RequiresAPI(t *testing.T) {
	type fields struct {
		PickupID string
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
			p := PickupReq{
				PickupID: tt.fields.PickupID,
			}
			if got := p.RequiresAPI(); got != tt.want {
				t.Errorf("PickupReq.RequiresAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
