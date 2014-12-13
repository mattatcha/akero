package database

import (
	"errors"

	"github.com/MattAitchison/akero/model"
	"github.com/jmoiron/sqlx"
)

type ItemGroupstore struct {
	*sqlx.DB
}

func NewItemGroupstore(db *sqlx.DB) *ItemGroupstore {
	return &ItemGroupstore{db}
}

func (db *ItemGroupstore) GetItemGroup(id int64) (*model.ItemGroup, error) {
	//sql, _, _ := psql.Select("*").From(tenantsT).Where(sq.Eq{"id": id}).ToSql()

	itemGroup := &model.ItemGroup{}
	err := db.Get(itemGroup, `SELECT * FROM item_group WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return itemGroup, nil
}

func (db *ItemGroupstore) GetItemGroupList() ([]*model.ItemGroup, error) {
	tenants := []*model.ItemGroup{}

	err := db.Select(&tenants, `SELECT * FROM item_group`)
	if err != nil {
		return nil, err
	}
	return tenants, nil
}
func (db *ItemGroupstore) PostItemGroup(itemGroup *model.ItemGroup) error {

	rows, err := db.NamedQuery(`INSERT INTO item_group (name,description) VALUES (:name,:description) RETURNING id`,
		itemGroup)
	if err != nil {
		return err
	}

	if rows.Next() {
		rows.Scan(&itemGroup.ID)
	}
	return nil
}
func (db *ItemGroupstore) PutItemGroup(itemGroup *model.ItemGroup) error {
	return errors.New("Not Impl")
}
func (db *ItemGroupstore) DelItemGroup(itemGroup *model.ItemGroup) error {
	return errors.New("Not Impl")
}
