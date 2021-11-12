package tachyonauth

import (
	"context"
	"errors"

	authpb "github.com/quantum-box/tachyon-sdk-go/service/auth/proto"
)

func (c *Client) Verify(ctx context.Context, token string) error {
	if token == "" {
		return errors.New("token must not be empty")
	}
	return c.verify(ctx, token)
}

func (c *Client) verify(ctx context.Context, token string) error {
	_, err := c.connection.VerifyToken(ctx, &authpb.AuthorizeTokenRequest{
		Token: token,
	})
	if err != nil {
		return err
	}
	return nil
}
