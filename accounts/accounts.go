package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//receive pointer of Account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance returns the balance
func (a Account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("Can't withdraw")

//Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney 
	}
	a.balance -= amount
	return nil
}

//ChangeOwner of the account
func (a Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance()) 
}

