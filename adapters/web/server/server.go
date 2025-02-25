package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Gabriel-Macedogmc/hexagonal-architecture/adapters/web/handler"
	"github.com/Gabriel-Macedogmc/hexagonal-architecture/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w *WebServer) Server() {

	r := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "Log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
