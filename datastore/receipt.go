package datastore

import "github.com/MattAitchison/akero/model"

type Receiptstore interface {
	GetReceipt(id int64) (*model.Receipt, error)
	GetReceiptList() ([]*model.Receipt, error)
	GetReceiptPurchaseList(purchaseID int64) ([]*model.Receipt, error)
	PostReceipt(receipt *model.Receipt) error
	PutReceipt(receipt *model.Receipt) error
	DelReceipt(receipt *model.Receipt) error
}
