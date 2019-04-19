package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RespondJSON is a function that sets response as JSON
func RespondJSON(i interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	j, err := json.Marshal(i)
	if err != nil {
		fmt.Printf("Failed to generate JSON: %s\n", err)
		return
	}

	w.Write(j)
}
