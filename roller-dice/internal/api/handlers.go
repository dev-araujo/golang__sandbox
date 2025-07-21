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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Erro ao codificar a resposta JSON: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
	}
}
