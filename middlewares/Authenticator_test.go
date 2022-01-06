package middlewares

// func TestAuthenticator(t *testing.T) {
// 	t.Parallel()
// 	type args struct {
// 		sessionId string
// 		expiry    time.Time
// 	}

// 	tests := []struct {
// 		name           string
// 		authenticated  bool
// 		sessionExists  bool
// 		expectedStatus int
// 		args           args
// 	}{
// 		{"valid session", true, true, http.StatusOK, args{sessionId: "TestAuthenticator1", expiry: time.Now().Add(time.Minute)}},
// 		{"invalid session", true, false, http.StatusUnauthorized, args{sessionId: "TestAuthenticator2", expiry: time.Now().Add(time.Minute)}},
// 		{"invalid token", false, true, http.StatusUnauthorized, args{sessionId: "TestAuthenticator3", expiry: time.Now().Add(time.Minute)}},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {

// 			rw := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodGet, "/", nil)
// 			handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 				util.WriteResponse(rw, util.Response{
// 					Status: http.StatusOK,
// 					Data:   nil,
// 				})
// 			})

// 			if test.authenticated {
// 				jwt := models.NewJWT([]byte(config.GetConfig().JwtSignKey), jwa.HS256)
// 				tokenString, err := jwt.Encode(test.args.sessionId, test.args.expiry)
// 				assert.NoError(t, err)
// 				req.AddCookie(&http.Cookie{Name: auth.JWTCookieName, Value: tokenString})
// 			}

// 			if test.sessionExists {
// 				session := models.NewSession(util.GetRedis())
// 				session.SessionId = test.args.sessionId
// 				err := session.SetSession(req.Context(), time.Until(test.args.expiry))
// 				assert.NoError(t, err)
// 			}

// 			authHandler := Authenticator(handler)
// 			authHandler.ServeHTTP(rw, req)
// 			assert.Equal(t, test.expectedStatus, rw.Code)
// 		})
// 	}
// }
