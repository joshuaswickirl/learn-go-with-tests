package contexting

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return // todo: log error
		}

		fmt.Fprint(w, data)
	}
}
