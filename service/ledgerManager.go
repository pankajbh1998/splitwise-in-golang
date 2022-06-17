package service

import (
	"splitwise-in-golang/models"
	"splitwise-in-golang/repo"
)

type LedgerManager struct {
	ledger *repo.LedgerMap
	user *repo.UserMap
}

func NewLedgerManager(l *repo.LedgerMap, u *repo.UserMap)*LedgerManager{
	return &LedgerManager{
		ledger: l,
		user: u,
	}
}

func (r *LedgerManager) addAnExpense(l *models.Ledger){
	// if Creditor  already owes money from Debitor
	led, ok := r.ledger.GetLedger(&l.Owe)
	if ok {
		l.Amount += led.Amount
		r.ledger.AddExpense(l)
		return
	}

	// if Debitor owes money from Creditor
	o := models.Owe{
		Debter: l.Creditor,
		Creditor: l.Debter,
	}
	led, ok = r.ledger.GetLedger(&o)
	if ok {
		if led.Amount > l.Amount {
			led.Amount -= l.Amount
			r.ledger.AddExpense(led)
			return
		}
		r.ledger.RemoveLedger(&o)
		if led.Amount < l.Amount {
			l.Amount -= led.Amount
			r.ledger.AddExpense(l)
		}
		return
	}

	r.ledger.AddExpense(l)
	return
}

func (r *LedgerManager) showAllExpense()[]*models.Ledger {
	return r.ledger.GetAllLedger()
}

func (r *LedgerManager) showExpenseUser(name string)([]*models.Ledger, error) {
	user,err := r.user.GetUserByName(name)
	if err != nil {
		return nil,err
	}
	ledger := r.ledger.GetAllLedger()
	ret := make([]*models.Ledger,0)
	for _, val := range ledger {
		if val.Creditor ==  user.ID || val.Debter == user.ID {
			ret = append(ret, val)
		}
	}
	return ret,nil
}
