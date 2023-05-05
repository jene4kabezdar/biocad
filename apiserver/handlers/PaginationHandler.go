package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jene4kabezdar/biocad/internal/app/model"
	"github.com/jene4kabezdar/biocad/internal/app/store"
)

func PaginatonHandler(w http.ResponseWriter, r *http.Request) {
	var store store.Store
	defer store.Close()
	store.ConfigureStore()
	store.Open()

	w.Header().Set("Access-Control-Allow-Origin", "*")

	number, err := strconv.Atoi(mux.Vars(r)["number"])
	if err != nil {
		log.Fatal(err)
	}

	messages, err := model.GetMessages(store, number)
	if err != nil {
		log.Fatal(err)
	}

	response, err := json.Marshal(messages)
	if err != nil {
		log.Fatal(err)
	}
	
	w.Write(response)
}