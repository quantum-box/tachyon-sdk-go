package tachyoncms

import (
	"context"
	"testing"
	"time"

	tachyonid "github.com/quantum-box/tachyon-sdk-go/internal/id"
	cmspb "github.com/quantum-box/tachyon-sdk-go/service/cms/proto"
)

func TestClient_Create(t *testing.T) {
	conn, err := NewCmsClient()
	if err != nil {
		t.Error(err)
	}

	type fields struct {
		connection cmspb.CmsApiClient
	}
	type args struct {
		ctx context.Context
		in  *AggregateDto
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
			fields: fields{conn.connection},
			args: args{WithAuth(context.Background(), "Bearer some-auth-token"), &AggregateDto{
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
			}
			if err := c.Create(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("Client.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
