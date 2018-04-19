package sessions

import (
	"net/http"
	"github.com/satori/go.uuid"
)

//setCookie
// var sessions map[string]string

// sessions["session_id"] = "user_id"

// userId, ok := sessions["some_session_id"]
// if !ok {
//   // Map doesn't have that session id in it
//}
//method on request Cookie
var dbSessions = map[string]string{}//

dbSessions["session_id"]
func SetCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie",
		Value: "013d27409g2837409y837a38943i00"
	})
}

func SetCookies(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		id := uuid.NewV4();
		c = &http.Cookie{
			Name: "session",
			Value: id.String(),
			httpOnly: true,
		}
		http.SetCookie(w, c)
	}


}