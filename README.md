# Amartha Billing Engine

This project is a simple billing engine system written in Go. It includes functionality for creating loans, making payments, checking delinquencies, and displaying billing schedules.

## Project Structure
amartha-test/

├── constant/

│ └── constants.go

├── model/

│ └── installment.go

│ └── loan.go

│ └── user.go

├── usecase/

│ ├── billingengine.go

│ └── billingengine_test.go

├── go.mod

└── README.md


### Directory Descriptions

- `constant/`: Contains constants used throughout the application.
- `model/`: Contains model definitions for loans and installments.
- `usecase/`: Contains the main business logic and test cases for the application.


## Unit Test
To run the unit testing use the command
```go
go test ./... </code>
```

## Installation

### Step 1
```go
git clone https://github.com/stevenandianto/amartha-test.git
```

### Step 2
```go
cd amartha-test
```

### Step 3
```go
go run main.go
```

or

```go 
./amartha-test
```



