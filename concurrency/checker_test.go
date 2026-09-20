package concurrency

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "test2"
}

func TestCheckWebstite(t *testing.T) {
	t.Run("simple check", func(t *testing.T) {
		urls := []string{
			"test1",
			"test2",
		}

		got := CheckWebsite(mockWebsiteChecker, urls)
		expected := map[string]bool {
			"test1": true,
			"test2": false,
		}

		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("Result is not as expetced. Got %v wanted %v", got, expected)
		}
	})
}

func mocSlowkWebsiteChecker(url string) bool {
	time.Sleep(25 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := range urls {
		urls[i] = fmt.Sprintf("next url is %d", i)
	}
	
	for b.Loop() {
		CheckWebsite(mocSlowkWebsiteChecker, urls)
	}
} 
