package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chat-funnel/database"
	"github.com/chat-funnel/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := database.New(os.Getenv("DB_CONN"))
	if err != nil {
		fmt.Printf("failed to initialize database: %s", err.Error())
		os.Exit(1)
	}
	router := httprouter.New()
	router.GET("/contacts/:id", handlers.GetContactHandler(db))
	router.GET("/contacts", handlers.GetContactsHandler(db))
	router.POST("/contacts", handlers.AddContactHandler(db))

	log.Fatal(http.ListenAndServe(":8080", router))
}
