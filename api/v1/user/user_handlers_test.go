package user

// func TestHandleGetUser(t *testing.T) {
// 	handler := http.HandlerFunc(HandleGetUser)

// 	type args struct {
// 		user *models.User
// 	}

// 	tests := []struct {
// 		name           string
// 		wantBool       bool
// 		expectedStatus int
// 		args           args
// 	}{
// 		{"valid context", true, http.StatusOK, args{&models.User{UserID: "1"}}},
// 		{"invalid context", false, http.StatusInternalServerError, args{&models.User{UserID: "1"}}},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			rw := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodGet, "/", nil)
// 			if test.wantBool {
// 				req = req.WithContext(context.WithValue(req.Context(), consts.CtxUser, test.args.user))
// 			}

// 			handler.ServeHTTP(rw, req)
// 			assert.Equal(t, test.expectedStatus, rw.Code)

// 		})
// 	}

// }
