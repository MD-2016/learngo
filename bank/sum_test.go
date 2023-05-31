package bank

import (
	"reflect"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("any size collections", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := Sum(nums)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSum := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make sums of tails", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSum(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiply all elements", func(t *testing.T) {
		mult := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, mult, 1), 6)
	})

	t.Run("concat strings", func(t *testing.T) {
		concat := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concat, ""), "abc")
	})
}

func TestFind(t *testing.T) {
	t.Run("find even num", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstNum, found := Find(nums, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstNum, 2)
	})

	type Linux struct {
		Name string
	}

	t.Run("find the best distro", func(t *testing.T) {
		distros := []Linux{
			{Name: "Debian"},
			{Name: "Arch"},
			{Name: "Gentoo"},
			{Name: "Slackware"},
		}

		king, found := Find(distros, func(l Linux) bool {
			return strings.Contains(l.Name, "Arch")
		})
		AssertTrue(t, found)
		AssertEqual(t, king, Linux{Name: "Arch"})
	})
}
