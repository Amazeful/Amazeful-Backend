package channel

import (
	"context"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/dataful"
	"github.com/Amazeful/dataful/models"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ch *ChannelHandler) ChannelFromId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		channelId, err := primitive.ObjectIDFromHex(chi.URLParam(req, "channelId"))
		if err != nil {
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		r := ch.DB.Repository(dataful.DBAmazeful, dataful.CollectionChannel)
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
