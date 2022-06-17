package repo

import "splitwise-in-golang/models"

type LedgerMap struct{
	mp map[models.Owe]float64
}

func NewLedgerMap()*LedgerMap{
	return &LedgerMap{
		mp: make(map[models.Owe]float64),
	}
}

func (l *LedgerMap)AddExpense(ledger *models.Ledger){
	l.mp[ledger.Owe] = ledger.Amount
}

func (l *LedgerMap)GetLedger(o *models.Owe)( *models.Ledger,bool ){
	val, ok := l.mp[*o]
	if !ok {
		return nil, ok
	}
	return &models.Ledger{
		Owe: *o,
		Amount: val,
	}, ok
}

func (l *LedgerMap)GetAllLedger() []*models.Ledger{
	ret := make([]*models.Ledger,0)
	for key,val := range l.mp {
		newLedger := &models.Ledger{
			Owe: key,
			Amount: val,
		}
		ret = append(ret,newLedger)
	}
	return ret
}

func (l *LedgerMap) RemoveLedger(o *models.Owe){
	delete(l.mp, *o)
}