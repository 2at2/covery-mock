package handler

import (
	"net/http"
	"github.com/mono83/xray"
)

type NodeNames struct {
	Ray xray.Ray
}

func (h NodeNames) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Ray.Info("NodeNames")
	sendResponse(w, map[string]interface{}{})
}
