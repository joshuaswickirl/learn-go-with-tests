package selecting

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func WebsiteRacer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-makeHTTPRequest(a):
		return a, nil
	case <-makeHTTPRequest(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func makeHTTPRequest(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, err := http.Get(url)
		if err != nil {
			ch <- struct{}{}
		}
		ch <- struct{}{}
	}()
	return ch
}
