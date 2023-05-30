package generics

import "testing"

func TestAssertFuncs(t *testing.T) {
	t.Run("assert on ints", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("assert on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "bob")
	})
}

func TestStack(t *testing.T) {
	t.Run("int stack", func(t *testing.T) {
		intStack := new(Stack[int])

		AssertTrue(t, intStack.IsEmpty())

		intStack.Push(123)
		AssertFalse(t, intStack.IsEmpty())

		intStack.Push(567)
		val, _ := intStack.Pop()
		AssertEqual(t, val, 567)
		val, _ = intStack.Pop()
		AssertEqual(t, val, 123)
		AssertTrue(t, intStack.IsEmpty())

		intStack.Push(1)
		intStack.Push(2)
		frn, _ := intStack.Pop()
		sec, _ := intStack.Pop()
		AssertEqual(t, frn+sec, 3)
	})
}
