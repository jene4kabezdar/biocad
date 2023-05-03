package main

import (
	"github.com/jene4kabezdar/biocad/internal/app/files"
	"github.com/jene4kabezdar/biocad/internal/app/model"
	"github.com/jene4kabezdar/biocad/internal/app/store"
	"github.com/jene4kabezdar/biocad/internal/app/util"
)

func main() {
	var store store.Store
	defer store.Close()
	store.ConfigureStore()
	store.Open()

	rows, err := files.ParseStartData()
	util.HandleError(err, store)

	messages, err := model.CreateMessagesByRows(rows)
	util.HandleError(err, store)

	for _, message := range messages {
		_, err := message.Add(store)
		util.HandleError(err, store)
	}
}
