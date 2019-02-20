package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/chat-funnel/database"
	"github.com/julienschmidt/httprouter"
)

//GetContactHandler - retrieve contact by the given id from the database
func GetContactHandler(db database.Database) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		log.Printf("got id: %s", id)

		c, err := db.GetContact(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		bytes, err := json.Marshal(c)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintln(w, err)
			return
		}
		w.Write([]byte(bytes))
	}
}
