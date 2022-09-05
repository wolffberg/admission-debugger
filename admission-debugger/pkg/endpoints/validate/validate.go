package validate

import (
  "log"
  "net/http"
  "net/http/httputil"

	"github.com/wolffberg/admission-debugger/pkg/router"
)

func init() {
  router.BaseRouter.HandleFunc(router.EndpointInitiator{
    Path: "/dump",
    Function: execute,
  })
}

func execute(w http.ResponseWriter, r *http.Request) {
  dump, _ := httputil.DumpRequest(r, true)
  log.Println(string(dump))
}
