package channel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/models"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ChannelFromSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		session, ok := req.Context().Value(consts.CtxSession).(*models.Session)
		if !ok {
			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		r := util.NewRepository(consts.DBAmazeful, consts.CollectionChannel)
		channel := models.NewChannel(r)
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

func ChannelFromId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		channelId, err := primitive.ObjectIDFromHex(chi.URLParam(req, "channelId"))
		if err != nil {
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		r := util.NewRepository(consts.DBAmazeful, consts.CollectionChannel)
		channel := models.NewChannel(r)
		err = channel.FindBylId(req.Context(), channelId)
		if err != nil {
			util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		if !channel.Loaded() {
			http.Error(rw, consts.ErrStrResourceDNE, http.StatusNotFound)
			return
		}

		ctx := context.WithValue(req.Context(), consts.CtxChannel, channel)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}

func ChannelFromParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		channelName := chi.URLParam(req, "channelName")
		if channelName == "" {
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		r := util.NewRepository(consts.DBAmazeful, consts.CollectionChannel)
		channel := models.NewChannel(r)
		err := channel.FindByChannelName(req.Context(), channelName)
		if err != nil {
			util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		if !channel.Loaded() {
			http.Error(rw, consts.ErrStrResourceDNE, http.StatusNotFound)
			return
		}
		ctx := context.WithValue(req.Context(), consts.CtxChannel, channel)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
