package main

import (
  "testing"

	"github.com/wolffberg/admission-debugger/pkg/router"
)

// Dummy function to initialize all endpoints 
// and check for dublicate routes.
func TestUniqueRoute(t *testing.T) {
  _ = router.BaseRouter.Paths
}
