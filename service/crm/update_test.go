package tachyoncrm

import (
	"context"
	"testing"
	"time"

	tachyonid "github.com/quantum-box/tachyon-sdk-go/internal/id"
	"github.com/quantum-box/tachyon-sdk-go/internal/testhelper"
	crmpb "github.com/quantum-box/tachyon-sdk-go/service/crm/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon"
)

func TestClient_Update(t *testing.T) {
	client, err := NewCrmClient(&tachyon.Config{
		AppID:     "01FKXKS0VVMZS86G1P7A5NNH5H",
		ProjectID: "01FKXKQTWW7HNYQ8D5PFXC693D",
	},
	)
	if err != nil {
		t.Error(err)
	}
	ctx := testhelper.NewContextWithToken()
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
				client.connection,
				client.config,
			},
			args: args{
				ctx,
				&CustomerDto{
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
			if err := c.Update(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("Client.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
