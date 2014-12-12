package datastore

import "github.com/MattAitchison/akero/model"

type ReceiptLinestore interface {
	GetReceiptLine(id int64) (*model.ReceiptLine, error)
	GetReceiptLineList() ([]*model.ReceiptLine, error)
	PostReceiptLine(line *model.ReceiptLine) error
	PutReceiptLine(line *model.ReceiptLine) error
	DelReceiptLine(line *model.ReceiptLine) error
}
