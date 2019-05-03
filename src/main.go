package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func runProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response := []byte("Hello Universe! (because World is too small...)")
	w.Write(response)
	return
}

//NewRouter api router
func newRouter() *mux.Router {
	apiRouter := mux.NewRouter().StrictSlash(true) //mainRouter.PathPrefix("/api").Subrouter().StrictSlash(true)
	//Create Routes
	apiRouter.HandleFunc("/", runProgram).Methods("GET")
	return apiRouter
}
func main() {
	// launch server in port 8500
	log.Fatal(http.ListenAndServe(":8500", newRouter()))
}
