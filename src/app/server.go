package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type (
	Server struct {
		logger    LoggingI
		prop      *properties.Properties
		processor *Processor
	}
)

func NewServer(logger LoggingI, prop *properties.Properties, proc *Processor) *Server {
	return &Server{logger: logger, prop: prop, processor: proc}
}

func (s *Server) handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(prometheusMiddleware)
	myRouter.Path("/metrics").Handler(promhttp.Handler())
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/products", s.processor.returnAllProducts)
	myRouter.HandleFunc("/product/{id}", s.processor.returnSingleProduct)
	myRouter.HandleFunc("/product", s.processor.createProduct).Methods("POST")
	myRouter.HandleFunc("/articles", s.processor.returnAllArticles)
	myRouter.HandleFunc("/article", s.processor.createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", s.processor.deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", s.processor.returnSingleArticle)
	//Http time out config,
	srv := &http.Server{
		Handler: myRouter,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	s.logger.Fatal(srv.ListenAndServe())
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
