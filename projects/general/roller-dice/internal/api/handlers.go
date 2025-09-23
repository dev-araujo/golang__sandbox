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

	if request.DiceType != nil {
		diceType := d.DiceType(*request.DiceType)
		result := rand.Intn(int(diceType)) + 1

		response := RollResponse{
			Result:   &result,
			DiceType: request.DiceType,
		}

		outcome := handleCriticalsResults(request, result)
		if outcome != "" {
			response.Outcome = outcome
		}

		respondWithJSON(w, http.StatusOK, response)
		return
	}

	if len(request.DiceTypes) > 0 {
		var rolls []Roll
		var sum uint
		for _, dt := range request.DiceTypes {
			diceType := d.DiceType(dt)
			result := rand.Intn(int(diceType)) + 1
			sum += uint(result)
			rolls = append(rolls, Roll{
				Result:   result,
				DiceType: int(diceType),
			})
		}
		response := RollResponse{
			Rolls:    rolls,
			TotalSum: &sum,
		}
		respondWithJSON(w, http.StatusOK, response)
		return
	}
}
