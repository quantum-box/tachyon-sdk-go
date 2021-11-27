package tachyoncrm

import (
	"context"
	"testing"
	"time"

	tachyonid "github.com/quantum-box/tachyon-sdk-go/internal/id"
	crmpb "github.com/quantum-box/tachyon-sdk-go/service/crm/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon"
	"google.golang.org/grpc"
)

type crmApiClientMock struct {
}

func (c *crmApiClientMock) GetByMail(ctx context.Context, in *crmpb.GetByMailRequest, opts ...grpc.CallOption) (*crmpb.GetByMailResponse, error) {
	return nil, nil
}
func (c *crmApiClientMock) Create(ctx context.Context, in *crmpb.CreateRequest, opts ...grpc.CallOption) (*crmpb.CreateResponse, error) {
	return nil, nil
}
func (c *crmApiClientMock) Update(ctx context.Context, in *crmpb.UpdateRequest, opts ...grpc.CallOption) (*crmpb.UpdateResponse, error) {
	return nil, nil
}
func (c *crmApiClientMock) Delete(ctx context.Context, in *crmpb.DeleteRequest, opts ...grpc.CallOption) (*crmpb.DeleteResponse, error) {
	return nil, nil
}

func TestClient_Create(t *testing.T) {
	client, err := NewCrmClient(&tachyon.Config{
		AppID:     "01FKXKS0VVMZS86G1P7A5NNH5H",
		ProjectID: "01FKXKQTWW7HNYQ8D5PFXC693D",
	})
	if err != nil {
		t.Error(err)
	}
	type fields struct {
		connection crmpb.CrmApiClient
		config     *tachyon.Config
	}
	type args struct {
		ctx context.Context
		in  *CustomerDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "unittest",
			fields: fields{
				&crmApiClientMock{},
				nil,
			},
			args: args{
				ctx: context.Background(),
				in: &CustomerDto{
					ID:             tachyonid.NewUlID(),
					RegisteredAt:   time.Now(),
					LastSignedInAt: time.Now(),
				},
			},
			wantErr: false,
		},
		{
			name: "integrationtest",
			fields: fields{
				client.connection,
				client.config,
			},
			args: args{
				ctx: context.Background(),
				in: &CustomerDto{
					ID:             tachyonid.NewUlID(),
					RegisteredAt:   time.Now(),
					LastSignedInAt: time.Now(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				connection: tt.fields.connection,
				config:     tt.fields.config,
			}
			if err := c.Create(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("Client.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
