package middlewares

// func UserFromSession(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
// 		session, ok := req.Context().Value(consts.CtxSession).(*models.Session)
// 		if !ok {
// 			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 			return
// 		}

// 		r := util.GetDB().Repository(dataful.DBAmazeful, dataful.CollectionUser)
// 		user := models.NewUser(r)
// 		err := user.FindBylId(req.Context(), session.User)
// 		if err != nil {
// 			util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
// 			return
// 		}
// 		if !user.Loaded() {
// 			err = fmt.Errorf("selected userid %s in session does not exist in DB", session.User.String())
// 			util.WriteError(rw, err, http.StatusInternalServerError, consts.ErrStrResourceDNE)
// 			return
// 		}
// 		ctx := context.WithValue(req.Context(), consts.CtxUser, user)
// 		next.ServeHTTP(rw, req.WithContext(ctx))
// 	})
// }
