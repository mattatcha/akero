package model

import "time"

//Tennant a user of our service
type Tennant struct {
	ID          int64
	Name        string
	Description string
}

// UnitOfMeasure or UOM for short
// describes a measurement type for
// an item.
type UnitOfMeasure struct {
	ID          int64
	TennantID   int64
	Name        string
	Description string
}

// ItemGroup is used to group similar products together.
type ItemGroup struct {
	ID          int64
	TennantID   int64
	Name        string
	Description string
}

// TransactionType is used to record what type a transaction is.
// I will need to work on getting a more thought out list.
// Maybe store in DB?
// type TransactionType struct {
// 	ID          int64
// 	TennantID   int64
// 	Name        string
// 	Description string
// }

// TransactionType is used to record what type a transaction is.
// I will need to work on getting a more thought out list.
// Maybe store in DB?
type TransactionType int

const (
	PurchaseType TransactionType = 1 << iota
	SaleType
)

// Contact could be a Vendor or a Supplier (or even both)
type Contact struct {
	ID        int64
	TennantID int64
	Name      string
	Vendor    bool
	Supplier  bool
}

// Item is a part or an end product
type Item struct {
	ID              int64
	TennantID       int64
	ItemGroupID     int64
	ItemGroup       ItemGroup
	Code            string
	Description     string
	UnitOfMeasureID int64 // Should be a Type
	UnitOfMeasure   UnitOfMeasure
	Comments        string
	Component       bool
	Assembly        bool
}

// BillOfMaterial is a list of items can be assembled to
// create an end product.
type BillOfMaterial struct {
	ID        int64
	TennantID int64
	ItemID    int64

	//
	Item     Item
	BOMLines []BOMLine
}

// BOMLine line item for BilBillOfMaterial
type BOMLine struct {
	ID               int64
	TennantID        int64
	BillOfMaterialID int64
	ItemID           int64
	Item             Item
	UsedQty          int64
	WasteQty         int64
}

// Transaction always involves an Item and is the
//  source of truth for Inventory counts and just
// about everyting else that deals with an item.
type Transaction struct {
	ID        int64
	TennantID int64
	ItemID    int64
	Item      Item
	Qty       int64
	Comments  string
	Date      time.Time
	TypeID    int64
	Type      TransactionType
}

// Purchase is a document that commits to purchasing items.
type Purchase struct {
	ID         int64
	TennantID  int64
	SupplierID int64

	Date       time.Time
	Approved   bool
	ApprovedBy string // Should be a persons name
	ApprovedOn time.Time

	//
	PurchaseLines []PurchaseLine
}

// PurchaseLine describes line for a Purchase
type PurchaseLine struct {
	ID        int64
	TennantID int64
	ItemID    int64
	Item      Item
	Qty       int64
	UnitCost  int64
}

// Receipt is used to receieve
// items in that were on a PO.
type Receipt struct {
	ID         int64
	TennantID  int64
	PurchaseID int64
	Purchase   Purchase
	Date       time.Time

	ReceiptLines []ReceiptLine
}

// ReceiptLine defines a line item for a Receipt
type ReceiptLine struct {
	ID        int64
	TennantID int64
	ReceiptID int64
	ItemID    int64
	Item      Item
	Qty       int64
}
