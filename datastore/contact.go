package datastore

import "github.com/MattAitchison/akero/model"

type Contactstore interface {
	GetContact(id int64) (*model.Contact, error)
	GetContactList() ([]*model.Contact, error)
	GetContactVendorList() ([]*model.Contact, error)
	GetContactSupplierList() ([]*model.Contact, error)
	PostContact(contact *model.Contact) error
	PutContact(contact *model.Contact) error
	DelContact(contact *model.Contact) error
}
