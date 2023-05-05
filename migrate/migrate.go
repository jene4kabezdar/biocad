package main

import (
	"log"

	"github.com/jene4kabezdar/biocad/internal/app/files"
	"github.com/jene4kabezdar/biocad/internal/app/model"
	"github.com/jene4kabezdar/biocad/internal/app/store"
	"github.com/jene4kabezdar/biocad/internal/app/util"
)

func main() {
	var store store.Store
	defer store.Close()
	store.ConfigureStore()
	log.Println("opening store")
	store.Open()

	log.Println("parse start data")
	rows, err := files.ParseStartData()
	util.HandleError(err, store)

	messages, err := model.CreateMessagesByRows(rows)
	util.HandleError(err, store)

	log.Println("inserting start data")
	for _, message := range messages {
		_, err := message.Add(store)
		util.HandleError(err, store)
	}
	log.Println("migrate successful")
}
