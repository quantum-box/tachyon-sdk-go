package tachyoncrm

import (
	"context"
	"reflect"
	"testing"

	crmpb "github.com/quantum-box/tachyon-sdk-go/service/crm/proto"
	"github.com/quantum-box/tachyon-sdk-go/tachyon"
)

func TestClient_GetByMail(t *testing.T) {
	type fields struct {
		connection crmpb.CrmApiClient
		config     *tachyon.Config
	}
	type args struct {
		ctx             context.Context
		aggregationName string
		mail            string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CustomerDto
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				connection: tt.fields.connection,
				config:     tt.fields.config,
			}
			got, err := c.GetByMail(tt.args.ctx, tt.args.aggregationName, tt.args.mail)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetByMail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetByMail() = %v, want %v", got, tt.want)
			}
		})
	}
}
