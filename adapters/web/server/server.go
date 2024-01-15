package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/glaydsonp/go-hexagonal/adapters/web/handler"
	"github.com/glaydsonp/go-hexagonal/application"
	"github.com/gorilla/mux"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {

	router := mux.NewRouter()
	negroni := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(router, *negroni, w.Service)

	http.Handle("/", router)

	server := &http.Server{
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
