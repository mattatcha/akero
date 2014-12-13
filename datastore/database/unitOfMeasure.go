package database

import (
	"errors"

	"github.com/MattAitchison/akero/model"
	"github.com/jmoiron/sqlx"
)

type UnitOfMeasurestore struct {
	*sqlx.DB
}

func NewUnitOfMeasurestore(db *sqlx.DB) *UnitOfMeasurestore {
	return &UnitOfMeasurestore{db}
}

func (db *UnitOfMeasurestore) GetUnitOfMeasure(id int64) (*model.UnitOfMeasure, error) {
	//sql, _, _ := psql.Select("*").From(tenantsT).Where(sq.Eq{"id": id}).ToSql()

	unitOfMeasure := &model.UnitOfMeasure{}
	err := db.Get(unitOfMeasure, `SELECT * FROM unit_of_measure WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return unitOfMeasure, nil
}

func (db *UnitOfMeasurestore) GetUnitOfMeasureList() ([]*model.UnitOfMeasure, error) {
	tenants := []*model.UnitOfMeasure{}

	err := db.Select(&tenants, `SELECT * FROM unit_of_measure`)
	if err != nil {
		return nil, err
	}
	return tenants, nil
}
func (db *UnitOfMeasurestore) PostUnitOfMeasure(unitOfMeasure *model.UnitOfMeasure) error {

	rows, err := db.NamedQuery(`INSERT INTO unit_of_measure (name,description) VALUES (:name,:description) RETURNING id`,
		unitOfMeasure)
	if err != nil {
		return err
	}

	if rows.Next() {
		rows.Scan(&unitOfMeasure.ID)
	}
	return nil
}
func (db *UnitOfMeasurestore) PutUnitOfMeasure(unitOfMeasure *model.UnitOfMeasure) error {
	return errors.New("Not Impl")
}
func (db *UnitOfMeasurestore) DelUnitOfMeasure(unitOfMeasure *model.UnitOfMeasure) error {
	return errors.New("Not Impl")
}
