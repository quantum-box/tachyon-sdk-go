package tachyoncms

import (
	"context"
	"testing"

	"github.com/quantum-box/tachyon-sdk-go/internal/testhelper"
	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon"
)

func TestClient_FindAll(t *testing.T) {
	client, err := NewCmsClient(&tachyon.Config{"01FKXKQTWW7HNYQ8D5PFXC693D", "01FKXKS0VVMZS86G1P7A5NNH5H"})
	if err != nil {
		t.Error(err)
	}
	ctx := testhelper.NewContextWithToken()
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
			args:    args{ctx, "test"},
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
