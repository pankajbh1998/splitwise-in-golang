package models

import (
	"math/rand"
	"time"
)

const (
	Equal   = "EQUAL"
	Exact   = "EXACT"
	Percent = "PERCENT"
)

const(
	Show    = "SHOW"
	Expense = "EXPENSE"
)

func GetUniqueID() int64 {
	return rand.Int63()
	return time.Now().Unix()
}


