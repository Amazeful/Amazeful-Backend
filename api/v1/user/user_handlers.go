package user

import (
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/dataful/models"
)

func HandleGetUser(rw http.ResponseWriter, req *http.Request) {
	user, ok := req.Context().Value(consts.CtxUser).(*models.User)
	if !ok {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	util.WriteResponse(rw, util.Response{Status: http.StatusOK, Data: user})
}
