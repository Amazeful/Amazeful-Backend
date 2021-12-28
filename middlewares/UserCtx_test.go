package middlewares

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
	"go.mongodb.org/mongo-driver/mongo"
)

func TestUserCtx(t *testing.T) {

	type args struct {
		session *models.Session
		user    *models.User
	}

	tests := []struct {
		name           string
		existsInDB     bool
		expectedStatus int
		args           args
	}{
		{"no session", false, http.StatusUnauthorized, args{nil, nil}},
		{"not in db", false, http.StatusInternalServerError, args{&models.Session{SelectedChannel: primitive.ObjectID{}, SessionId: "1"}, &models.User{UserID: "123"}}},
		{"valid", true, http.StatusOK, args{&models.Session{SelectedChannel: primitive.ObjectID{}, SessionId: "1"}, &models.User{UserID: "123"}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if test.args.session != nil {
				req = req.WithContext(context.WithValue(req.Context(), consts.CtxSession, test.args.session))
			}
			handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				user, ok := r.Context().Value(consts.CtxUser).(*models.User)
				if !ok {
					util.WriteError(rw, consts.ErrNoContextValue, http.StatusInternalServerError, consts.ErrStrResourceDNE)
					return
				}

				util.WriteResponse(rw, util.Response{
					Status: http.StatusOK,
					Data:   user,
				})
			})
			mockDB := new(mocks.IDB)
			mockCollection := new(mocks.ICollection)
			mockedResult := new(mocks.ISingleResult)
			mockDB.On("Collection", consts.CollectionUser).Return(mockCollection)

			if test.existsInDB {
				mockedResult.On("Decode", models.NewUser(mockCollection)).Return(nil).Run(func(args mock.Arguments) {
					arg := args.Get(0).(*models.User)
					arg.UserID = test.args.user.UserID
				})
				mockCollection.On("FindOne", req.Context(), bson.M{"_id": test.args.session.User}, mock.Anything).Return(mockedResult)
			} else {
				mockedResult.On("Decode", mock.Anything).Return(mongo.ErrNoDocuments)
				mockCollection.On("FindOne", req.Context(), mock.Anything, mock.Anything).Return(mockedResult)
			}

			util.SetDB(mockDB)

			testHandler := UserCtx(handler)
			testHandler.ServeHTTP(rw, req)
			assert.Equal(t, test.expectedStatus, rw.Code)

			if test.existsInDB {
				result := rw.Result()
				returnedUser := &models.User{}
				err := json.NewDecoder(result.Body).Decode(returnedUser)
				assert.NoError(t, err)

				assert.Equal(t, returnedUser.UserID, test.args.user.UserID)
			}

		})
	}
}
