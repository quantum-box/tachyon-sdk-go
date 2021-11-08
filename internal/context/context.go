package context

import (
	"context"
	"errors"
)

type contextProjectID string
type contextAppID string
type contextAuthzToken string

const contextProjectIDKey contextProjectID = "projectID"
const contextAppIDKey contextAppID = "appID"
const contextAuthzTokenKey contextAuthzToken = "authzToken"

func CurrentProjectID(ctx context.Context) (string, error) {
	if v, ok := ctx.Value(contextProjectIDKey).(string); ok && v != "" {
		return v, nil
	}
	return "", errors.New("projectID is not found in context")
}

func WithCurrentProjectID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, contextProjectIDKey, id)
}

func CurrentAppID(ctx context.Context) (string, error) {
	if v, ok := ctx.Value(contextAppIDKey).(string); ok && v != "" {
		return v, nil
	}
	return "", errors.New("appID is not found in context")
}

func WithCurrentAppID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, contextAppIDKey, id)
}

func CurrentAuthzToken(ctx context.Context) (string, error) {
	if v, ok := ctx.Value(contextAuthzTokenKey).(string); ok && v != "" {
		return v, nil
	}
	return "", errors.New("authzToken is not found in context")
}

func WithCurrentAuthzToken(ctx context.Context, tokenStr string) context.Context {
	return context.WithValue(ctx, contextAuthzTokenKey, tokenStr)
}
