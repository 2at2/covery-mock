package handler

import (
	"net/http"
	"github.com/mono83/xray"
)

type Ping struct {
	Ray xray.Ray
}

func (h Ping) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Ray.Info("Ping")

	res := map[string]interface{}{
		"access": map[string]interface{}{
			"decision": true,
		},
	}

	sendResponse(w, res)
}
