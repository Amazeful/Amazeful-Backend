package middlewares

// func TestChannelFromSession(t *testing.T) {
// 	t.Parallel()
// 	type args struct {
// 		session *models.Session
// 		channel *models.Channel
// 	}

// 	tests := []struct {
// 		name           string
// 		existsInDB     bool
// 		expectedStatus int
// 		args           args
// 	}{
// 		{"no session", false, http.StatusUnauthorized, args{nil, nil}},
// 		{"not in db", false, http.StatusInternalServerError, args{&models.Session{SelectedChannel: primitive.ObjectID{}, SessionId: "1"}, &models.Channel{ChannelId: "TestChannelFromSession1"}}},
// 		{"valid", true, http.StatusOK, args{&models.Session{SelectedChannel: primitive.ObjectID{}, SessionId: "1"}, &models.Channel{ChannelId: "TestChannelFromSession1"}}},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			rw := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodGet, "/", nil)
// 			if test.args.session != nil {
// 				req = req.WithContext(context.WithValue(req.Context(), consts.CtxSession, test.args.session))
// 			}
// 			handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 				channel, ok := r.Context().Value(consts.CtxChannel).(*models.Channel)
// 				if !ok {
// 					util.WriteError(rw, consts.ErrNoContextValue, http.StatusInternalServerError, consts.ErrStrResourceDNE)
// 					return
// 				}
// 				util.WriteResponse(rw, util.Response{
// 					Status: http.StatusOK,
// 					Data:   channel,
// 				})
// 			})
// 			if test.existsInDB {
// 				test.args.channel.R = util.GetDB().Repository(dataful.DBAmazeful, dataful.CollectionChannel)
// 				err := test.args.channel.Create(req.Context())
// 				assert.NoError(t, err)
// 			}

// 			testHandler := ChannelFromSession(handler)
// 			testHandler.ServeHTTP(rw, req)
// 			assert.Equal(t, test.expectedStatus, rw.Code)

// 			if test.existsInDB {
// 				result := rw.Result()
// 				returnedChannel := &models.Channel{}
// 				err := json.NewDecoder(result.Body).Decode(returnedChannel)
// 				assert.NoError(t, err)

// 				assert.Equal(t, test.args.channel.ChannelId, returnedChannel.ChannelId)

// 				err = test.args.channel.Delete(req.Context())
// 				assert.NoError(t, err)

// 			}
// 		})
// 	}
// }
