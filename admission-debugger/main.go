package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/wolffberg/admission-debugger/pkg/router"
	"github.com/wolffberg/admission-debugger/pkg/cert"
  
	_ "github.com/wolffberg/admission-debugger/pkg/endpoints/validate"
	_ "github.com/wolffberg/admission-debugger/pkg/endpoints/healthz"
)

func main() {
  cert.CreateCertFiles("")
  cert.CreateCertSecret("")

	go func() {
		if err := router.BaseRouter.Server.ListenAndServeTLS("tls.crt", "tls.key"); err != nil {
			fmt.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	var wait time.Duration
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	router.BaseRouter.Server.Shutdown(ctx)
	fmt.Println("shutting down")
	os.Exit(0)
}
