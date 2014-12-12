package datastore

import "github.com/MattAitchison/akero/model"

type Tennantstore interface {
	GetTennant(id int64) (*model.Tennant, error)
	GetTennantList() ([]*model.Tennant, error)
	PostTennant(tennant *model.Tennant) error
	PutTennant(tennant *model.Tennant) error
	DelTennant(tennant *model.Tennant) error
}
