package database

import (
	"github.com/MattAitchison/akero/model"

	sq "github.com/lann/squirrel"
)

// users := sq.Select("*").From("users").Join("emails USING (email_id)")

const tennantsT = "tennants"

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Tennantstore struct {
}

func NewTennantstore() *Tennantstore {
	return &Tennantstore{}
}

func (db *Tennantstore) GetTennant(id int64) (*model.Tennant, error) {
	sql, _, _ := psql.Select("*").From(tennantsT).Where(sq.Eq{"id": id}).ToSql()

	return nil, nil
}

func (db *Tennantstore) GetTennantList() ([]*model.Tennant, error) {

	return nil, nil
}
func (db *Tennantstore) PostTennant(tennant *model.Tennant) error {
	return nil
}
func (db *Tennantstore) PutTennant(tennant *model.Tennant) error {
	return nil
}
func (db *Tennantstore) DelTennant(tennant *model.Tennant) error {
	return nil
}

// type Tennantstore interface {
// 	GetTennant(id int64) (*model.Tennant, error)
// 	GetTennantList() ([]*model.Tennant, error)
// 	PostTennant(tennant *model.Tennant) error
// 	PutTennant(tennant *model.Tennant) error
// 	DelTennant(tennant *model.Tennant) error
// }
