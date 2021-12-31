package channel

import (
	"encoding/json"
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

func HandleUpdateChannel(rw http.ResponseWriter, req *http.Request) {
	channel, ok := req.Context().Value(consts.CtxChannel).(*models.Channel)
	if !ok {
		util.WriteError(rw, consts.ErrNoContextValue, http.StatusInternalServerError, consts.ErrStrResourceDNE)
		return
	}

	newChannel := &models.Channel{}
	err := json.NewDecoder(req.Body).Decode(newChannel)
	if err != nil {
		util.WriteError(rw, err, http.StatusBadRequest, consts.ErrStrDecode)
		return
	}

	channel.UpdateCustomFields(newChannel)

	err = channel.Update(req.Context())
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, consts.ErrStrDB)
		return
	}

	util.WriteResponse(rw, util.Response{Status: http.StatusOK, Data: channel})
}
