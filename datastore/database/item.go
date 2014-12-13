package database

import (
	"errors"

	"github.com/MattAitchison/akero/model"
	"github.com/jmoiron/sqlx"
)

type Itemstore struct {
	*sqlx.DB
}

func NewItemstore(db *sqlx.DB) *Itemstore {
	return &Itemstore{db}
}

func (db *Itemstore) GetItem(id int64) (*model.Item, error) {
	item := &model.Item{}

	err := db.Get(item, `SELECT * FROM item WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (db *Itemstore) GetItemList() ([]*model.Item, error) {
	return db.getItemListRaw("")
}
func (db *Itemstore) GetItemAssemblyList() ([]*model.Item, error) {
	return db.getItemListRaw("WHERE assembly = true")
}
func (db *Itemstore) GetItemComponentList() ([]*model.Item, error) {
	return db.getItemListRaw("WHERE component = true")
}
func (db *Itemstore) getItemListRaw(qAppend string) ([]*model.Item, error) {
	contacts := []*model.Item{}

	query := `SELECT * FROM item `
	if len(qAppend) != 0 {
		query = query + qAppend
	}

	err := db.Select(&contacts, query)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (db *Itemstore) PostItem(item *model.Item) error {

	rows, err := db.NamedQuery(`INSERT INTO item
		(code, description, comments, component, assembly, item_group_id, unit_of_measure_id)
		VALUES (:code, :description, :comments, :component, :assembly, :item_group_id, :unit_of_measure_id)
		RETURNING id`, item)
	if err != nil {
		return err
	}

	if rows.Next() {
		rows.Scan(&item.ID)
	}
	return nil
}
func (db *Itemstore) PutItem(item *model.Item) error {
	return errors.New("Not Impl")
}

func (db *Itemstore) DelItem(item *model.Item) error {
	return errors.New("Not Impl")
}
