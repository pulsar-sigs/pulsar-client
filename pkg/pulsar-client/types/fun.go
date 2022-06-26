package types

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// go tool pprof -http=:1234 http://localhost:9494/debug/pprof/goroutine
}

func RunReadnessAPI() {
	http.Handle("/readness", &handler{})
	err := http.ListenAndServe(":9494", nil)
	if err != nil {
		log.Fatal("start http server failed!", err)
	}
}
