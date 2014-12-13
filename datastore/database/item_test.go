package database

import (
	"testing"

	"github.com/MattAitchison/akero/model"
	. "github.com/smartystreets/goconvey/convey"
)

func createItem(is *Itemstore, ig *ItemGroupstore, um *UnitOfMeasurestore) (*model.Item, error) {
	itemGroup := &model.ItemGroup{
		Name:        "Misc",
		Description: "Misc Items",
	}

	err := ig.PostItemGroup(itemGroup)
	if err != nil {
		return nil, err
	}
	uom := &model.UnitOfMeasure{
		Name:        "EA",
		Description: "Each",
	}

	err = um.PostUnitOfMeasure(uom)
	if err != nil {
		return nil, err
	}

	item := &model.Item{
		ItemGroupID:     itemGroup.ID,
		UnitOfMeasureID: uom.ID,
		Code:            "BOTTLE",
	}

	err = is.PostItem(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
func TestItem(t *testing.T) {
	db := mustConnectTest()
	is := NewItemstore(db)
	ig := NewItemGroupstore(db)
	um := NewUnitOfMeasurestore(db)
	defer db.Close()

	Convey("Itemstore", t, func() {
		// Delete all before each test.
		db.Exec("DELETE FROM item")
		db.Exec("DELETE FROM item_group")
		db.Exec("DELETE FROM unit_of_measure")
		Convey("Should Post a Item", func() {
			_, err := createItem(is, ig, um)
			So(err, ShouldBeNil)
		})

		Convey("Should Put a Item", nil)

		Convey("Should Delete a Item", nil)

		Convey("Should Get a Item", func() {
			item, err := createItem(is, ig, um)
			So(err, ShouldBeNil)
			item2, err := is.GetItem(item.ID)

			So(err, ShouldBeNil)
			So(item2.Code, ShouldEqual, item.Code)

		})

		Convey("Should Get a Item List", func() {
			createItem(is, ig, um)
			res, err := is.GetItemList()

			So(err, ShouldBeNil)
			So(len(res), ShouldEqual, 1)
		})
	})

}
