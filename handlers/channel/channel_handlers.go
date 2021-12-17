package channel

import (
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/models"
	"github.com/Amazeful/Amazeful-Backend/util"
)

func HandleGetChannel(rw http.ResponseWriter, req *http.Request) {
	collection := util.GetCollection(consts.CollectionChannel)

	channel := models.NewChannel(collection)

	err := channel.FindByChannelId(req.Context(), 123)
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, consts.ErrStrRetrieveData)
		return
	}

	util.WriteResponse(rw, util.Response{Status: http.StatusOK, Data: channel})

}
