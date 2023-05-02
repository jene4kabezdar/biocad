package util

import "log"

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
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