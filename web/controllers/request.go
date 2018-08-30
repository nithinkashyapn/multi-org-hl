package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
)

func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {

	value1 := r.FormValue("key1")
	fmt.Printf("%v\n", value1)

	var data = simplejson.New()

	txid, err := app.Fabric.InvokeHello(value1)
	if err != nil {
		data.Set("success", false)
	} else {
		data.Set("TransactionId", txid)
		data.Set("success", true)
	}

	payload, err := data.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
