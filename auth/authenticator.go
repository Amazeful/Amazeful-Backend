package auth

import (
	"context"
	"net/http"
)

type Authenticator interface {
	CreateToken(ctx context.Context, rw http.ResponseWriter, userId, channelId string) error
	Authenticate(next http.Handler) http.Handler
}
