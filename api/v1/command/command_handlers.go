package command

import (
	"encoding/json"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/dataful"
	"github.com/Amazeful/dataful/models"
)

//HandleGetCommand returns the current command that is in request context.
func HandleGetCommand(rw http.ResponseWriter, req *http.Request) {
	command, ok := req.Context().Value(consts.CtxCommand).(*models.Command)
	if !ok {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	util.WriteResponse(rw, util.Response{
		Status: http.StatusOK,
		Data:   command,
	})
}

//HandleCreateCommand creates a new command.
func HandleCreateCommand(rw http.ResponseWriter, req *http.Request) {
	r := util.GetDB().Repository(dataful.DBAmazeful, dataful.CollectionCommand)
	command := models.NewCommand(r)
	err := json.NewDecoder(req.Body).Decode(command)
	if err != nil {
		util.WriteError(rw, err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = command.Create(req.Context())
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	util.WriteResponse(rw, util.Response{
		Status: http.StatusCreated,
		Data:   command,
	})
}

//HandleUpdateCommand gets the command from request context and updates it with new data in request body.
func HandleUpdateCommand(rw http.ResponseWriter, req *http.Request) {
	command, ok := req.Context().Value(consts.CtxCommand).(*models.Command)
	if !ok {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	r := util.GetDB().Repository(dataful.DBAmazeful, dataful.CollectionCommand)
	updatedCommand := models.NewCommand(r)

	err := json.NewDecoder(req.Body).Decode(updatedCommand)
	if err != nil {
		util.WriteError(rw, err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
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

	err = command.Update(req.Context())
	if err != nil {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	util.WriteResponse(rw, util.Response{
		Status: http.StatusOK,
		Data:   command,
	})
}

//HandleDeleteCommand deletes the command.
func HandleDeleteCommand(rw http.ResponseWriter, req *http.Request) {
	command, ok := req.Context().Value(consts.CtxCommand).(*models.Command)
	if !ok {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err := command.Delete(req.Context())
	if err != nil {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	util.WriteResponse(rw, util.Response{
		Status: http.StatusOK,
	})
}
