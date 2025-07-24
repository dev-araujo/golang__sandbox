package api

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/dev-araujo/golang__sandbox/roller-dice/pkg"
)

func HandleRollDice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var request pkg.RollRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Corpo do pedido inválido", http.StatusBadRequest)
		return
	}

	if *request.CriticalSuccess > request.DiceType || *request.CriticalFailure > request.DiceType {
		request.CriticalSuccess = nil
		request.CriticalFailure = nil
		http.Error(w, "Corpo do pedido inválido, os valores críticos não podem ser maiores que o tipo de dado", http.StatusBadRequest)
		return
	}

	diceChoiced := pkg.DiceType(request.DiceType)

	if !diceChoiced.IsValid() {
		http.Error(w, "O tipo de dado é inválido. Os dados possíveis são os de 4,6,8,10,20 lados", http.StatusBadRequest)
		return
	}

	result := rand.Intn(int(diceChoiced)) + 1

	response := pkg.RollResponse{
		Result:   result,
		DiceType: int(diceChoiced),
	}

	if handleCriticalsResults(request, result) != "" {
		response.Outcome = handleCriticalsResults(request, result)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Erro ao codificar a resposta JSON: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
	}
}

func handleCriticalsResults(r pkg.RollRequest, result int) string {
	if r.CriticalSuccess != nil && result == *r.CriticalSuccess {
		return "Sucesso crítico"
	}
	if r.CriticalFailure != nil && result == *r.CriticalFailure {
		return "Falha crítica"
	}

	return ""

}
