package sessions

//setCookie
//var sessions map[string]string

// sessions["session_id"] = "user_id"

// userId, ok := sessions["some_session_id"]
// if !ok {
//   // Map doesn't have that session id in it
//}
//method on request Cookie
//var dbSessions = map[string]string{} //

//dbSessions["session_id"]
// func RequestIDMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		requestID := r.Header.Get("X-Request-ID")
// 		if requestID == "" {
// 			r.Header.Set("X-Request-ID", uuid.NewV4().String())
		}

// // SessionMiddleware is used in another file
// func SessionMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		c, err := r.Cookie("my-cookie")
// 		if err != nil {
// 			id := uuid.NewV4()
// 			c = &http.Cookie{
// 				Name:     "session",
// 				Value:    id.String(),
// 				HttpOnly: true,
// 			}
// 			http.SetCookie(w, c)
// 		}
// 	})
// }
