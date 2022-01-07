package command

import (
	"encoding/json"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/dataful"
	"github.com/Amazeful/dataful/models"
)

func HandleCreateCommand(rw http.ResponseWriter, req *http.Request) {
	channel, ok := req.Context().Value(consts.CtxChannel).(*models.Channel)
	if !ok {
		util.WriteError(rw, consts.ErrNoContextValue, http.StatusInternalServerError, consts.ErrUnexpected)
		return
	}

	r := util.GetDB().Repository(dataful.DBAmazeful, dataful.CollectionCommand)
	command := models.NewCommand(r)

	err := json.NewDecoder(req.Body).Decode(command)
	if err != nil {
		util.WriteError(rw, err, http.StatusBadRequest, consts.ErrStrDecode)
		return
	}

	command.ID = channel.ID

	err = command.Create(req.Context())
	if err != nil {
		util.WriteError(rw, err, http.StatusBadRequest, consts.ErrStrDB)
		return
	}

	util.WriteResponse(rw, util.Response{
		Status: http.StatusCreated,
		Data:   command,
	})
}

func HandleGetCommand(rw http.ResponseWriter, req *http.Request) {
	command, ok := req.Context().Value(consts.CtxCommand).(*models.Command)
	if !ok {
		util.WriteError(rw, consts.ErrNoContextValue, http.StatusInternalServerError, consts.ErrUnexpected)
		return
	}

	util.WriteResponse(rw, util.Response{
		Status: http.StatusOK,
		Data:   command,
	})
}

func HandleUpdateCommand(rw http.ResponseWriter, req *http.Request) {
	command, ok := req.Context().Value(consts.CtxCommand).(*models.Command)
	if !ok {
		util.WriteError(rw, consts.ErrNoContextValue, http.StatusInternalServerError, consts.ErrUnexpected)
		return
	}

	r := util.GetDB().Repository(dataful.DBAmazeful, dataful.CollectionCommand)
	updatedCommand := models.NewCommand(r)

	err := json.NewDecoder(req.Body).Decode(updatedCommand)
	if err != nil {
		util.WriteError(rw, err, http.StatusBadRequest, consts.ErrStrDecode)
		return
	}

	command.Enabled = updatedCommand.Enabled
	command.Cooldowns = updatedCommand.Cooldowns
	command.Role = updatedCommand.Role
	command.Stream = updatedCommand.Stream
	command.Response = updatedCommand.Response
	command.Aliases = updatedCommand.Aliases
	command.Attributes = updatedCommand.Attributes
	command.Timer = updatedCommand.Timer
	command.Attributes = updatedCommand.Attributes

	err = command.Create(req.Context())
	if err != nil {
		util.WriteError(rw, err, http.StatusBadRequest, consts.ErrStrDB)
		return
	}

	util.WriteResponse(rw, util.Response{
		Status: http.StatusOK,
		Data:   command,
	})
}

func HandleDeleteCommand(rw http.ResponseWriter, req *http.Request) {
	command, ok := req.Context().Value(consts.CtxCommand).(*models.Command)
	if !ok {
		util.WriteError(rw, consts.ErrNoContextValue, http.StatusInternalServerError, consts.ErrUnexpected)
		return
	}

	err := command.Delete(req.Context())
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, consts.ErrStrDB)
		return
	}

	util.WriteResponse(rw, util.Response{
		Status: http.StatusOK,
		Data:   command,
	})
}
