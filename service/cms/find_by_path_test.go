package tachyoncms

import (
	"context"
	"testing"

	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
)

func TestClient_FindByPath(t *testing.T) {
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
			args:    args{WithAuth(context.Background(), "Bearer some-auth-token"), "test", []string{"data", "message"}, "ok"},
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
