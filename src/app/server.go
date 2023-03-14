package main

import (
	"fmt"
	"go_template/src/logger"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type (
	Server struct {
		logger    logger.LoggingI
		prop      *properties.Properties
		processor *Processor
	}
)

func NewServer(logger logger.LoggingI, prop *properties.Properties, proc *Processor) *Server {
	return &Server{logger: logger, prop: prop, processor: proc}
}

func (s *Server) handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/products", s.processor.returnAllProducts)
	myRouter.HandleFunc("/product/{id}", s.processor.returnSingleProduct)
	myRouter.HandleFunc("/product", s.processor.createProduct).Methods("POST")
	myRouter.HandleFunc("/articles", s.processor.returnAllArticles)
	myRouter.HandleFunc("/article", s.processor.createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", s.processor.deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", s.processor.returnSingleArticle)
	myRouter.Path("/metrics").Handler(promhttp.Handler())

	s.logger.Fatal(http.ListenAndServe(":8080", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
