package datastore

import "github.com/MattAitchison/akero/model"

type Itemstore interface {
	GetItem(id int64) (*model.Item, error)
	GetItemList() ([]*model.Item, error)
	GetItemAssemblyList() ([]*model.Item, error)
	GetItemComponentList() ([]*model.Item, error)
	PostItem(item *model.Item) error
	PutItem(item *model.Item) error
	DelItem(item *model.Item) error
}
