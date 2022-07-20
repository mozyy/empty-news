package oauth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-session/session"
)

func NewServerHttp(port int) {
	srv := newSrv()
	srv.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		store, err := session.Start(r.Context(), w, r)
		if err != nil {
			return
		}

		// TODO: LoggedInUserID 哪里设置的 ???
		uid, ok := store.Get("LoggedInUserID")
		if !ok {
			if r.Form == nil {
				r.ParseForm()
			}

			store.Set("ReturnUri", r.Form)
			store.Save()

			w.Header().Set("Location", r.Header.Get("x-envoy-origin")+"/login")
			w.WriteHeader(http.StatusFound)
			return
		}

		userID = uid.(string)
		store.Delete("LoggedInUserID")
		store.Save()
		return
	})
	http.HandleFunc("/oauth/authorize", func(w http.ResponseWriter, r *http.Request) {
		log.Println("resave oauth/authorize")

		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		log.Println("resave oauth/token")

		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Println("oauth server err:", err)
	}
}
