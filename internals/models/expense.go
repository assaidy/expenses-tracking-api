package models

import (
	"time"
)

type Expense struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	AddedAt     time.Time `json:"addedAt"`
}

type ExpenseCreateOrUpdateRequest struct {
	Amount      float64 `json:"amount" validate:"required,numeric"`
	Category    string  `json:"category" validate:"required"`
	Description string  `json:"description" validate:"required"`
}
