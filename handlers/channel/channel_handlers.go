package channel

import (
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/models"
	"github.com/Amazeful/Amazeful-Backend/util"
)

func HandleGetChannel(rw http.ResponseWriter, req *http.Request) {
	channel, ok := req.Context().Value(consts.CtxChannel).(*models.Channel)
	if !ok {
		util.WriteError(rw, consts.ErrNoContextValue, http.StatusInternalServerError, consts.ErrStrResourceDNE)
		return
	}

	util.WriteResponse(rw, util.Response{Status: http.StatusOK, Data: channel})
}
