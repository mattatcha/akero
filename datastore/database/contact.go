package database

import (
	"errors"

	"github.com/MattAitchison/akero/model"
	"github.com/jmoiron/sqlx"
)

type Contactstore struct {
	*sqlx.DB
}

func NewContactstore(db *sqlx.DB) *Contactstore {
	return &Contactstore{db}
}

func (db *Contactstore) GetContact(id int64) (*model.Contact, error) {
	contact := &model.Contact{}

	err := db.Get(contact, `SELECT * FROM contact WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (db *Contactstore) GetContactList() ([]*model.Contact, error) {
	return db.getContactListRaw("")
}
func (db *Contactstore) getContactListRaw(qAppend string) ([]*model.Contact, error) {
	contacts := []*model.Contact{}

	query := `SELECT * FROM contact `
	if len(qAppend) != 0 {
		query = query + qAppend
	}

	err := db.Select(&contacts, query)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}
func (db *Contactstore) GetContactVendorList() ([]*model.Contact, error) {
	return db.getContactListRaw("WHERE vendor = true")
}
func (db *Contactstore) GetContactSupplierList() ([]*model.Contact, error) {
	return db.getContactListRaw("WHERE supplier = true")
}

func (db *Contactstore) PostContact(contact *model.Contact) error {

	rows, err := db.NamedQuery(`INSERT INTO contact (name, vendor, supplier) VALUES (:name,:vendor,:supplier) RETURNING id`,
		contact)
	if err != nil {
		return err
	}

	if rows.Next() {
		rows.Scan(&contact.ID)
	}
	return nil
}
func (db *Contactstore) PutContact(contact *model.Contact) error {
	return errors.New("Not Impl")
}

func (db *Contactstore) DelContact(contact *model.Contact) error {
	return errors.New("Not Impl")
}
