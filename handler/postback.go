package handler

import (
	"net/http"
	"github.com/mono83/xray"
)

type Postback struct {
	Ray xray.Ray
}

func (h Postback) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Ray.Info("Postback")
	sendResponse(w, map[string]interface{}{})
}
