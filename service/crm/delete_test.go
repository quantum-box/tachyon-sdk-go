package tachyoncrm

import (
	"context"
	"testing"

	tachyonid "github.com/quantum-box/tachyon-sdk-go/internal/id"
	crmpb "github.com/quantum-box/tachyon-sdk-go/service/crm/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon"
)

func TestClient_Delete(t *testing.T) {
	type fields struct {
		connection crmpb.CrmApiClient
		config     *tachyon.Config
	}
	type args struct {
		ctx             context.Context
		aggregationName string
		id              string
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
				ctx:             context.Background(),
				aggregationName: "unittest",
				id:              tachyonid.NewUlID(),
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
			if err := c.Delete(tt.args.ctx, tt.args.aggregationName, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Client.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
