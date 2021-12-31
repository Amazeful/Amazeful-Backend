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
