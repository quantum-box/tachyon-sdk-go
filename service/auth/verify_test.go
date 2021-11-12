package tachyonauth

import (
	"context"
	"testing"
)

func TestClient_Verify(t *testing.T) {
	type fields struct {
		connection authpb.AuthorityApiClient
		config     *tachyon.Config
	}
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				connection: tt.fields.connection,
				config:     tt.fields.config,
			}
			if err := c.Verify(tt.args.ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("Client.Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
