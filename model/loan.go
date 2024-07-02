package model

type Loan struct {
	LoanID            int64   `json:"loan_id"`
	UserID            int64   `json:"user_id"`
	Amount            float64 `json:"amount"`
	OutstandingAmount float64 `json:"outstanding_amount"`
	Status            int     `json:"status"`
}
