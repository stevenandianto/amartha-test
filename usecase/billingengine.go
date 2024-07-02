package usecase

import (
	"amartha-test/constant"
	"amartha-test/model"
	"fmt"
	"time"
)

// global variable to store the created loan
var loan model.Loan
var installment model.Installment
var billings []model.Installment

// getOutstanding returns the current outstanding on a loan, 0 if no outstanding(or closed)
func GetOutstanding() {
	fmt.Printf("Outstanding amount : %2.f \n", loan.OutstandingAmount)
}

// isDelinquent returns true if there are more than 2 weeks of Non payment of the loan amount
func IsDelinquent(now time.Time) bool {
	var unpaidCount = 0
	for _, installment := range billings {
		if installment.DueDate.Before(now) && installment.Status == constant.StatusUnpaid {
			unpaidCount += 1
		}
	}

	return unpaidCount > 2
}

// makePayment make a payment of certain amount on the loan
func MakePayment(loanID int64, amount float64) {
	// assume that borrower can only pay the exact amount of payable that week or not pay at all
	for index, installment := range billings {
		if installment.Status == constant.StatusUnpaid && installment.Amount == amount {
			billings[index].Status = constant.StatusPaid
			loan.OutstandingAmount = loan.OutstandingAmount - amount
			break
		}
	}
}

func ShowBillingSchedule() {
	fmt.Println("Here is your billing schedule")
	for _, installment := range billings {
		fmt.Printf("%s : %2.f , Status: %d \n", installment.DueDate.Format(time.RFC822), installment.Amount, installment.Status)
	}
}

func CreateLoan(loanID int64, userID int64, amount float64) {
	// set loan
	loan.LoanID = loanID
	loan.UserID = userID
	loan.Amount = amount + constant.Interest*amount
	loan.Status = constant.StatusUnpaid
	loan.OutstandingAmount = loan.Amount

	var principalAmount = amount + constant.Interest*amount
	// set installment
	for i := 0; i < constant.Tenure; i++ {
		installment.InstallmentID = int64(i + 1)
		installment.Amount = principalAmount / constant.Tenure
		installment.Status = constant.StatusUnpaid
		installment.DueDate = time.Now().AddDate(0, 0, 7*(i+1))
		billings = append(billings, installment)
	}
}
