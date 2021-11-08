package tachyoncms

import (
	"context"
	"testing"

	"github.com/quantum-box/tachyon-sdk-go/internal/testhelper"
	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon"
)

func TestClient_FindByPath(t *testing.T) {
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
		paths           []string
		value           string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*AggregateDto
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "integration test for find_by_path",
			fields:  fields{client.connection},
			args:    args{ctx, "test", []string{"data", "message"}, "ok"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				connection: tt.fields.connection,
			}
			got, err := c.FindByPath(tt.args.ctx, tt.args.aggregationName, tt.args.paths, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.FindByPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
