package main

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"time"

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

	entries, err := os.ReadDir("indata")
	util.HandleError(err, store)

	amountEntries := len(entries)
	names := make([]string, amountEntries)
	for i, e := range entries {
		names[i] = e.Name()
	}

	for {
		time.Sleep(15 * time.Second)
		log.Println("reading data...")
		entries, err = os.ReadDir("indata")
		util.HandleError(err, store)
		if len(entries) > amountEntries {
			for _, e := range entries {
				if !util.InArrStr(e.Name(), names) {
					names = append(names, e.Name())
					f, err := os.Open("indata/" + e.Name())
					util.HandleError(err, store)

					reader := csv.NewReader(f)
					reader.FieldsPerRecord = -1
					reader.Comma = '\t'

					rows, err := reader.ReadAll()
					util.HandleError(err, store)

					f.Close()

					messages, err := model.CreateMessagesByRows(files.ProcessingData(rows))
					util.HandleError(err, store)

					for _, message := range messages {
						if message.Unit_guid == "" {
							util.HandleError(errors.New("not enough unit_guid"), store)
						}
						_, err := message.Add(store)
						util.HandleError(err, store)
					}
				}
			}
			amountEntries = len(entries)
		}
	}
}
