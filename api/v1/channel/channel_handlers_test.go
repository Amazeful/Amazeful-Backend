package channel

import (
	"context"
	"testing"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/stretchr/testify/assert"
)

func TestHandleGetChannel(t *testing.T) {
	err := config.LoadConfig()
	assert.NoError(t, err)

	err = util.InitDB(context.Background())
	assert.NoError(t, err)
}

// func TestHandleGetChannel(t *testing.T) {
// 	channelHandler := NewChannelHandler(&util.Resources{})
// 	handler := http.HandlerFunc(channelHandler.HandleGetChannel)

// 	type args struct {
// 		channel *models.Channel
// 	}

// 	tests := []struct {
// 		name           string
// 		wantBool       bool
// 		expectedStatus int
// 		args           args
// 	}{
// 		{"valid context", true, http.StatusOK, args{&models.Channel{BroadcasterName: "Amazeful"}}},
// 		{"invalid context", false, http.StatusInternalServerError, args{&models.Channel{BroadcasterName: "Amazeful"}}},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			rw := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodGet, "/", nil)
// 			if test.wantBool {
// 				req = req.WithContext(context.WithValue(req.Context(), consts.CtxChannel, test.args.channel))
// 			}

// 			handler.ServeHTTP(rw, req)
// 			assert.Equal(t, test.expectedStatus, rw.Code)
// 			if test.wantBool {
// 				result := &models.Channel{}
// 				json.NewDecoder(rw.Result().Body).Decode(result)
// 				assert.Equal(t, test.args.channel.ChannelId, result.ChannelId)
// 			}

// 		})
// 	}

// }

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
