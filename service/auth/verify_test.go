package tachyonauth

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	authpb "github.com/quantum-box/tachyon-sdk-go/service/auth/proto"
	mock_authpb "github.com/quantum-box/tachyon-sdk-go/service/auth/proto/mock_authority_grpc"
	"github.com/quantum-box/tachyon-sdk-go/tachyon"
)

func TestClient_Verify(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := mock_authpb.NewMockAuthorityApiClient(ctrl)
	ctx := context.Background()
	type fields struct {
		connection authpb.AuthorityApiClient
		config     *tachyon.Config
	}
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockFunc func()
		wantErr  bool
	}{
		{
			name:   "unittest success",
			fields: fields{mockClient, &tachyon.Config{}},
			args:   args{ctx, "some-token"},
			mockFunc: func() {
				mockClient.EXPECT().VerifyToken(ctx, &authpb.AuthorizeTokenRequest{
					Token: "some-token",
				}).Return(nil, nil)
			},
			wantErr: false,
		},
		{
			name:   "unittest failure",
			fields: fields{mockClient, &tachyon.Config{}},
			args:   args{ctx, "failure"},
			mockFunc: func() {
				mockClient.EXPECT().VerifyToken(ctx, &authpb.AuthorizeTokenRequest{
					Token: "failure",
				}).Return(nil, errors.New("not authorized"))
			},
			wantErr: true,
		},
		{
			name:     "unittest failure token is empty",
			fields:   fields{mockClient, &tachyon.Config{}},
			args:     args{ctx, ""},
			mockFunc: func() {},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			c := &Client{
				connection: tt.fields.connection,
				config:     tt.fields.config,
			}
			if err := c.Verify(tt.args.ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("Client.Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
