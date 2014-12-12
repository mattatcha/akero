package datastore

import "github.com/MattAitchison/akero/model"

type Transactionstore interface {
	GetTransaction(id int64) (*model.Transaction, error)
	GetTransactionList() ([]*model.Transaction, error)
	GetTransactionItemList(itemID int64) ([]*model.Transaction, error)
	PostTransaction(transaction *model.Transaction) error
	PutTransaction(transaction *model.Transaction) error
	DelTransaction(transaction *model.Transaction) error
}
