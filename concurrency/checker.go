package concurrency

type WebsiteChecker func(url string) bool

type Result struct {
	url string
	ok bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool, len(urls))
	resultChannel := make(chan Result)

	for _, url := range urls {
		go func() {
			resultChannel <- Result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.url] = r.ok
	}

	return results
}
