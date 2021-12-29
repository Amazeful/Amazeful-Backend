package channel

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/models"
	"github.com/stretchr/testify/assert"
)

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

		})
	}

}
