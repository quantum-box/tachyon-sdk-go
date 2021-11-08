package testhelper

import (
	"context"

	myContext "github.com/quantum-box/tachyon-sdk-go/internal/context"
)

func NewContextWithToken() context.Context {
	return myContext.WithCurrentAuthzToken(context.Background(), "Bearer some-auth-token")
}
