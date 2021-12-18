package user

import (
	"encoding/json"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/models"
	"github.com/Amazeful/Amazeful-Backend/util"
)

func HandleGetUser(rw http.ResponseWriter, req *http.Request) {
	collection := util.GetCollection(consts.CollectionUser)

	user := models.NewUser(collection)

	err := user.FindByUserId(req.Context(), 123)
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, consts.ErrStrRetrieveData)
		return
	}

	util.WriteResponse(rw, util.Response{Status: http.StatusOK, Data: user})

}

func HandleCreateUser(rw http.ResponseWriter, req *http.Request) {
	collection := util.GetCollection(consts.CollectionUser)

	user := models.NewUser(collection)

	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, consts.ErrStrDecode)
		return
	}

	err = user.Create(req.Context())
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, consts.ErrStrInsert)
		return
	}

	util.WriteResponse(rw, util.Response{Status: http.StatusCreated, Data: user})
}
