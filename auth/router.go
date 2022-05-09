package auth

import (
	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/rest"
	"github.com/Amazeful/dataful/db"
	cachestore "github.com/Amazeful/dataful/store/cache"
	mongostore "github.com/Amazeful/dataful/store/mongo"

	chi "github.com/go-chi/chi/v5"
)

type AuthRouter struct {
	*rest.RouterCommon
	twitchConfig *config.TwitchConfig
}

func NewAuthRouter(common *rest.RouterCommon) *AuthRouter {
	return &AuthRouter{
		RouterCommon: common,
	}
}

func (ar *AuthRouter) Process(r chi.Router) {
	authenticator := NewJWTAuthenticator(cachestore.NewCacheSessionStore(ar.RouterCommon.Cache), []byte(ar.RouterCommon.ServerConfig.JwtSignKey))
	channelStore := mongostore.NewMongoChannelStore(ar.RouterCommon.DB.Collection(db.DBAmazeful, db.CollectionChannel))
	userStore := mongostore.NewMongoUserStore(ar.RouterCommon.DB.Collection(db.DBAmazeful, db.CollectionUser))
	twitch := NewTwitchAuth(ar.twitchConfig, ar.RouterCommon.ServerConfig, ar.RouterCommon.ResponseWriter, channelStore, userStore, authenticator)

	r.Route("/twitch", func(r chi.Router) {
		r.Get("/login", twitch.HandleTwitchLogin)
		r.Get("/callback", twitch.HandleTwitchCallback)
	})

}
