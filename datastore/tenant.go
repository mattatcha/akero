package datastore

import "github.com/MattAitchison/akero/model"

type Tenantstore interface {
	GetTenant(id int64) (*model.Tenant, error)
	GetTenantList() ([]*model.Tenant, error)
	PostTenant(tenant *model.Tenant) error
	PutTenant(tenant *model.Tenant) error
	DelTenant(tenant *model.Tenant) error
}
