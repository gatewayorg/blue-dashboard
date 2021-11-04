package auth_bak

import (
	"context"
	"encoding/json"
	"github.com/gatewayorg/blue-dashboard/pkg/jwt"
	"net/http"
)

type loginResp struct {
	AccessToken string `json:"access_token"`
}

func (l *loginResp) Bytes() []byte {
	bytes, _ := json.Marshal(l)
	return bytes
}

func (l *loginResp) Write(w http.ResponseWriter) {
	w.Write(l.Bytes())
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	username := r.Form.Get("username")
	passwd := r.Form.Get("passwd")
	var (
		id  uint64
		err error
	)
	if username != "" && passwd != "" {
		id, err = auth.Authenticate(context.Background(), username, passwd)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token, err := jwt.Sign.Sign(id, username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		l := loginResp{AccessToken: token}

		l.Write(w)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}
