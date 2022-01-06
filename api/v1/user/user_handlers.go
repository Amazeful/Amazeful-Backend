package user

// func HandleGetUser(rw http.ResponseWriter, req *http.Request) {
// 	user, ok := req.Context().Value(consts.CtxUser).(*models.User)
// 	if !ok {
// 		util.WriteError(rw, consts.ErrNoContextValue, http.StatusInternalServerError, consts.ErrStrResourceDNE)
// 		return
// 	}

// 	util.WriteResponse(rw, util.Response{Status: http.StatusOK, Data: user})
// }
