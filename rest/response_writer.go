package rest

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type ResponseWriter interface {
	WriteJson(rw http.ResponseWriter, data interface{}, status int)
	WriteError(rw http.ResponseWriter, status int, err error, message string)
	WriteCookie(rw http.ResponseWriter, cookie *http.Cookie)
}

type DefaultResponseWriter struct {
	logger *zap.Logger
}

func NewDefaultResponseWriter() *DefaultResponseWriter {
	return &DefaultResponseWriter{}
}

func (drw *DefaultResponseWriter) WriteJson(rw http.ResponseWriter, data *SuccessResponse, status int) {
	b, err := json.Marshal(data)
	if err != nil {
		drw.WriteError(rw, NewErrorResponse(nil, http.StatusInternalServerError, MessageInternalServerError, err))
		return
	}
	drw.addJsonHeaders(rw)
	rw.WriteHeader(status)
	rw.Write(b)
}

func (drw *DefaultResponseWriter) WriteError(rw http.ResponseWriter, errResponse *ErrorResponse) {
	drw.logger.Error("responding with error", zap.Any("errResponse", errResponse), zap.Error(errResponse.err))

	b, err := json.Marshal(errResponse)
	if err != nil {
		http.Error(rw, errResponse.Error, errResponse.Status)
		return
	}
	drw.addJsonHeaders(rw)
	rw.WriteHeader(errResponse.Status)
	rw.Write(b)
}

func (drw *DefaultResponseWriter) WriteCookie(rw http.ResponseWriter, cookie *http.Cookie) {
	http.SetCookie(rw, cookie)
}

func (drw *DefaultResponseWriter) addJsonHeaders(rw http.ResponseWriter) {
	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Add("X-Frame-Options", "deny")
}
