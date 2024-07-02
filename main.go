package main

import (
	"amartha-test/model"
	"amartha-test/usecase"
	"fmt"
	"time"
)

// global variable to store the created loan
var loan model.Loan
var installment model.Installment
var billings []model.Installment

func main() {
	var loanID = int64(100)
	var userID = int64(1)
	var amount = float64(5000000)

	// create a loan
	usecase.CreateLoan(loanID, userID, amount)

	// check billing schedule
	fmt.Println("======================================================================")
	usecase.ShowBillingSchedule()

	// check outstanding, user havent paid so outstanding = principal
	usecase.GetOutstanding()

	/*
	* first payment case
	* first installment become paid
	* outstanding amount decreased
	 */
	fmt.Println("======================================================================")
	usecase.MakePayment(loanID, 110000)
	usecase.ShowBillingSchedule()
	usecase.GetOutstanding()

	/*
	* second payment case
	* first and second installment become paid
	* outstanding amount decreased
	 */
	fmt.Println("======================================================================")
	usecase.MakePayment(loanID, 110000)
	usecase.ShowBillingSchedule()
	usecase.GetOutstanding()

	/*
	* skipped payment case
	* already paid 2 weeks
	* assume no payment in next 3 weeks (W3, W4, W5)
	 */
	fmt.Println("======================================================================")
	//check isDelinquent on W4 or 28 weeks it returns false
	fmt.Println("Is the loan Delinquent?", usecase.IsDelinquent(time.Now().AddDate(0, 0, 28)))

	//check isDelinquent on W5 or 35 weeks it returns true
	fmt.Println("Is the loan Delinquent?", usecase.IsDelinquent(time.Now().AddDate(0, 0, 35)))

}
