package selects

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(urls ...string) (string, error) {
	return ConfigurableRacer(tenSecondTimeout, urls...)
}

func ConfigurableRacer(timeout time.Duration, urls ...string) (string, error) {
	// Kanał buforowany o rozmiarze len(urls) zapobiega wyciekowi pamięci (goroutine leak)
	// - goroutines, które przegrają wyścig, nie zablokują się na zapisie.
	ch := make(chan string, len(urls))

	for _, url := range urls {
		// Pamiętaj o przekazaniu zmiennej jako argument do domknięcia (closure)
		go func(u string) {
			resp, err := http.Get(u)
			if err == nil {
				resp.Body.Close()
			}
			ch <- u
		}(url)
	}

	// Czekamy na pierwszy wynik lub timeout
	select {
	case winner := <-ch:
		return winner, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout - żaden serwer nie odpowiedział na czas")
	}
}
