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

// RespondError takes and error and writes a json map with error inside as well as the 500 status code
func RespondError(err error, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-type", "application/json")
	d := map[string]string{
		"status": "500",
		"error":  err.Error(),
	}

	j, _ := json.Marshal(d)
	w.Write(j)
}
