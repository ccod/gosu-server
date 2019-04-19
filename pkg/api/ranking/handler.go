package ranking

import (
	"fmt"
	"net/http"
)

func mock(s string) http.HanlderFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ranking route emits: %s", s)
	}
}

func list(w http.ResponseWriter, r *http.Request)    {}
func index(w http.ResponseWriter, r *http.Request)   {}
func create(w http.ResponseWriter, r *http.Request)  {}
func update(w http.ResponseWriter, r *http.Request)  {}
func delete(w http.ResponseWriter, r *http.Request)  {}
func promote(w http.ResponseWriter, r *http.Request) {}
