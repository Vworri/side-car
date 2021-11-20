package services

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func StartServer() {
	log.Info("Starting Sidecar API server")
	router := mux.NewRouter()
	router.HandleFunc("/health", Healthz).Methods("GET")
	http.ListenAndServe(":8000", router)
}
