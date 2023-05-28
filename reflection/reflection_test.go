package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string",
			struct{ Name string }{"MD"},
			[]string{"MD"},
		},
		{
			"struct with two strings",
			struct {
				Name string
				City string
			}{"MD", "Lexington"},
			[]string{"MD", "Lexington"},
		},
		{
			"struct with non string",
			struct {
				Name string
				Age  int
			}{"MD", 32},
			[]string{"MD"},
		},
		{
			"nested",
			Person{
				"MD",
				Profile{32, "Lexington"},
			},
			[]string{"MD", "Lexington"},
		},
		{
			"pointers to things",
			&Person{
				"MD",
				Profile{32, "Lexington"},
			},
			[]string{"MD", "Lexington"},
		},
		{
			"arrays",
			[2]Profile{
				{32, "Lexington"},
				{31, "Huntington"},
			},
			[]string{"Lexington", "Huntington"},
		},
		{
			"slices",
			[]Profile{
				{32, "Lexington"},
				{31, "Huntington"},
			},
			[]string{"Lexington", "Huntington"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		testMap := map[string]string{
			"Foo": "Bar",
			"Faz": "Foz",
		}

		var got []string
		walk(testMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Foz")
	})

	t.Run("with channels", func(t *testing.T) {
		testChan := make(chan Profile)

		go func() {
			testChan <- Profile{32, "Austin"}
			testChan <- Profile{33, "Dallas"}
			close(testChan)
		}()

		var got []string
		want := []string{"Austin", "Dallas"}

		walk(testChan, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFun := func() (Profile, Profile) {
			return Profile{32, "Austin"}, Profile{33, "Dallas"}
		}

		var got []string
		want := []string{"Austin", "Dallas"}

		walk(aFun, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t testing.TB, hay []string, needle string) {
	t.Helper()
	contains := false
	for _, y := range hay {
		if y == needle {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", hay, needle)
	}
}
