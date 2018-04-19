package sessions

import (
	"net/http"
	"github.com/satori/go.uuid"
)

//setCookie

//method on request Cookie

func SetCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie",
		Value: "013d27409g2837409y837a38943i00"
	})
}

func Read(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintln(w, "Your Cookie:", c)
}