package tachyoncms

import (
	"context"
	"testing"
	"time"

	tachyonid "github.com/quantum-box/tachyon-sdk-go/internal/id"
	"github.com/quantum-box/tachyon-sdk-go/internal/testhelper"
	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon/config"
)

func TestClient_Create(t *testing.T) {
	conn, err := NewCmsClient(&config.Config{
		ProjectID: "01FKXKQTWW7HNYQ8D5PFXC693D", AppID: "01FKXKS0VVMZS86G1P7A5NNH5H"})
	if err != nil {
		t.Error(err)
	}
	ctx := testhelper.NewContextWithToken()

	type fields struct {
		connection cmspb.CmsApiClient
		config     *config.Config
	}
	type args struct {
		ctx             context.Context
		aggregationName string
		in              *AggregateDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			// cms-service connection required
			name:   "integrate test",
			fields: fields{conn.connection, conn.config},
			args: args{ctx, "test", &AggregateDto{
				ID:        tachyonid.NewUlID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
				Data: map[string]interface{}{
					"message": "ok",
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				connection: tt.fields.connection,
				config:     tt.fields.config,
			}
			if err := c.Create(tt.args.ctx, tt.args.aggregationName, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("Client.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
