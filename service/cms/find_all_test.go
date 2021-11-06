package tachyoncms

import (
	"context"
	"testing"

	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
)

func TestClient_FindAll(t *testing.T) {
	client, err := NewCmsClient()
	if err != nil {
		t.Error(err)
	}
	type fields struct {
		connection cmspb.CmsApiClient
	}
	type args struct {
		ctx             context.Context
		aggregationName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "integrate test",
			fields:  fields{client.connection},
			args:    args{WithAuth(context.Background(), "Bearer some-auth-token"), "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				connection: tt.fields.connection,
			}
			got, err := c.FindAll(tt.args.ctx, tt.args.aggregationName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
