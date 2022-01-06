package channel

import (
	"encoding/json"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/dataful/models"
)

func (ch *ChannelHandler) HandleGetChannel(rw http.ResponseWriter, req *http.Request) {
	channel, ok := req.Context().Value(consts.CtxChannel).(*models.Channel)
	if !ok {
		util.WriteError(rw, consts.ErrNoContextValue, http.StatusInternalServerError, consts.ErrStrResourceDNE)
		return
	}

	util.WriteResponse(rw, util.Response{Status: http.StatusOK, Data: channel})
}

func (ch *ChannelHandler) HandleUpdateChannel(rw http.ResponseWriter, req *http.Request) {
	channel, ok := req.Context().Value(consts.CtxChannel).(*models.Channel)
	if !ok {
		util.WriteError(rw, consts.ErrNoContextValue, http.StatusInternalServerError, consts.ErrUnexpected)
		return
	}

	newChannel := &models.Channel{}
	err := json.NewDecoder(req.Body).Decode(newChannel)
	if err != nil {
		util.WriteError(rw, err, http.StatusBadRequest, consts.ErrStrDecode)
		return
	}

	channel.Joined = newChannel.Joined
	channel.Silenced = newChannel.Silenced
	channel.Prefix = newChannel.Prefix

	err = channel.Update(req.Context())
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, consts.ErrStrDB)
		return
	}

	util.WriteResponse(rw, util.Response{Status: http.StatusOK, Data: channel})
}
