package api

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"github.com/haashemi/BotManagerBot/config"
	"github.com/haashemi/BotManagerBot/manager"
)

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
		// TODO: write a proper error, not this empty object.
		// TODO: cache the error's bytes somewhere and don't marshal every time.
		if !api.manager.IsBotExists(botToken) {
			invalidData, _ := json.Marshal(map[string]any{})

			w.WriteHeader(http.StatusForbidden)
			w.Write(invalidData)
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
