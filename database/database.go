package database

import (
	"database/sql"

	"github.com/chat-funnel/models"
	_ "github.com/go-sql-driver/mysql" // blank import to load mysql driver
)

type database struct {
	DB *sql.DB
}

const NotFound = "NOT_FOUND"

// Database interface encapsulates all database access
type Database interface {
	GetContact(id int64) (*models.Contact, error)
	GetContacts() ([]*models.Contact, error)
	AddContact(contact models.Contact) (int64, error)
}

//New - get a new connection to the database
func New(connectionString string) (Database, error) {
	dbConn, err := sql.Open("mysql", connectionString) // This will need to change from being hard coded
	if err != nil {
		return nil, err
	}

	var db database
	db.DB = dbConn
	return &db, nil
}

func (db *database) GetContacts() ([]*models.Contact, error) {

	rows, err := db.DB.Query("SELECT id, first_name, last_name, email FROM contacts")

	// .Scan(&firstName, &lastName, &email)
	if err != nil {
		return nil, err
	}
	var contacts []*models.Contact
	for rows.Next() {
		var (
			id        int64
			firstName *string
			lastName  *string
			email     *string
		)
		rows.Scan(&id, &firstName, &lastName, &email)
		c := models.Contact{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		}
		contacts = append(contacts, &c)

	}

	return contacts, err
}

func (db *database) GetContact(id int64) (*models.Contact, error) {
	var (
		firstName *string
		lastName  *string
		email     *string
	)

	err := db.DB.QueryRow("SELECT first_name, last_name, email FROM contacts WHERE id = ?", id).Scan(&firstName, &lastName, &email)
	if err != nil {
		return nil, err
	}
	c := models.Contact{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	return &c, err
}

func (db *database) AddContact(contact models.Contact) (int64, error) {
	res, err := db.DB.Exec(`INSERT contacts SET
	 first_name = ?,
	 last_name = ?,
	 email = ? `,
		contact.FirstName,
		contact.LastName,
		contact.Email,
	)
	if err != nil {
		return 0, nil
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, err
}
