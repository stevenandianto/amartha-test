package usecase

import (
	"amartha-test/constant"
	"amartha-test/model"
	"bytes"
	"io"
	"os"
	"testing"
	"time"
)

func TestCreateLoan(t *testing.T) {
	CreateLoan(1, 1, 1000)
	if loan.LoanID != 1 || loan.UserID != 1 || loan.Amount != 1100 || loan.OutstandingAmount != 1100 {
		t.Error("Loan creation failed")
	}
	if len(billings) != 50 {
		t.Error("Expected 50 installments")
	}
}

func TestGetOutstanding(t *testing.T) {
	loan = model.Loan{
		OutstandingAmount: 1000,
	}
	expected := "Outstanding amount : 1000 \n"
	if actual := captureOutput(GetOutstanding); actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestIsDelinquent(t *testing.T) {
	now := time.Now()
	billings = []model.Installment{
		{DueDate: now.AddDate(0, 0, -21), Status: constant.StatusUnpaid},
		{DueDate: now.AddDate(0, 0, -14), Status: constant.StatusUnpaid},
		{DueDate: now.AddDate(0, 0, -7), Status: constant.StatusUnpaid},
	}
	if !IsDelinquent(now) {
		t.Error("Expected loan to be delinquent")
	}

	billings = []model.Installment{
		{DueDate: now.AddDate(0, 0, -14), Status: constant.StatusUnpaid},
		{DueDate: now.AddDate(0, 0, -7), Status: constant.StatusUnpaid},
	}
	if IsDelinquent(now) {
		t.Error("Expected loan to not be delinquent")
	}
}

func TestMakePayment(t *testing.T) {
	loan = model.Loan{
		OutstandingAmount: 1000.0,
	}
	billings = []model.Installment{
		{InstallmentID: 1, Amount: 250.0, Status: constant.StatusUnpaid},
		{InstallmentID: 2, Amount: 250.0, Status: constant.StatusUnpaid},
	}
	MakePayment(1, 250.0)
	if loan.OutstandingAmount != 750.0 {
		t.Errorf("Expected OutstandingAmount to be 750.0 but got %f", loan.OutstandingAmount)
	}
	if billings[0].Status != constant.StatusPaid {
		t.Error("Expected first installment to be paid")
	}
}

func TestShowBillingSchedule(t *testing.T) {
	now := time.Now()
	billings = []model.Installment{
		{DueDate: now.AddDate(0, 0, 7), Amount: 250, Status: constant.StatusUnpaid},
	}
	expected := "Here is your billing schedule\n" +
		now.AddDate(0, 0, 7).Format(time.RFC822) + " : 250 , Status: 1 \n"
	if actual := captureOutput(ShowBillingSchedule); actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = writer

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		outC <- buf.String()
	}()

	f()
	writer.Close()
	os.Stdout = stdout
	return <-outC
}
