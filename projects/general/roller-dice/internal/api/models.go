package api

import (
	"fmt"

	d "github.com/dev-araujo/golang__sandbox/roller-dice/internal/domain"
)

type RollRequest struct {
	DiceType        *int  `json:"dice_type,omitempty"`
	DiceTypes       []int `json:"dice_types,omitempty"`
	CriticalSuccess *int  `json:"critical_success,omitempty"`
	CriticalFailure *int  `json:"critical_failure,omitempty"`
}

func (r *RollRequest) Validate() error {
	if r.DiceType != nil && len(r.DiceTypes) > 0 {
		return fmt.Errorf("não é possível rolar um único dado e vários dados ao mesmo tempo")
	}

	if r.DiceType != nil {
		dice := d.DiceType(*r.DiceType)
		if !dice.Validate() {
			return fmt.Errorf("o tipo de dado é inválido. Os dados possíveis são os de 4, 6, 8, 10, 20 lados")
		}

		if r.CriticalSuccess != nil {
			if *r.CriticalSuccess <= 0 {
				return fmt.Errorf("o valor de sucesso crítico não pode ser menor ou igual a zero")
			}
			if *r.CriticalSuccess > *r.DiceType {
				return fmt.Errorf("o valor de sucesso crítico não pode ser maior que o tipo de dado")
			}
		}

		if r.CriticalFailure != nil {
			if *r.CriticalFailure <= 0 {
				return fmt.Errorf("o valor de falha crítica não pode ser menor ou igual a zero")
			}
			if *r.CriticalFailure > *r.DiceType {
				return fmt.Errorf("o valor de falha crítica não pode ser maior que o tipo de dado")
			}
		}

		if r.CriticalSuccess != nil && r.CriticalFailure != nil && *r.CriticalSuccess == *r.CriticalFailure {
			return fmt.Errorf("os valores de sucesso e falha crítica não podem ser iguais")
		}
	}

	if len(r.DiceTypes) > 0 {
		if r.CriticalSuccess != nil || r.CriticalFailure != nil {
			return fmt.Errorf("não é possível definir sucesso ou falha crítica para rolagens de múltiplos dados")
		}
		for _, dt := range r.DiceTypes {
			dice := d.DiceType(dt)
			if !dice.Validate() {
				return fmt.Errorf("um dos tipos de dado é inválido. Os dados possíveis são os de 4, 6, 8, 10, 20 lados")
			}
		}
	}

	if r.DiceType == nil && len(r.DiceTypes) == 0 {
		return fmt.Errorf("nenhum tipo de dado foi especificado")
	}

	return nil
}

type Roll struct {
	Result   int `json:"result"`
	DiceType int `json:"dice_type"`
}

type RollResponse struct {
	Rolls    []Roll `json:"rolls,omitempty"`
	DiceType *int   `json:"dice_type"`
	Result   *int   `json:"result,omitempty"`
	TotalSum *uint  `json:"total_sum,omitempty"`
	Outcome  string `json:"outcome,omitempty"`
}
