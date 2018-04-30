package main

import (
	"github.com/mono83/xray"
	xrayOs "github.com/mono83/xray/out/os"
	"flag"
	"net/http"
	"github.com/2at2/covery-mock/handler"
	"strconv"
	"github.com/mono83/xray/args"
)

func main() {
	xray.ROOT.On(xrayOs.StdOutLogger(xray.INFO))
	ray := xray.ROOT.Fork().WithLogger("covery")

	var port int
	flag.IntVar(&port, "port", 80, "")
	flag.Parse()

	http.Handle("/api/ping", handler.Ping{Ray: ray})
	http.Handle("/api/makeDecision", handler.MakeDecision{Ray: ray})
	http.Handle("/api/postback", handler.Postback{Ray: ray})
	http.Handle("/manage/nodeNames", handler.NodeNames{Ray: ray})
	http.Handle("/manage/trustListCheck", handler.TrustListCheck{Ray: ray})

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)

	if err != nil {
		ray.Error("Unable to start http server - :err", args.Error{Err: err})
	}
}
