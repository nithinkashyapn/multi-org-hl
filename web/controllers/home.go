package controllers

import (
	"fmt"
	"log"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Request is %v\n", r)

	var data = simplejson.New()

	helloValue, err := app.Fabric.QueryHello()
	if err != nil {
		data.Set("success", false)
	} else {
		data.Set("helloValue", helloValue)
		data.Set("success", true)
	}

	payload, err := data.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
