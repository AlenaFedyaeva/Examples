package main

import (
	"fmt"
	"net/http"

	_ "swagexample/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// @host localhost:8989
// @BasePath /tmp
// @query.collection.format multi
func main() {
	fmt.Println("srv")
	r := mux.NewRouter()

	r.HandleFunc("/tmp", tmpHandle).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8989/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	)).Methods(http.MethodGet)

	// r.HandleFunc("/sw", tmpHandle).Methods(sw.)

	http.ListenAndServe(":8989", r)
}

// @Summary tmpHandle
// @ID create-todo
// @Produce json
func tmpHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've requested:  ", "tmp")
}
