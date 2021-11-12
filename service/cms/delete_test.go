package tachyoncms

import (
	"context"
	"testing"

	"github.com/quantum-box/tachyon-sdk-go/internal/testhelper"
	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon/config"
)

func TestClient_Delete(t *testing.T) {
	conn, err := NewCmsClient(&config.Config{
		ProjectID: "01FKXKQTWW7HNYQ8D5PFXC693D", AppID: "01FKXKS0VVMZS86G1P7A5NNH5H"})
	if err != nil {
		t.Error(err)
	}
	ctx := testhelper.NewContextWithToken()
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
			args:    args{ctx, "01FKQJ5P12J772SMB4EMC21YS9", "test"},
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
