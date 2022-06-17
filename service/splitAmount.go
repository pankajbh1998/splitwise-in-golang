package service

import (
	"fmt"
	"strconv"
)

func splitEqual(numberOfUsers int , expense float64)([]float64, error){
	ret := make([]float64,0)
	for i:=0;i<numberOfUsers;i++ {
		ret = append(ret, expense/float64(numberOfUsers))
	}
	return ret,nil
}

func splitExact(numberOfUsers int , totalExpense float64, command []string)([]float64,error){
	ret := make([]float64,0)
	sumExpense := 0
	for _, val := range command {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return nil,err
		}
		sumExpense += intVal
		ret = append(ret, float64(intVal))
	}

	if totalExpense != float64(sumExpense) {
		return nil, fmt.Errorf("the expense must be equal to split amount")
	}

	return ret,nil
}

func splitPercent(numberOfUsers int , expense float64, command []string)([]float64,error){
	ret := make([]float64,0)
	sumPercent := 0
	for _, val := range command {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return nil,err
		}
		sumPercent += intVal
		ret = append(ret, expense*float64(intVal)/100)
	}

	if sumPercent != 100 {
		return nil, fmt.Errorf("the expense percent must be equal 100")
	}

	return ret,nil
}
