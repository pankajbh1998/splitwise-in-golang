package service

import (
	"fmt"
	"splitwise-in-golang/models"
	"strconv"
)

type LedgerManager struct {
	expense []models.Expense
	users 	[]models.User
	ledger  map[models.Owe]float64
}

func NewLedgerManager(users []models.User)*LedgerManager{
	return &LedgerManager{
		expense: make([]models.Expense,0),
		users : users,
		ledger: make(map[models.Owe]float64,0),
	}
}

func( l *LedgerManager)AddUser(u *models.User){
	l.users = append(l.users, *u)
}

func (l *LedgerManager)addToExpense(expenseType string, payee string,  amount []float64, users []string){
	exp := models.Expense{
		ID:                 models.GetUniqueID(),
		Payee: payee,
		ExpenseType:        expenseType,
		DistributionAmount: amount,
		UserDistribution:   users,
	}
	l.expense = append(l.expense, exp)
}

func (r *LedgerManager) addToLedger(l *models.Ledger){
	// if Creditor  already owes money from Debitor
	amt, ok := r.ledger[l.Owe]
	if ok {
		r.ledger[l.Owe] += l.Amount
		return
	}

	// if Debitor owes money from Creditor
	o := models.Owe{
		Debter: l.Creditor,
		Creditor: l.Debter,
	}
	amt, ok = r.ledger[o]
	if ok {
		if amt > l.Amount {
			r.ledger[o] -= l.Amount
			return
		}
		delete(r.ledger, o)
		if amt < l.Amount {
			r.ledger[l.Owe] = l.Amount - amt
		}
		return
	}
	// if no expense for that user
	r.ledger[l.Owe] = l.Amount
	return
}

func (r *LedgerManager) showAllExpense()[]*models.Ledger {
	ret := make([]*models.Ledger,0)
	for key, val := range r.ledger {
			ret = append(ret, &models.Ledger{
				Owe:    key,
				Amount: val,
			})
	}
	return ret
}

func (r *LedgerManager) showExpenseUser(user string)([]*models.Ledger, error) {
	ret := make([]*models.Ledger,0)
	for key, val := range r.ledger {
		if key.Creditor ==  user || key.Debter == user {
			ret = append(ret, &models.Ledger{
				Owe:    key,
				Amount: val,
			})
		}
	}
	return ret, nil
}

func (l *LedgerManager) formatOutput(ledger []*models.Ledger)[]string{
	str := make([]string,0)
	if len(ledger) == 0 {
		str = append(str, "No balances")
	} else {
		for _, val := range ledger {
			str = append(str, fmt.Sprintf("%v owes %v : %v\n",val.Creditor, val.Debter, val.Amount))
		}
	}
	return str
}

func (l *LedgerManager) ShowExpense(str[]string)([]string,error) {
	if len(str) == 1 {
		return l.formatOutput(l.showAllExpense()), nil
	}
	userLedger, err := l.showExpenseUser(str[1])
	if err != nil {
		return nil, err
	}
	return l.formatOutput(userLedger),nil
}

func (l *LedgerManager) AddExpense(str[]string)([]string,error) {
	payee := str[1]

	totalExpense, err := strconv.Atoi(str[2])
	if err != nil {
		return nil,err
	}

	numberOfUsers, err := strconv.Atoi(str[3])
	if err != nil {
		return nil,err
	}

	index := 4
	users := make([]string,0)
	for j:=0;j<numberOfUsers;j++ {
		users = append(users, str[index])
		index++
	}

	expenseType := str[index]
	index++
	var splitedExpense []float64
	switch expenseType {
	case models.Equal:
		splitedExpense,err = splitEqual(numberOfUsers,float64(totalExpense))
	case models.Exact:
		splitedExpense,err = splitExact(numberOfUsers,float64(totalExpense), str[index:])
	case models.Percent:
		splitedExpense,err = splitPercent(numberOfUsers,float64(totalExpense), str[index:])
	}
	if err != nil {
		return nil, err
	}

	l.addToExpense(expenseType, payee, splitedExpense, users)
	for i:=0;i<numberOfUsers;i++ {
		if payee == users[i] {
			continue
		}
		ledger := &models.Ledger{
			Owe: models.Owe{
				Debter: payee,
				Creditor: users[i],
			},
			Amount: splitedExpense[i],
		}
		l.addToLedger(ledger)
	}

	return nil,nil
}