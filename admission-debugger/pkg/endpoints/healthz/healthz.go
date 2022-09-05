package validate

import (
  "net/http"

	"github.com/wolffberg/admission-debugger/pkg/router"
)

func init() {
  router.BaseRouter.HandleFunc(router.EndpointInitiator{
    Path: "/healthz",
    Function: execute,
  })
}

func execute(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Works!"))
}
