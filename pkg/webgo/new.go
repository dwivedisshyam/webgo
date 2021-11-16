package webgo

import (
	"net/http"
	"os"

	"github.com/dwivedisshyam/webgo/pkg/log"
	"github.com/dwivedisshyam/webgo/pkg/webgo/config"
	"github.com/dwivedisshyam/webgo/pkg/webgo/datastore"
)

type WebGo struct {
	datastore.Datastore
	Config Config
	Logger log.Logger
	Server *Server
}

func New() (w *WebGo) {
	var (
		logger       = log.NewLogger()
		configFolder string
	)

	if _, err := os.Stat("./configs"); err == nil {
		configFolder = "./configs"
	} else if _, err := os.Stat("../configs"); err == nil {
		configFolder = "../configs"
	} else {
		configFolder = "../../configs"
	}

	return NewWithConfig(config.NewGoDotEnvProvider(logger, configFolder))
}

func NewWithConfig(c Config) *WebGo {
	w := &WebGo{
		Logger: log.NewLogger(),
		Config: c,
	}

	w.Server = NewServer(c, w)

	initilizeDatastores(w)

	return w
}

func (w *WebGo) Get(path string, handler Handler) {
	w.Server.Router.Route(http.MethodGet, path, handler)
}

func (w *WebGo) Post(path string, handler Handler) {
	w.Server.Router.Route(http.MethodPost, path, handler)
}

func (w *WebGo) Put(path string, handler Handler) {
	w.Server.Router.Route(http.MethodPut, path, handler)
}

func (w *WebGo) Delete(path string, handler Handler) {
	w.Server.Router.Route(http.MethodDelete, path, handler)
}

func initilizeDatastores(w *WebGo) {
	initilizeSQL(w)
}

func initilizeSQL(w *WebGo) {
	dc := &datastore.SQLConfig{
		Host:     w.Config.Get("DB_HOST"),
		Port:     w.Config.Get("DB_PORT"),
		Name:     w.Config.Get("DB_NAME"),
		User:     w.Config.Get("DB_USER"),
		Password: w.Config.Get("DB_PASS"),
		Dialect:  w.Config.Get("DB_DIALECT"),
	}

	if dc.Host == "" || dc.Dialect == "" {
		return
	}

	db := datastore.NewSQL(w.Logger, dc)
	w.Datastore.SetDB(db)
}
