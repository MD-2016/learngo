package bank

import "testing"

func TestBadBank(t *testing.T) {
	var (
		bob  = Account{Name: "Bob", Balance: 100}
		jim  = Account{Name: "Jim", Balance: 75}
		john = Account{Name: "John", Balance: 200}

		trans = []Transaction{
			NewTransaction(jim, bob, 100),
			NewTransaction(john, jim, 25),
		}
	)

	NewBalanceFor := func(acc Account) float64 {
		return NewBalanceFor(acc, trans).Balance
	}

	AssertEqual(t, NewBalanceFor(bob), 200)
	AssertEqual(t, NewBalanceFor(jim), 0)
	AssertEqual(t, NewBalanceFor(john), 175)
}
