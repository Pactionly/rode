package auth

import (
	"context"
	"encoding/base64"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/rode/rode/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type authenticator struct {
	authConfig *config.AuthConfig
}

type Authenticator interface {
	Authenticate(ctx context.Context) (context.Context, error)
}

func NewAuthenticator(authConfig *config.AuthConfig) Authenticator {
	return &authenticator{
		authConfig: authConfig,
	}
}

func (a *authenticator) Authenticate(ctx context.Context) (context.Context, error) {
	if a.authConfig.Basic.Username != "" && a.authConfig.Basic.Password != "" {
		return a.basic(ctx)
	}

	if a.authConfig.JWT.Issuer != "" {
		return a.jwt(ctx)
	}

	return ctx, nil
}

func (a *authenticator) basic(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "basic")
	if err != nil {
		return nil, err
	}

	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "error decoding auth token: %v", err)
	}

	parts := strings.Split(string(data), ":")
	if len(parts) != 2 {
		return nil, status.Errorf(codes.Unauthenticated, "expected auth token to follow format ${username}:${password}")
	}

	if a.authConfig.Basic.Username == parts[0] && a.authConfig.Basic.Password == parts[1] {
		return ctx, nil
	}

	return nil, status.Error(codes.Unauthenticated, "invalid username or password")
}

func (a *authenticator) jwt(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	_, err = a.authConfig.JWT.Verifier.Verify(ctx, token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "error validating jwt: %v", err)
	}

	return ctx, nil
}
