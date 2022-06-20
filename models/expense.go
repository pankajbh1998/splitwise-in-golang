package models

type Expense struct {
	ID int64
	Payee string
	ExpenseType string
	DistributionAmount []float64
	UserDistribution []string
}
