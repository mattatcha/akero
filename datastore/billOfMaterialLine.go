package datastore

import "github.com/MattAitchison/akero/model"

type BOMLinestore interface {
	GetBOMLine(id int64) (*model.BOMLine, error)
	GetBOMLineList() ([]*model.BOMLine, error)
	GetBOMLineForList(bomID int64) ([]*model.BOMLine, error)
	PostBOMLine(line *model.BOMLine) error
	PutBOMLine(line *model.BOMLine) error
	DelBOMLine(line *model.BOMLine) error
}
