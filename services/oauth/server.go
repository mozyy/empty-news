package oauth

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/mozyy/empty-news/services/user"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-session/session"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dumpvar bool = true
	portvar int  = 9096
)

var TokenStore oauth2.TokenStore

func New() {
	flag.Parse()
	if dumpvar {
		log.Println("Dumping requests")
	}

	manager := manage.NewDefaultManager()

	// use mysql token store
	// manager.MapTokenStorage(NewStoreToken())
	tokenStore, err := store.NewMemoryTokenStore()
	if err != nil {
		panic(err)
	}
	TokenStore = tokenStore
	manager.MapTokenStorage(tokenStore)

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	// manager.MapAccessGenerate(generates.NewAccessGenerate())

	manager.MapClientStorage(NewStoreClient())

	srv := server.NewServer(server.NewConfig(), manager)

	userStore := user.NewUserStore()

	srv.SetPasswordAuthorizationHandler(func(ctx context.Context, mobile, password string) (ID string, err error) {
		user, err := userStore.Get(mobile, password)
		if err != nil {
			return
		}
		return strconv.FormatUint(uint64(user.ID), 10), nil
	})

	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		re = errors.NewResponse(err, 400)
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/oauth/authorize", func(w http.ResponseWriter, r *http.Request) {
		if dumpvar {
			dumpRequest(os.Stdout, "authorize", r)
		}

		log.Println("resave oauth/authorize")

		err = srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		if dumpvar {
			_ = dumpRequest(os.Stdout, "token", r) // Ignore the error
		}

		fmt.Println(123123)

		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Printf("Server is running at %d port.\n", portvar)
	go func() {
		err = http.ListenAndServe(fmt.Sprintf(":%d", portvar), nil)
		if err != nil {
			log.Println("oauth server err:", err)
		}
	}()
}

func dumpRequest(writer io.Writer, header string, r *http.Request) error {
	data, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}
	writer.Write([]byte("\n" + header + ": \n"))
	writer.Write(data)
	return nil
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	if dumpvar {
		_ = dumpRequest(os.Stdout, "userAuthorizeHandler", r) // Ignore the error
	}
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

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
}
