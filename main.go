package main

import (
	"fmt"
	"learngo/accounts"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())

	err := account.Withdraw(10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
