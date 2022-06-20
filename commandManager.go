package main

import (
	"splitwise-in-golang/models"
	"splitwise-in-golang/service"
	"strings"
)

func RunCommand(input string, l *service.LedgerManager)(output []string,err error){
	str := strings.Split(input," ")
	switch str[0] {
	case models.ShowTag:
		output,err = l.ShowExpense(str)
	case models.ExpenseTag:
		output,err = l.AddExpense(str)
	}
	return
}

