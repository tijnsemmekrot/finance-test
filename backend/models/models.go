package models

type ExpenseRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Bank        string  `json:"bank"`
}
