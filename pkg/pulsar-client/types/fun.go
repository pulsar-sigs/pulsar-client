package types

import (
	"log"
	"net/http"
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func RunReadnessAPI() {
	http.Handle("/readness", &handler{})
	err := http.ListenAndServe(":9494", nil)
	if err != nil {
		log.Fatal("start http server failed!", err)
	}
}
