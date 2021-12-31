package channel

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/models"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/Amazeful-Backend/util/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestChannelFromSession(t *testing.T) {

	type args struct {
		session *models.Session
		channel *models.Channel
	}

	tests := []struct {
		name           string
		existsInDB     bool
		expectedStatus int
		args           args
	}{
		{"no session", false, http.StatusUnauthorized, args{nil, nil}},
		{"not in db", false, http.StatusInternalServerError, args{&models.Session{SelectedChannel: primitive.ObjectID{}, SessionId: "1"}, &models.Channel{ChannelId: "123"}}},
		{"valid", true, http.StatusOK, args{&models.Session{SelectedChannel: primitive.ObjectID{}, SessionId: "1"}, &models.Channel{ChannelId: "123"}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if test.args.session != nil {
				req = req.WithContext(context.WithValue(req.Context(), consts.CtxSession, test.args.session))
			}
			handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				channel, ok := r.Context().Value(consts.CtxChannel).(*models.Channel)
				if !ok {
					util.WriteError(rw, consts.ErrNoContextValue, http.StatusInternalServerError, consts.ErrStrResourceDNE)
					return
				}
				util.WriteResponse(rw, util.Response{
					Status: http.StatusOK,
					Data:   channel,
				})
			})
			mockR := new(mocks.Repository)
			util.SetRepoGetter(func(dbName consts.MongoDatabase, collection consts.MongoCollection) util.Repository {
				return mockR
			})
			if test.existsInDB {
				mockR.On("FindOne", req.Context(), bson.M{"_id": primitive.ObjectID{}}, models.NewChannel(mockR)).Return(nil).Run(func(args mock.Arguments) {
					arg := args.Get(2).(*models.Channel)
					arg.ChannelId = test.args.channel.ChannelId
					arg.SetLoaded(true)
				})
			} else {
				mockR.On("FindOne", req.Context(), bson.M{"_id": primitive.ObjectID{}}, models.NewChannel(mockR)).Return(nil)

			}

			testHandler := ChannelFromSession(handler)
			testHandler.ServeHTTP(rw, req)
			assert.Equal(t, test.expectedStatus, rw.Code)

			if test.existsInDB {
				result := rw.Result()
				returnedChannel := &models.Channel{}
				err := json.NewDecoder(result.Body).Decode(returnedChannel)
				assert.NoError(t, err)

				assert.Equal(t, returnedChannel.ChannelId, test.args.channel.ChannelId)
			}
		})
	}
}
