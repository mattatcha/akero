package database

import (
	"errors"

	"github.com/MattAitchison/akero/model"
	"github.com/jmoiron/sqlx"
)

type Tenantstore struct {
	*sqlx.DB
}

func NewTenantstore(db *sqlx.DB) *Tenantstore {
	return &Tenantstore{db}
}

func (db *Tenantstore) GetTenant(id int64) (*model.Tenant, error) {
	//sql, _, _ := psql.Select("*").From(tenantsT).Where(sq.Eq{"id": id}).ToSql()

	tenant := &model.Tenant{}
	err := db.Get(tenant, `SELECT * FROM tenant WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return tenant, nil
}

func (db *Tenantstore) GetTenantList() ([]*model.Tenant, error) {
	tenants := []*model.Tenant{}

	err := db.Select(&tenants, `SELECT * FROM tenant`)
	if err != nil {
		return nil, err
	}
	return tenants, nil
}
func (db *Tenantstore) PostTenant(tenant *model.Tenant) error {

	rows, err := db.NamedQuery(`INSERT INTO tenant (name,description) VALUES (:name,:description) RETURNING id`,
		tenant)
	if err != nil {
		return err
	}

	if rows.Next() {
		rows.Scan(&tenant.ID)
	}
	return nil
}
func (db *Tenantstore) PutTenant(tenant *model.Tenant) error {
	return errors.New("Not Impl")
}
func (db *Tenantstore) DelTenant(tenant *model.Tenant) error {
	return errors.New("Not Impl")
}
