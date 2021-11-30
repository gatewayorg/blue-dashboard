package auth_bak

import "net/http"

var (
	auth author
)

func InitGlobal(user author) {
	auth = user
}

func RegisterREST(mux *http.ServeMux) {
	mux.HandleFunc("/user/login", login)
}
