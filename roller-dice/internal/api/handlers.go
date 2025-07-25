package api

import (
	"encoding/json"
	"math/rand"
	"net/http"

	d "github.com/dev-araujo/golang__sandbox/roller-dice/internal/domain"
)

func HandleRollDice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	var request RollRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Corpo do pedido inválido")
		return
	}

	if err := request.Validate(); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	diceType := d.DiceType(request.DiceType)
	result := rand.Intn(int(diceType)) + 1

	response := RollResponse{
		Result:   result,
		DiceType: int(diceType),
	}

	outcome := handleCriticalsResults(request, result)
	if outcome != "" {
		response.Outcome = outcome
	}

	respondWithJSON(w, http.StatusOK, response)
}
