package tachyoncms

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/quantum-box/tachyon-sdk-go/internal/testhelper"
	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon/config"
)

func TestClient_Update(t *testing.T) {
	client, err := NewCmsClient(&config.Config{"01FKXKQTWW7HNYQ8D5PFXC693D", "01FKXKS0VVMZS86G1P7A5NNH5H"})
	if err != nil {
		t.Error(err)
	}
	ctx := testhelper.NewContextWithToken()
	type fields struct {
		connection cmspb.CmsApiClient
	}
	type args struct {
		ctx             context.Context
		in              *AggregateDto
		aggregationName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "integration update test",
			fields: fields{client.connection},
			args: args{ctx, &AggregateDto{
				ID:        "01FKNYSFQ4P7F8ETGR0ZGE9N6B",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
				Data: map[string]interface{}{
					"message": "updated!",
				},
			}, "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				connection: tt.fields.connection,
			}
			if err := c.Update(tt.args.ctx, tt.args.aggregationName, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("Client.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_convUpdateRequest(t *testing.T) {
	type args struct {
		in              *AggregateDto
		aggregationName string
	}
	tests := []struct {
		name    string
		args    args
		want    *cmspb.UpdateRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convUpdateRequest(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("convUpdateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convUpdateRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
