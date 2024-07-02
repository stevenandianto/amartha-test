package model

import (
	"time"
)

type Installment struct {
	LoanID        int64     `json:"loan_id"`
	InstallmentID int64     `json:"installment_id"`
	Amount        float64   `json:"amount"`
	Status        int       `json:"status"`
	DueDate       time.Time `json:"due_date"`
}
