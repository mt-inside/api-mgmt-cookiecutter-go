package health

import (
	"encoding/json"
	"net/http"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.project_slug}}/version"
)

func ListenHealthChecks(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleLive)
	mux.HandleFunc("/ready", handleReady)
	http.ListenAndServe(addr, mux)
}

type liveResponse struct {
	Status  string
	App     string
	Version string
}

func handleLive(w http.ResponseWriter, r *http.Request) {
	live := &liveResponse{
		Status:  "ok",
		App:     version.Name,
		Version: version.Version,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(live)
}

type readyResponse struct {
	Ready   string
	App     string
	Version string
}

func handleReady(w http.ResponseWriter, r *http.Request) {
	ready := &readyResponse{
		Ready:   "ok",
		App:     version.Name,
		Version: version.Version,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ready)
}
