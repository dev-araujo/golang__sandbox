package pkg

type RollRequest struct {
	DiceType int `json:"dice_type"`
}

type RollResponse struct {
	Result   int `json:"result"`
	DiceType int `json:"dice_type"`
}
