package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cubereum/mew/web/controllers"
	"github.com/gorilla/mux"
)

func Serve(app *controllers.Application) {

	router := mux.NewRouter()

	// Test API
	router.HandleFunc("/test", testApi).Methods("GET")

	// Main APIs
	router.HandleFunc("/home", app.HomeHandler).Methods("GET")
	router.HandleFunc("/request", app.RequestHandler).Methods("POST")

	fmt.Println("Listening on http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func testApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test Successful!")
}
