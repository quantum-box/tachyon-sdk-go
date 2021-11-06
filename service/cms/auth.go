package tachyoncms

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func WithAuth(ctx context.Context, token string) context.Context {
	md := metadata.New(map[string]string{"authorization": token})
	return metadata.NewOutgoingContext(ctx, md)
}
