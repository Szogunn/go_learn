package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func getHttpHandler(delay time.Duration) http.HandlerFunc {
	someFunc := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(delay))
		w.WriteHeader(http.StatusOK)
	}

	return someFunc
}

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := httptest.NewServer(getHttpHandler(20 * time.Millisecond))
		fastServer := httptest.NewServer(getHttpHandler(0))
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("want %q and got %q", want, got)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		slowServer := httptest.NewServer(getHttpHandler(101 * time.Millisecond))
		fastServer := httptest.NewServer(getHttpHandler(102 * time.Millisecond))
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		_, err := ConfigurableRacer(100 * time.Millisecond, slowURL, fastURL)
		if err == nil {
			t.Errorf("Expected Error but didn't get one")
		}
	})

}
