package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chat-funnel/database"
	"github.com/chat-funnel/models"
	"github.com/julienschmidt/httprouter"
)

//GetContactsHandler - retrieve contacts from the database
func GetContactsHandler(db database.Database) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		c, err := db.GetContacts()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		res := models.Contacts{
			Contacts: c,
		}
		bytes, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintln(w, err)
			return
		}
		w.Write([]byte(bytes))
	}
}
