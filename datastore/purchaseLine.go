package datastore

import "github.com/MattAitchison/akero/model"

type PurchaseLinestore interface {
	GetPurchaseLine(id int64) (*model.PurchaseLine, error)
	GetPurchaseLineList() ([]*model.PurchaseLine, error)
	GetPurchaseLineForList(purchaseID int64) ([]*model.PurchaseLine, error)
	PostPurchaseLine(line *model.PurchaseLine) error
	PutPurchaseLine(line *model.PurchaseLine) error
	DelPurchaseLine(line *model.PurchaseLine) error
}
