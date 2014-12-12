package datastore

import "github.com/MattAitchison/akero/model"

type ItemGroupstore interface {
	GetItemGroup(id int64) (*model.ItemGroup, error)
	GetItemGroupList() ([]*model.ItemGroup, error)
	PostItemGroup(group *model.ItemGroup) error
	PutItemGroup(group *model.ItemGroup) error
	DelItemGroup(group *model.ItemGroup) error
}
