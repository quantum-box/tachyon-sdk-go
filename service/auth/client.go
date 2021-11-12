package tachyonauth

import "context"

type TachyonAuthorityDriver interface {
	Verify(ctx context.Context, token string) error
}
