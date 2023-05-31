package bank

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func NewBalanceFor(acc Account, trans []Transaction) Account {
	return Reduce(
		trans,
		applyTransaction,
		acc,
	)
}

func applyTransaction(a Account, trans Transaction) Account {
	if trans.From == a.Name {
		a.Balance -= trans.Sum
	}
	if trans.To == a.Name {
		a.Balance += trans.Sum
	}

	return a
}
