package pkg

type RollRequest struct {
	DiceType        int  `json:"dice_type"`
	CriticalSuccess *int `json:"critical_success"`
	CriticalFailure *int `json:"critical_failure"`
}

type RollResponse struct {
	Result   int    `json:"result"`
	DiceType int    `json:"dice_type"`
	Outcome  string `json:"outcome,omitempty"`
}
