package command

import (
	"encoding/json"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/models"
	"github.com/Amazeful/Amazeful-Backend/util"
)

func HandleCreateCommand(rw http.ResponseWriter, req *http.Request) {
	channel, ok := req.Context().Value(consts.CtxChannel).(*models.Channel)
	if !ok {
		util.WriteError(rw, consts.ErrNoContextValue, http.StatusInternalServerError, consts.ErrUnexpected)
		return
	}

	r := util.NewRepository(consts.DBAmazeful, consts.CollectionCommand)
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

	r := util.NewRepository(consts.DBAmazeful, consts.CollectionCommand)
	updatedCommand := models.NewCommand(r)

	err := json.NewDecoder(req.Body).Decode(updatedCommand)
	if err != nil {
		util.WriteError(rw, err, http.StatusBadRequest, consts.ErrStrDecode)
		return
	}

	command.UpdateCustomFields(updatedCommand)
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
