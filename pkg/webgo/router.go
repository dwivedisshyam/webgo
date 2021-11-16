package webgo

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	http.Handler

	Use(...mux.MiddlewareFunc)
	Route(string, string, Handler)
	CatchAllRoute(h Handler)
}

type router struct {
	mux.Router
}

func (r *router) Use(m ...mux.MiddlewareFunc) {
	r.Router.Use(m...)
}

func (r *router) Route(method, path string, handler Handler) {
	r.HandleFunc(path, handler.ServeHTTP).Methods(method)
}

func (r *router) CatchAllRoute(h Handler) {
	r.Router.PathPrefix("/").Handler(h)
}

func NewRouter() Router {
	muxRouter := mux.NewRouter().StrictSlash(false)
	r := router{Router: *muxRouter}

	return &r
}
