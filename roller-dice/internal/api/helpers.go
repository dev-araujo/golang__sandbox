package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Erro ao codificar a resposta JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Erro interno do servidor"}`))
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func handleCriticalsResults(r RollRequest, result int) string {
	if r.CriticalSuccess != nil && result == *r.CriticalSuccess {
		return "Sucesso crítico"
	}
	if r.CriticalFailure != nil && result == *r.CriticalFailure {
		return "Falha crítica"
	}

	return ""

}
