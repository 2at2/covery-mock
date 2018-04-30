package handler

import (
	"net/http"
	"github.com/mono83/xray"
)

type TrustListCheck struct {
	Ray xray.Ray
}

func (h TrustListCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Ray.Info("TrustListCheck")
	sendResponse(w, map[string]interface{}{})
}
