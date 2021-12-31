package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Amazeful/Amazeful-Backend/models"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/Amazeful-Backend/util/mocks"
	"github.com/go-redis/redis/v8"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticator(t *testing.T) {

	type args struct {
		sessionId string
		expiry    time.Time
	}

	tests := []struct {
		name           string
		authenticated  bool
		sessionExists  bool
		expectedStatus int
		args           args
	}{
		{"valid session", true, true, http.StatusOK, args{sessionId: "1", expiry: time.Now().Add(time.Minute)}},
		{"invalid session", true, false, http.StatusUnauthorized, args{sessionId: "1", expiry: time.Now().Add(time.Minute)}},
		{"invalid token", false, true, http.StatusUnauthorized, args{sessionId: "1", expiry: time.Now().Add(time.Minute)}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				util.WriteResponse(rw, util.Response{
					Status: http.StatusOK,
					Data:   nil,
				})
			})

			if test.authenticated {
				jwt := models.NewJWT([]byte("test"), jwa.HS256)
				tokenString, err := jwt.Encode(test.args.sessionId, test.args.expiry)
				assert.NoError(t, err)
				req.AddCookie(&http.Cookie{Name: JWTCookieName, Value: tokenString})
			}

			mockedRedis := new(mocks.Redis)
			util.SetRedis(mockedRedis)
			session := models.NewSession(mockedRedis)
			result := &redis.StringCmd{}

			if test.sessionExists {
				session.SessionId = test.args.sessionId
				data, err := json.Marshal(session)
				assert.NoError(t, err)
				result.SetVal(string(data))
			} else {
				result.SetErr(errors.New("not found"))
			}
			mockedRedis.On("Get", req.Context(), "session-"+test.args.sessionId).Return(result)
			authHandler := Authenticator(handler)
			authHandler.ServeHTTP(rw, req)
			assert.Equal(t, test.expectedStatus, rw.Code)
		})
	}
}
