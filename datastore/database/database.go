package database

import "github.com/MattAitchison/akero/datastore"

func NewDatastore() datastore.Datastore {

	return struct {
		*Tennantstore
	}{
		NewTennantstore(),
	}
}
