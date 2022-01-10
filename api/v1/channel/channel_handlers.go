package channel

import (
	"encoding/json"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/dataful"
	"github.com/Amazeful/dataful/models"
)

//HandleGetChannel returns the current channel that is in request context.
func HandleGetChannel(rw http.ResponseWriter, req *http.Request) {
	channel, ok := req.Context().Value(consts.CtxChannel).(*models.Channel)
	if !ok {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	util.WriteResponse(rw, util.Response{Status: http.StatusOK, Data: channel})
}

//HandleUpdateChannel gets the channel from request context and updates it with new data in request body.
func HandleUpdateChannel(rw http.ResponseWriter, req *http.Request) {
	channel, ok := req.Context().Value(consts.CtxChannel).(*models.Channel)
	if !ok {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	newChannel := &models.Channel{}
	err := json.NewDecoder(req.Body).Decode(newChannel)
	if err != nil {
		util.WriteError(rw, err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	//We only want to update customizable values. Other values can only be updated using twitch api.
	channel.Joined = newChannel.Joined
	channel.Silenced = newChannel.Silenced
	channel.Prefix = newChannel.Prefix

	err = channel.Update(req.Context())
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	util.WriteResponse(rw, util.Response{Status: http.StatusOK, Data: channel})
}

//HandleGetChannelCommands gets the list of all commands for the channel.
func HandleGetChannelCommands(rw http.ResponseWriter, req *http.Request) {
	channel, ok := req.Context().Value(consts.CtxChannel).(*models.Channel)
	if !ok {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	r := util.GetDB().Repository(dataful.DBAmazeful, dataful.CollectionCommand)

	commandList := models.NewCommandList(r)

	err := commandList.LoadAllByChannel(req.Context(), channel.ID)
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	util.WriteResponse(rw, util.Response{
		Status: http.StatusOK,
		Data:   commandList.List,
	})

}

//HandleGetChannelFilters gets the channel filters.
func HandleGetChannelFilters(rw http.ResponseWriter, req *http.Request) {
	channel, ok := req.Context().Value(consts.CtxChannel).(*models.Channel)
	if !ok {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	r := util.GetDB().Repository(dataful.DBAmazeful, dataful.CollectionFilters)

	filters := models.NewFilters(r)
	err := filters.LoadByChannel(req.Context(), channel.ID)
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if !filters.Loaded() {
		http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	util.WriteResponse(rw, util.Response{
		Status: http.StatusOK,
		Data:   filters,
	})

}
