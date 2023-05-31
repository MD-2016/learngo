package bank

func Find[T any](items []T, predicate func(T) bool) (val T, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}

func Reduce[T, U any](collection []T, accumulator func(U, T) U, initialValue U) U {
	var res = initialValue
	for _, x := range collection {
		res = accumulator(res, x)
	}
	return res
}
