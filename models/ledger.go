package models

type Owe struct {
	Debter   string
	Creditor string
}

type Ledger struct{
	Owe
	Amount float64
}
