package api

import (
	"fmt"

	d "github.com/dev-araujo/golang__sandbox/roller-dice/internal/domain"
)

type RollRequest struct {
	DiceType        int  `json:"dice_type"`
	CriticalSuccess *int `json:"critical_success,omitempty"`
	CriticalFailure *int `json:"critical_failure,omitempty"`
}

func (r *RollRequest) Validate() error {
	dice := d.DiceType(r.DiceType)
	if !dice.IsValid() {
		return fmt.Errorf("o tipo de dado é inválido. Os dados possíveis são os de 4, 6, 8, 10, 20 lados")
	}

	if r.CriticalSuccess != nil {
		if *r.CriticalSuccess <= 0 {
			return fmt.Errorf("o valor de sucesso crítico não pode ser menor ou igual a zero")
		}
		if *r.CriticalSuccess > r.DiceType {
			return fmt.Errorf("o valor de sucesso crítico não pode ser maior que o tipo de dado")
		}
	}

	if r.CriticalFailure != nil {
		if *r.CriticalFailure <= 0 {
			return fmt.Errorf("o valor de falha crítica não pode ser menor ou igual a zero")
		}
		if *r.CriticalFailure > r.DiceType {
			return fmt.Errorf("o valor de falha crítica não pode ser maior que o tipo de dado")
		}
	}

	if r.CriticalSuccess != nil && r.CriticalFailure != nil && *r.CriticalSuccess == *r.CriticalFailure {
		return fmt.Errorf("os valores de sucesso e falha crítica não podem ser iguais")
	}

	return nil
}

type RollResponse struct {
	Result   int    `json:"result"`
	DiceType int    `json:"dice_type"`
	Outcome  string `json:"outcome,omitempty"`
}
