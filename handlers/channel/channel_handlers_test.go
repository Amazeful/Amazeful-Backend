package channel

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/models"
	"github.com/Amazeful/Amazeful-Backend/util/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func addTestChannelToContext(channel *models.Channel, req *http.Request) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), consts.CtxChannel, channel))
}

func TestHandleGetChannel(t *testing.T) {
	handler := http.HandlerFunc(HandleGetChannel)

	type args struct {
		channel *models.Channel
	}

	tests := []struct {
		name           string
		wantBool       bool
		expectedStatus int
		args           args
	}{
		{"valid context", true, http.StatusOK, args{&models.Channel{ChannelId: "1"}}},
		{"invalid context", false, http.StatusInternalServerError, args{&models.Channel{ChannelId: "1"}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if test.wantBool {
				req = req.WithContext(context.WithValue(req.Context(), consts.CtxChannel, test.args.channel))
			}

			handler.ServeHTTP(rw, req)
			assert.Equal(t, test.expectedStatus, rw.Code)
			if test.wantBool {
				result := &models.Channel{}
				json.NewDecoder(rw.Result().Body).Decode(result)
				assert.Equal(t, test.args.channel.ChannelId, result.ChannelId)
			}

		})
	}

}

func TestHandleUpdateChannel(t *testing.T) {
	handler := http.HandlerFunc(HandleUpdateChannel)

	changedChannel := &models.Channel{Joined: false, Prefix: "%", Silenced: true}
	b, err := json.Marshal(changedChannel)
	require.NoError(t, err)

	tests := []struct {
		name           string
		wantErr        bool
		updateErr      bool
		hasContext     bool
		expectedStatus int
	}{
		{"successfull update", false, false, true, http.StatusOK},
		{"unsuccessfull update", true, true, true, http.StatusInternalServerError},
		{"invalid context", true, false, false, http.StatusInternalServerError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(b))
			if test.hasContext {
				collection := new(mocks.ICollection)

				channel := models.NewChannel(collection)
				channel.Joined = true
				channel.Silenced = false
				channel.Prefix = "!"

				req = addTestChannelToContext(channel, req)
				if test.updateErr {
					collection.On("ReplaceOne", req.Context(), mock.Anything, channel).Return(&mongo.UpdateResult{MatchedCount: 0}, mongo.ErrNoDocuments)
				} else {
					collection.On("ReplaceOne", req.Context(), mock.Anything, channel).Return(&mongo.UpdateResult{MatchedCount: 1}, nil)
				}
			}
			handler.ServeHTTP(rw, req)

			assert.Equal(t, test.expectedStatus, rw.Code)

			if !test.wantErr {
				received := &models.Channel{}
				json.NewDecoder(rw.Result().Body).Decode(received)
				assert.Equal(t, changedChannel.Title, received.Title)
				assert.Equal(t, changedChannel.Silenced, received.Silenced)
				assert.Equal(t, changedChannel.Prefix, received.Prefix)

			}

		})
	}

}
