package main

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := execute(); err != nil {
		log.Errorf("%+v", err)
		os.Exit(1)
	}
}

func execute() error {
	addr := ":" + os.Getenv("PORT")

	c := Controller{}
	r := httprouter.New()
	r.HandlerFunc(http.MethodGet, `/`, c.IndexHandler)
	r.HandlerFunc(http.MethodGet, `/healthz`, c.HealthcheckHandler)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Infof("listening on: %s", addr)
	return srv.ListenAndServe()
}
