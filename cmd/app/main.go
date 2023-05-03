package main

import (
	"github.com/jene4kabezdar/biocad/internal/app/store"
	"github.com/jene4kabezdar/biocad/internal/app/watcher"
)

func main() {
	var store store.Store
	defer store.Close()
	store.ConfigureStore()
	store.Open()

	watcher.Start(store)
}
