package middlewares

import (
	"context"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/go-chi/chi/v5"
)

func ChannelCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		channelId := chi.URLParam(r, consts.CTXChannelId)
		//TODO: Get channel from DB
		ctx := context.WithValue(r.Context(), consts.CTXChannelId, channelId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
