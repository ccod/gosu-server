package history

import (
	"fmt"
	"net/http"
	"strconv"

	m "github.com/ccod/gosu-server/pkg/middleware"
	"github.com/ccod/gosu-server/pkg/models"
	re "github.com/ccod/gosu-server/pkg/response"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func mock(s string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from %s\n", s)
	}
}

func getHistory(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	history := models.GetHistory(db, id)
	re.RespondJSON(history, w, r)
}
