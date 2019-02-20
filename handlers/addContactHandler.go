package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chat-funnel/database"
	"github.com/chat-funnel/models"
	"github.com/julienschmidt/httprouter"
)

//AddContactHandler - add a contact to the database
func AddContactHandler(db database.Database) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		buf, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		var contact models.Contact
		err := json.Unmarshal(buf, &contact)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}

		id, err := db.AddContact(contact)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		idResponse := models.ID{
			ID: id,
		}
		bytes, err := json.Marshal(idResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintln(w, err)
			return
		}
		w.Write([]byte(bytes))
	}
}
