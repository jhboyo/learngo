package main

import (
	"fmt"
	//"learngo/accounts"
	"learngo/mydict"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

		dictionary := mydict.Dictionary{}
		baseWord := "hello"

		dictionary.Add(baseWord, "First")
		dictionary.Search(baseWord)
		dictionary.Delete(baseWord)

		word, err := dictionary.Search(baseWord)
		if err != nil {
			fmt.Println(err)
		} else {	
			fmt.Println(word)
		}
 

	// account := accounts.NewAccount("nico")
	// account.Deposit(10)
	// fmt.Println(account)

	// err := account.Withdraw(10)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance())
}
