package player

import (
	"fmt"
	"net/http"
)

func mock(s string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from %s\n", s)
	}
}
