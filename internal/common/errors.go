package common

import (
	"log"
	"net/http"
)

func HandleErrorPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HandleErrorReturn(err error) {
	if err != nil {
		return
	}
}

func HandleErrorWithParam(err error, str string) {
	if err != nil {
		log.Fatalf(str+": %v", err)
	}
}

func HandleInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
