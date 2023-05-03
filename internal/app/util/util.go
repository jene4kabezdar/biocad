package util

import (
	"time"

	"github.com/jene4kabezdar/biocad/internal/app/files"
	"github.com/jene4kabezdar/biocad/internal/app/store"
)

func HandleError(err error, store store.Store) {
	if err != nil {
		insertError(err.Error(), store)
		files.WriteErrorLog(err.Error())
	}
}

func InArrStr(str string, strArr []string) bool {
	for _, s := range strArr {
		if s == str {
			return true
		}
	}
	return false
}

func insertError(message string, store store.Store) {

	store.DB.QueryRow(
		`INSERT INTO errors (error_text, error_date) VALUES ($1, $2) RETURNING n`,
		message, time.Now().Format("02-01-2006 15:04"),
	)
}
