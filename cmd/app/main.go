package main

import (
	"log"

	"github.com/jene4kabezdar/biocad/internal/app/store"
	"github.com/jene4kabezdar/biocad/internal/app/watcher"
)

func main() {
	var store store.Store
	defer store.Close()
	store.ConfigureStore()
	log.Println("opening store")
	store.Open()
	
	watcher.Start(store)
}
