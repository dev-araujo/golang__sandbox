package task

type Task struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
