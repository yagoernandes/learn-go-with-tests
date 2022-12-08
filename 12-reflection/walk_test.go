package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Yago"},
			[]string{"Yago"},
		},
		{
			"struct with two strings field",
			struct {
				Name    string
				Surname string
			}{"Yago", "Ernandes"},
			[]string{"Yago", "Ernandes"},
		},
		{
			"struct with non string field",
			Profile{26, "Brasília"},
			[]string{"Brasília"},
		},
		{
			"nested fields",
			Person{"Yago", Profile{26, "Brasília"}},
			[]string{"Yago", "Brasília"},
		},
		{
			"pointers to things",
			&Person{
				"Yago",
				Profile{26, "Brasília"},
			},
			[]string{"Yago", "Brasília"},
		},
		{
			"slices",
			[]Profile{
				{26, "Brasília"},
				{18, "Correntina"},
				{15, "Barreiras"},
			},
			[]string{"Brasília", "Correntina", "Barreiras"},
		},
		{
			"arrays",
			[3]Profile{
				{26, "Brasília"},
				{18, "Correntina"},
				{15, "Barreiras"},
			},
			[]string{"Brasília", "Correntina", "Barreiras"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %q, want %q", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Name": "Yago",
			"City": "Brasília",
		}

		var got []string

		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Yago")
		assertContains(t, got, "Brasília")
	})

	t.Run("with chans", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{26, "Brasília"}
			aChannel <- Profile{18, "Correntina"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Brasília", "Correntina"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{26, "Brasília"}, Profile{18, "Correntina"}
		}

		var got []string
		want := []string{"Brasília", "Correntina"}

		Walk(aFunction, func(s string) {
			got = append(got, s)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, key := range haystack {
		if key == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q, but it didn't", haystack, needle)
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
