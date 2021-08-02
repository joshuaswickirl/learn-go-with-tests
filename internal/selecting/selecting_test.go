package selecting_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/selecting"
)

func TestWebsiteRacer(t *testing.T) {

	t.Run("returns faster request url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Second)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.CloseClientConnections()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := selecting.WebsiteRacer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns error after 10s", func(t *testing.T) {
		server := makeDelayedServer(11 * time.Second)
		defer server.CloseClientConnections() // force close unneeded client connection

		_, err := selecting.ConfigurableRacer(server.URL, server.URL, 1*time.Millisecond)

		if err == nil {
			t.Error("expected error and didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
