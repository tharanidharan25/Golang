/*
Assignment

Update Balance
Textio needs a new way to update user's account balance.

Assignment
Implement the updateBalance function. It should take a customer pointer and a transaction, and return an error. Depending on the transactionType, it should either add or subtract the amount from the customer's balance. If the customer does not have enough money, it should return the error insufficient funds. If the transactionType isn't a withdrawal or deposit, it should return the error unknown transaction type. Otherwise, it should process the transaction and return nil.

alice := customer{id: 1, balance: 100.0}
deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

updateBalance(&alice, deposit)
// id: 1 balance: 150
*/

package main

import (
	"errors"
	"fmt"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(c *customer, t *transaction) error {
	if t.transactionType == transactionWithdrawal && c.balance < t.amount {
		return errors.New("Insufficient Balance")
	} else if t.transactionType != transactionDeposit && t.transactionType != transactionWithdrawal {
		return errors.New("Unknown Transaction Type")
	} else if t.transactionType == transactionDeposit {
		c.balance += t.amount
	} else if t.transactionType == transactionWithdrawal {
		c.balance -= t.amount
	}
	return nil
}

func main() {
	// Deposit operation
	c := customer{id: 1, balance: 100.0}
	t := transaction{customerID: 1, amount: 50.0, transactionType: transactionDeposit}
	ok := updateBalance(&c, &t)
	if ok == nil {
		fmt.Println(c.balance)
	} else {
		fmt.Println(ok)
	}
	
	// Withdrawal operation
	c = customer{id: 2, balance: 200.0}
	t = transaction{customerID: 2, amount: 100.0, transactionType: transactionWithdrawal}
	ok = updateBalance(&c, &t)
	if ok == nil {
		fmt.Println(c.balance)
	} else {
		fmt.Println(ok)
	}
	
	// insufficient funds for withdrawal
	c = customer{id: 3, balance: 50.0}
	t = transaction{customerID: 3, amount: 100.0, transactionType: transactionWithdrawal}
	ok = updateBalance(&c, &t)
	if ok == nil {
		fmt.Println(c.balance)
	} else {
		fmt.Println(ok)
	}
	
	// unknown transaction type
	c = customer{id: 4, balance: 100.0}
	t = transaction{customerID: 4, amount: 50.0, transactionType: "unknown"}
	ok = updateBalance(&c, &t)
	if ok == nil {
		fmt.Println(c.balance)
	} else {
		fmt.Println(ok)
	}
}