package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://foo" {
		return false
	}
	return true
}

func TestHelloWorld(t *testing.T) {
	websites := []string{
		"http://google.com",
		"waat://foo",
	}

	want := map[string]bool{
		"http://google.com": true,
		"waat://foo":        false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
