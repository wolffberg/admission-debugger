package router

import (
  "net/http"
  "time"
  "fmt"

	"github.com/gorilla/mux"
)

type Router struct {
  Router *mux.Router
  Server *http.Server
  Paths map[string]int
}

type EndpointInitiator struct {
  Path     string
  Function func(w http.ResponseWriter, r *http.Request)
}

var BaseRouter *Router

func init() {
  BaseRouter = &Router{
    Router: mux.NewRouter(),
    Paths:  make(map[string]int),
  }

	http.Handle("/", BaseRouter.Router)

	BaseRouter.Server = &http.Server{
		Addr: "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      BaseRouter.Router,
	}
}

func (r *Router)HandleFunc(e EndpointInitiator) error {
  if _, ok := r.Paths[e.Path]; ok {
    panic(fmt.Sprintf("Path '%s' already exists!", e.Path))
  }

  r.Paths[e.Path] = 0
  r.Router.HandleFunc(e.Path, e.Function)
  return nil
}
