package main

import (
	"github.com/jene4kabezdar/biocad/internal/app/files"
	"github.com/jene4kabezdar/biocad/internal/app/model"
	"github.com/jene4kabezdar/biocad/internal/app/store"
	"github.com/jene4kabezdar/biocad/internal/app/util"
)

func main() {
	rows, err := files.ParseStartData()
	util.HandleError(err)

	var store store.Store
	defer store.Close()
	store.ConfigureStore()
	store.Open()

	messages, err := model.CreateMessagesByRows(rows)
	util.HandleError(err)

	for _, message := range messages {
		_, err := message.Add(store)
		util.HandleError(err)
	}
}
