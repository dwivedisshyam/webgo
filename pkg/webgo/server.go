package webgo

import (
	ctx "context"
	"fmt"

	"net/http"
	"sync"

	"github.com/dwivedisshyam/webgo/pkg/log"
	middeware "github.com/dwivedisshyam/webgo/pkg/middleware"
	"github.com/dwivedisshyam/webgo/pkg/webgo/types"
)

type contextKey int

const webGoContextKey contextKey = 1

type Server struct {
	contextPool sync.Pool

	Router Router
	HTTP   *HTTP
	done   chan bool
}

type HTTP struct {
	Port string
}

func NewServer(config Config, w *WebGo) *Server {
	srv := &Server{
		Router: NewRouter(),
		done:   make(chan bool),
		HTTP:   new(HTTP),
	}

	srv.contextPool.New = func() interface{} {
		return NewContext(nil, nil, w)
	}

	srv.HTTP.Port = config.GetOrDefault("HTTP_PORT", "8000")

	return srv
}

func (s *Server) Start(logger log.Logger) {
	s.Router.Use(s.ctxInjector)
	s.Router.Use(middeware.Logging(logger))
	s.Router.CatchAllRoute(func(c *Context) (i interface{}, err error) {
		return nil, &types.Error{StatusCode: http.StatusNotFound, Reason: fmt.Sprintf("Route %s %s not found", c.req.Method, c.req.URL)}
	})

	go func() {
		addr := ":" + s.HTTP.Port
		logger.Infof("starting http server at %s", addr)

		srv := &http.Server{
			Addr:    addr,
			Handler: s.Router,
		}

		err := srv.ListenAndServe()
		if err != nil {
			s.done <- true
		}
	}()

	<-s.done
	logger.Info("Server received on done channel. Stopping")
}

// ctxInjector injects *Context variable into every request using a middleware
func (s *Server) ctxInjector(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := s.contextPool.Get().(*Context)
		c.reset(w, r)
		c.Context = r.Context()
		*r = *r.WithContext(ctx.WithValue(c, webGoContextKey, c))

		inner.ServeHTTP(w, r)

		s.contextPool.Put(c)
	})
}
