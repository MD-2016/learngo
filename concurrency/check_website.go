package concurrency

import "net/http"

func CheckWebsite(url string) bool {
	res, err := http.Head(url)
	if err != nil {
		return false
	}

	return res.StatusCode == http.StatusOK
}
