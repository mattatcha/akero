package datastore

import "github.com/MattAitchison/akero/model"

type BillOfMaterialstore interface {
	GetBillOfMaterial(id int64) (*model.BillOfMaterial, error)
	GetBillOfMaterialList() ([]*model.BillOfMaterial, error)
	PostBillOfMaterial(bom *model.BillOfMaterial) error
	PutBillOfMaterial(bom *model.BillOfMaterial) error
	DelBillOfMaterial(bom *model.BillOfMaterial) error
}
