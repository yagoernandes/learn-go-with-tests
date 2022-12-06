package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestWebsiteChecker(t *testing.T) {
	websites := []string{
		"ok.com",
		"ok2.com",
		"wrong_site_url",
	}

	want := map[string]bool{
		"ok.com":         true,
		"ok2.com":        true,
		"wrong_site_url": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func BenchmarkChekWebsites(b *testing.B) {
	websites := make([]string, 100)

	for i := 0; i < len(websites); i++ {
		websites[i] = "url_test"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, websites)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func mockWebsiteChecker(url string) bool {
	return url != "wrong_site_url"
}
