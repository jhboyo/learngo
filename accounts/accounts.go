package accounts

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

// GetOwner returns the account owner
/*func (a *Account) GetOwner() string {
	return a.owner
}*/
