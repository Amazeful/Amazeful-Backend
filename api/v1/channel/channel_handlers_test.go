package channel

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/dataful"
	"github.com/Amazeful/dataful/models"
	"github.com/go-playground/assert/v2"
)

// import (
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/Amazeful/Amazeful-Backend/consts"
// 	"github.com/Amazeful/dataful/models"
// 	"github.com/go-playground/assert/v2"
// )

// // func TestHandleGetChannel(t *testing.T) {
// // 	err := config.LoadConfig()
// // 	assert.NoError(t, err)

// // 	err = util.InitDB(context.Background())
// // 	assert.NoError(t, err)
// // }

func TestHandleGetChannel(t *testing.T) {
	r := util.DB().Repository(dataful.DBAmazeful, dataful.CollectionChannel)
	testChannel := &models.Channel{
		BaseModel:       dataful.NewBaseModel(r),
		ChannelId:       "138760387",
		BroadcasterName: "Amazeful",
		Language:        "en",
		GameId:          "1",
		GameName:        "Just Chatting",
		Title:           "FeelsDankMan",
		Joined:          true,
		Prefix:          "!",
		Shard:           1,
		StartedAt:       time.Now().Add(-time.Hour).UTC(),
		Moderator:       true,
	}

	handler := http.HandlerFunc(HandleGetChannel)

	tests := []struct {
		name           string
		hasContext     bool
		expectedStatus int
	}{
		{"valid context", true, http.StatusOK},
		{"invalid context", false, http.StatusInternalServerError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if test.hasContext {
				req = req.WithContext(context.WithValue(req.Context(), consts.CtxChannel, testChannel))
			}

			handler.ServeHTTP(rw, req)
			assert.Equal(t, test.expectedStatus, rw.Code)
			if test.hasContext {
				result := &models.Channel{
					BaseModel: dataful.NewBaseModel(r),
				}

				json.NewDecoder(rw.Result().Body).Decode(result)

				assert.Equal(t, result, testChannel)
			}

		})
	}

}

// func TestHandleUpdateChannel(t *testing.T) {
// 	t.Parallel()
// 	changedChannel := &models.Channel{Joined: false, Prefix: "%", Silenced: true}
// 	b, err := json.Marshal(changedChannel)
// 	require.NoError(t, err)

// 	tests := []struct {
// 		name           string
// 		wantErr        bool
// 		updateErr      bool
// 		hasContext     bool
// 		expectedStatus int
// 	}{
// 		{"successfull update", false, false, true, http.StatusOK},
// 		{"unsuccessfull update", true, true, true, http.StatusInternalServerError},
// 		{"invalid context", true, false, false, http.StatusInternalServerError},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			rw := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(b))
// 			db := new(mocks.Database)
// 			repository := new(mocks.Repository)
// 			db.On("Repository", dataful.DBAmazeful, dataful.CollectionChannel).Return(repository)
// 			if test.hasContext {

// 				channel := models.NewChannel(repository)
// 				channel.SetLoaded(true)
// 				channel.Joined = true
// 				channel.Silenced = false
// 				channel.Prefix = "!"

// 				req = req.WithContext(context.WithValue(req.Context(), consts.CtxChannel, channel))
// 				if test.updateErr {
// 					repository.On("ReplaceOne", req.Context(), mock.Anything, channel, mock.Anything).Return(mongo.ErrNoDocuments)
// 				} else {
// 					repository.On("ReplaceOne", req.Context(), mock.Anything, channel, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
// 						arg := args.Get(2).(*models.Channel)
// 						arg.Joined = changedChannel.Joined
// 						arg.Silenced = changedChannel.Silenced
// 						arg.Prefix = changedChannel.Prefix
// 					})
// 				}
// 			}

// 			channelHandler := NewChannelHandler(&util.Resources{DB: db})
// 			handler := http.HandlerFunc(channelHandler.HandleUpdateChannel)
// 			handler.ServeHTTP(rw, req)

// 			assert.Equal(t, test.expectedStatus, rw.Code)

// 			if !test.wantErr {
// 				received := &models.Channel{}
// 				json.NewDecoder(rw.Result().Body).Decode(received)
// 				assert.Equal(t, changedChannel.Title, received.Title)
// 				assert.Equal(t, changedChannel.Silenced, received.Silenced)
// 				assert.Equal(t, changedChannel.Prefix, received.Prefix)

// 			}

// 		})
// 	}

// }
