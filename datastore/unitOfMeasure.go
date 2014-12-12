package datastore

import "github.com/MattAitchison/akero/model"

type UnitOfMeasurestore interface {
	GetUnitOfMeasure(id int64) (*model.UnitOfMeasure, error)
	GetUnitOfMeasureList() ([]*model.UnitOfMeasure, error)
	PostUnitOfMeasure(uom *model.UnitOfMeasure) error
	PutUnitOfMeasure(uom *model.UnitOfMeasure) error
	DelUnitOfMeasure(uom *model.UnitOfMeasure) error
}
