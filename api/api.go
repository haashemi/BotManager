package api

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"github.com/haashemi/BotManager/config"
	"github.com/haashemi/BotManager/manager"
)

var UnauthorizedResponse = []byte(`{"ok":false,"error_code":403,"description":"Token is not authorized in the BotManager"}`)

type RunFunc func() error

type API struct {
	manager *manager.Manager
}

// NewAPI initializes an API instance and returns the ListerAndServe method
func NewAPI(config *config.Config, manager *manager.Manager) (RunFunc, error) {
	target, err := url.Parse(config.TelegramBotAPI.Host)
	if err != nil {
		return nil, err
	}

	api := &API{manager: manager}

	r := mux.NewRouter()
	r.HandleFunc("/{token}/{method}", api.methodHandler(target))
	r.HandleFunc("/file/{token}/{dir}/{file}", api.fileHandler(config.TelegramBotAPI.Dir))

	srv := &http.Server{Handler: r, Addr: config.API.Addr}
	return func() error { return srv.ListenAndServe() }, nil
}

// methodHandler handles api method calls
func (api *API) methodHandler(target *url.URL) http.HandlerFunc {
	prx := httputil.NewSingleHostReverseProxy(target)

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		botToken := strings.TrimPrefix(vars["token"], "bot")

		// Don't proxy the request if bot is not whitelisted.
		//
		// TODO: send an alert via the bot to the admins.
		if !api.manager.IsBotExists(botToken) {
			w.WriteHeader(http.StatusForbidden)
			w.Write(UnauthorizedResponse)
			return
		}

		r.Host = target.Host
		prx.ServeHTTP(w, r)
	}
}

// fileHandler handles file downloads
func (api *API) fileHandler(dir string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(dir))

	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/file/bot")
		fs.ServeHTTP(w, r)
	}
}
