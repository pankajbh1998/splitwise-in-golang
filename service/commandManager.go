package service

import (
	"fmt"
	"splitwise-in-golang/models"
	"strconv"
	"strings"
)

func RunCommand(input string, l *LedgerManager)(output []string,err error){
	str := strings.Split(input," ")
	switch str[0] {
	case models.Show:
		output,err = l.showExpense(str)
	case models.Expense:
		output,err = l.addExpense(str)
	}
	return
}

func (l *LedgerManager) formatOutput(ledger []*models.Ledger)[]string{
	str := make([]string,0)
	if len(ledger) == 0 {
		str = append(str, "No balances")
	} else {
		for _, val := range ledger {
			str = append(str, fmt.Sprintf("%v owes %v : %v\n",l.user.GetUserByID(val.Creditor).Name, l.user.GetUserByID(val.Debter).Name, val.Amount))
		}
	}
	return str
}

func (l *LedgerManager) showExpense(str[]string)([]string,error) {
	if len(str) == 1 {
		return l.formatOutput(l.showAllExpense()), nil
	}
	userLedger, err := l.showExpenseUser(str[1])
	if err != nil {
		return nil, err
	}
	return l.formatOutput(userLedger),nil
}

func (l *LedgerManager) addExpense(str[]string)([]string,error) {
	payee, err := l.user.GetUserByName(str[1])
	if err != nil {
		return nil, err
	}

	totalExpense, err := strconv.Atoi(str[2])
	if err != nil {
		return nil,err
	}

	numberOfUsers, err := strconv.Atoi(str[3])
	if err != nil {
		return nil,err
	}

	index := 4
	userIDs := make([]int64,0)
	for j:=0;j<numberOfUsers;j++ {
		user, err := l.user.GetUserByName(str[index])
		if err != nil {
			return nil,err
		}
		userIDs = append(userIDs, user.ID)
		index++
	}

	typeOfPartition := str[index]
	index++
	var splittedExpense []float64
	switch typeOfPartition {
	case models.Equal:
		splittedExpense,err = splitEqual(numberOfUsers,float64(totalExpense))
	case models.Exact:
		splittedExpense,err = splitExact(numberOfUsers,float64(totalExpense), str[index:])
	case models.Percent:
		splittedExpense,err = splitPercent(numberOfUsers,float64(totalExpense), str[index:])
	}
	if err != nil {
		return nil, err
	}
	for i:=0;i<numberOfUsers;i++ {
		if payee.ID == userIDs[i] {
			continue
		}
		owe := models.Owe{
			Debter: payee.ID,
			Creditor: userIDs[i],
		}
		ledger := &models.Ledger{
			Owe: owe,
			Amount: splittedExpense[i],
		}
		l.addAnExpense(ledger)
	}

	return nil,nil
}

