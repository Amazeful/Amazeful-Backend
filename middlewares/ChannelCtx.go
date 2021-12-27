package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/models"
	"github.com/Amazeful/Amazeful-Backend/util"
)

func ChannelCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		session, ok := req.Context().Value(consts.CtxSession).(*models.Session)
		if !ok {
			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		collection := util.GetCollection(consts.CollectionChannel)
		channel := models.NewChannel(collection)
		err := channel.FindBylId(req.Context(), session.SelectedChannel)
		if err != nil {
			util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		if !channel.Loaded() {
			err = fmt.Errorf("selected channelid %s in session does not exist in DB", session.SelectedChannel.String())
			util.WriteError(rw, err, http.StatusInternalServerError, consts.ErrStrResourceDNE)
			return
		}
		ctx := context.WithValue(req.Context(), consts.CtxChannel, channel)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
