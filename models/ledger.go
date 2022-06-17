package models

type Owe struct {
	Debter   int64
	Creditor int64
}

type Ledger struct{
	Owe
	Amount float64
}
