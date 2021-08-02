package reflection_test

import (
	"reflect"
	"testing"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/reflection"
)

func TestWalk(t *testing.T) {
	type Profile struct {
		Age  int
		City string
	}
	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Josh"},
			[]string{"Josh"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Josh", "Fort Collins"},
			[]string{"Josh", "Fort Collins"},
		},
		{
			"struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Josh", 30},
			[]string{"Josh"},
		},
		{
			"nested fields",
			Person{
				"Josh",
				Profile{30, "Fort Collins"},
			},
			[]string{"Josh", "Fort Collins"},
		},
		{
			"x is a pointer",
			&Person{
				"Josh",
				Profile{30, "Fort Collins"},
			},
			[]string{"Josh", "Fort Collins"},
		},
		{
			"slice of structs",
			[]Profile{
				{30, "Fort Collins"},
				{7, "Fort Collins"},
			},
			[]string{"Fort Collins", "Fort Collins"},
		},
		{
			"arrays",
			[2]Profile{
				{30, "Fort Collins"},
				{7, "Fort Collins"},
			},
			[]string{"Fort Collins", "Fort Collins"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			reflection.Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		reflection.Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{30, "Fort Collins"}
			aChannel <- Profile{7, "Fort Collins"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Fort Collins", "Fort Collins"}

		reflection.Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{30, "Fort Collins"}, Profile{7, "Fort Collins"}
		}

		var got []string
		want := []string{"Fort Collins", "Fort Collins"}

		reflection.Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
