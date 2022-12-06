package website_racer

import "testing"

func TestRacer(t *testing.T) {
	slowUrl := "http://globo.com"
	fastUrl := "http://quii.dev"

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
