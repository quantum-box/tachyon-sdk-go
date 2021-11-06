package tachyoncms

import (
	"context"
	"testing"

	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
)

func TestClient_Delete(t *testing.T) {
	conn, err := NewCmsClient()
	if err != nil {
		t.Error(err)
	}
	type fields struct {
		connection cmspb.CmsApiClient
	}
	type args struct {
		ctx             context.Context
		id              string
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
			fields:  fields{conn.connection},
			args:    args{WithAuth(context.Background(), "Bearer some-auth-token"), "01FKQJ5P12J772SMB4EMC21YS9", "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				connection: tt.fields.connection,
			}
			if err := c.Delete(tt.args.ctx, tt.args.id, tt.args.aggregationName); (err != nil) != tt.wantErr {
				t.Errorf("Client.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
