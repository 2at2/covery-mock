package handler

import (
	"net/http"
	"encoding/json"
)

func sendResponse(w http.ResponseWriter, entity interface{}) {
	if err, ok := entity.(error); ok {
		sendErrorResponse(w, err)
	} else {
		bts, err := json.Marshal(entity)

		if err != nil {
			sendErrorResponse(w, err)
		} else {
			w.WriteHeader(200)
			w.Write(bts)
		}
	}
}

func sendErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Add("X-Maxwell-Status", "Exception")
	w.Header().Add("X-Maxwell-Error-Message", err.Error())
	w.WriteHeader(500)
}
