package database

import (
	"testing"

	"github.com/MattAitchison/akero/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestItemGroup(t *testing.T) {
	db := mustConnectTest()
	ts := NewItemGroupstore(db)
	defer db.Close()
	Convey("ItemGroupstore", t, func() {
		// Delete all before each test.
		db.Exec("DELETE FROM item_group")
		Convey("Should Post a ItemGroup", func() {
			err := ts.PostItemGroup(&model.ItemGroup{
				Name:        "Oil",
				Description: "Oil",
			})
			So(err, ShouldBeNil)
		})

		Convey("Should Put a ItemGroup", nil)

		Convey("Should Delete a ItemGroup", nil)

		Convey("Should Get a ItemGroup", func() {
			itemGroup := &model.ItemGroup{
				Name:        "Oil",
				Description: "Oil",
			}
			ts.PostItemGroup(itemGroup)

			tenant2, err := ts.GetItemGroup(itemGroup.ID)

			So(err, ShouldBeNil)
			So(tenant2.Name, ShouldEqual, itemGroup.Name)

		})

		Convey("Should Get a ItemGroup List", func() {
			ts.PostItemGroup(&model.ItemGroup{
				Name:        "Oil",
				Description: "Oil",
			})
			res, err := ts.GetItemGroupList()

			So(err, ShouldBeNil)
			So(len(res), ShouldEqual, 1)
		})

	})

}
