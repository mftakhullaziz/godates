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
		log.Fatalf("%s: %v", str, err)
	}
}

func HandleInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleErrorDefault(err error) error {
	if err != nil {
		return err
	}
	return nil
}
