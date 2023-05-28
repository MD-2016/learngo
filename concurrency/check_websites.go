package concurrency

type websiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc websiteChecker, urls []string) map[string]bool {
	res := make(map[string]bool)
	reschan := make(chan result)

	for _, url := range urls {
		go func(u string) {
			reschan <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-reschan
		res[r.string] = r.bool
	}

	return res
}
