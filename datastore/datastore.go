package datastore

type Datastore interface {
	// BillOfMaterialstore
	// BOMLinestore
	Contactstore
	Itemstore
	ItemGroupstore
	// Purchasestore
	// PurchaseLinestore
	// Receiptstore
	// ReceiptLinestore
	Tenantstore
	// Transactionstore
	UnitOfMeasurestore
}
