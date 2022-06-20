package main

import (
	"fmt"
	"splitwise-in-golang/models"
	"splitwise-in-golang/repo"
	"splitwise-in-golang/service"
)

func main(){
	numberOfUsers, commands := getInput()
	userMap := repo.NewUserMap()
	ledgerMap := repo.NewLedgerMap()
	ledgerManager := service.NewLedgerManager(ledgerMap,userMap)
	addUsers(numberOfUsers, userMap)
	for _, val := range commands {
		output,err := service.RunCommand(val, ledgerManager)
		if err != nil {
			fmt.Println(err)
		} else if len(output) != 0 {
			fmt.Println(output)
		}
	}

}

func addUsers(numberOfUsers int, mp *repo.UserMap){
	for i:=1;i<=numberOfUsers;i++ {
		user := &models.User{
			Name: fmt.Sprintf("u%v",i),
		}
		mp.AddUser(user)
	}
}

func getInput()(int,[]string){
	numberOfUsers := 5
	commands := []string {
		"SHOW",
		"SHOW u1",
		"EXPENSE u1 1000 4 u1 u2 u3 u4 EQUAL",
		"SHOW u4",
		"SHOW u1",
		"EXPENSE u1 1250 2 u2 u3 EXACT 370 880",
		"SHOW",
		"EXPENSE u4 1200 4 u1 u2 u3 u4 PERCENT 40 20 20 20",
		"SHOW u1",
		"SHOW",
	}
	return numberOfUsers, commands
}

// Output
//[No balances]
//[No balances]
//[u4 owes u1 : 250
//]
//[u2 owes u1 : 250
// u3 owes u1 : 250
// u4 owes u1 : 250
//]
//[u2 owes u1 : 620
// u3 owes u1 : 1130
// u4 owes u1 : 250
//]
//[u2 owes u1 : 620
// u3 owes u1 : 1130
// u1 owes u4 : 230
//]
//[u3 owes u4 : 240
// u2 owes u1 : 620
// u3 owes u1 : 1130
// u1 owes u4 : 230
// u2 owes u4 : 240
//]