package util

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"
)

type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

//WriteError writes errString to http and logs the actual error
func WriteError(rw http.ResponseWriter, err error, errCode int, errString string) {
	http.Error(rw, errString, errCode)
	GetLogger().Error(errString, zap.Error(err))
}

//WriteResponse writes http response
func WriteResponse(rw http.ResponseWriter, data Response) {
	js, err := json.Marshal(data.Data)
	if err != nil {
		WriteError(rw, errors.New("request failed. please try again later"), http.StatusInternalServerError, err.Error())
		return
	}

	if data.Message != "" {
		GetLogger().Info(data.Message)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(data.Status)
	rw.Write(js)
}
