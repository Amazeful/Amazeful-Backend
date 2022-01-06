package channel

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/dataful"
	"github.com/Amazeful/dataful/mocks"
	"github.com/Amazeful/dataful/models"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestChannelFromSession(t *testing.T) {
	t.Parallel()
	type args struct {
		id      string
		channel *models.Channel
	}

	tests := []struct {
		name           string
		existsInDB     bool
		existsInURL    bool
		expectedStatus int
		args           args
	}{
		{"not in url", false, false, http.StatusBadRequest, args{id: "507f1f77bcf86cd799439011", channel: &models.Channel{BroadcasterName: "Amazeful"}}},
		{"not in db", false, true, http.StatusNotFound, args{id: "507f1f77bcf86cd799439011", channel: &models.Channel{BroadcasterName: "Amazeful"}}},
		{"all valid", true, true, http.StatusOK, args{id: "507f1f77bcf86cd799439011", channel: &models.Channel{BroadcasterName: "Amazeful"}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if test.existsInURL {
				ctx := chi.NewRouteContext()
				ctx.URLParams.Add("channelId", test.args.id)
				req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
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
			db := new(mocks.Database)
			repository := new(mocks.Repository)
			db.On("Repository", dataful.DBAmazeful, dataful.CollectionChannel).Return(repository)
			if test.existsInDB {
				repository.On("FindOne", req.Context(), mock.Anything, models.NewChannel(repository), mock.Anything).Return(nil).Run(func(args mock.Arguments) {
					arg := args.Get(2).(*models.Channel)
					arg.BroadcasterName = test.args.channel.BroadcasterName
					arg.SetLoaded(true)
				})
			} else {
				repository.On("FindOne", req.Context(), mock.Anything, models.NewChannel(repository), mock.Anything).Return(nil)
			}

			channelHandler := NewChannelHandler(&util.Resources{DB: db})

			server := channelHandler.ChannelFromId(handler)

			server.ServeHTTP(rw, req)

			assert.Equal(t, test.expectedStatus, rw.Code)
			if test.existsInDB {
				result := rw.Result()
				returnedChannel := &models.Channel{}
				err := json.NewDecoder(result.Body).Decode(returnedChannel)
				assert.NoError(t, err)

				assert.Equal(t, test.args.channel.BroadcasterName, returnedChannel.BroadcasterName)
			}
		})
	}
}
