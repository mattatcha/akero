package datastore

import "github.com/MattAitchison/akero/model"

type Purchasestore interface {
	GetPurchase(id int64) (*model.Purchase, error)
	GetPurchaseList() ([]*model.Purchase, error)
	GetPurchaseApprovedList() ([]*model.Purchase, error)
	GetPurchaseDraftList() ([]*model.Purchase, error)
	PostPurchase(purchase *model.Purchase) error
	PutPurchase(purchase *model.Purchase) error
	DelPurchase(purchase *model.Purchase) error
}
