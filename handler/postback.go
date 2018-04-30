package handler

import (
	"net/http"
	"github.com/mono83/xray"
	"io/ioutil"
	"github.com/mono83/xray/args"
)

type Postback struct {
	Ray xray.Ray
}

func (h Postback) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Ray.Info("Postback")

	bts, err := ioutil.ReadAll(r.Body)

	if err != nil {
		sendResponse(w, err)
		return
	}

	h.Ray.Debug("Postback request :val", args.String{N: "val", V: string(bts)})

	sendResponse(w, map[string]interface{}{})
}
