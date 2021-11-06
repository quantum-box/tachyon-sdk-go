package tachyoncms

import (
	"context"
	"reflect"
	"testing"
	"time"

	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
)

func TestClient_Update(t *testing.T) {
	client, err := NewCmsClient()
	if err != nil {
		t.Error(err)
	}
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
			args: args{WithAuth(context.Background(), "Bearer some-auth-token"), &AggregateDto{
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
			if err := c.Update(tt.args.ctx, tt.args.in, tt.args.aggregationName); (err != nil) != tt.wantErr {
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
			got, err := convUpdateRequest(tt.args.in, tt.args.aggregationName)
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
