package channel

import (
	"encoding/json"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/dataful"
	"github.com/Amazeful/dataful/models"
)

func HandleGetChannel(rw http.ResponseWriter, req *http.Request) {
	channel, ok := req.Context().Value(consts.CtxChannel).(*models.Channel)
	if !ok {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	util.WriteResponse(rw, util.Response{Status: http.StatusOK, Data: channel})
}

func HandleUpdateChannel(rw http.ResponseWriter, req *http.Request) {
	channel, ok := req.Context().Value(consts.CtxChannel).(*models.Channel)
	if !ok {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
		util.WriteError(rw, err, http.StatusInternalServerError, consts.ErrStrDB)
		return
	}

	if !commandList.Loaded() {
		http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	util.WriteResponse(rw, util.Response{
		Status: http.StatusOK,
		Data:   commandList.List,
	})

}

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
		util.WriteError(rw, err, http.StatusInternalServerError, consts.ErrStrDB)
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

// func HandleCreateCommand(rw http.ResponseWriter, req *http.Request) {
// 	channel, ok := req.Context().Value(consts.CtxChannel).(*models.Channel)
// 	if !ok {
// 		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	r := util.GetDB().Repository(dataful.DBAmazeful, dataful.CollectionCommand)
// 	command := models.NewCommand(r)
// 	err := json.NewDecoder(req.Body).Decode(command)
// 	if err != nil {
// 		util.WriteError(rw, err, http.StatusBadRequest, consts.ErrStrDecode)
// 		return
// 	}

// 	command.Channel = channel.ID
// 	err = command.Create(req.Context())

// 	util.WriteResponse(rw, util.Response{
// 		Status: http.StatusOK,
// 		Data:   commandList.List,
// 	})

// }
