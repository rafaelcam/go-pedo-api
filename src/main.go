package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/apex/log"
	"github.com/rafaelcam/go-pedo-api/src/controllers"
	"time"
)

func main() {
	router := createRouter()

	srv := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
	}

	log.Info("Running server!")
	if err := srv.ListenAndServe(); err != nil {
		log.WithError(err).Fatal("Server fail")
	}
}

func createRouter() *mux.Router {
	router := mux.NewRouter()
	router.Methods("GET").Path("/health").HandlerFunc(controllers.AppHealth)

	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.Methods("GET").Path("/pedo").HandlerFunc(controllers.GetPeDo)

	return router
}