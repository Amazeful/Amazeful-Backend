package middlewares

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

//CommandFromId middleware adds the command data to request context using commandId url param.
func CommandFromId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		commandId, err := primitive.ObjectIDFromHex(chi.URLParam(req, "commandId"))
		if err != nil {
			util.WriteError(rw, err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}

		r := util.DB().Repository(dataful.DBAmazeful, dataful.CollectionCommand)
		command := models.NewCommand(r)
		err = command.LoadBylId(req.Context(), commandId)
		if err != nil {
			util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		if !command.Loaded() {
			http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		ctx := context.WithValue(req.Context(), consts.CtxCommand, command)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
