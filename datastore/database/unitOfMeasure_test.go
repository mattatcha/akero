package database

import (
	"testing"

	"github.com/MattAitchison/akero/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitOfMeasure(t *testing.T) {
	db := mustConnectTest()
	ts := NewUnitOfMeasurestore(db)
	defer db.Close()
	Convey("UnitOfMeasurestore", t, func() {
		// Delete all before each test.
		db.Exec("DELETE FROM unit_of_measure")
		Convey("Should Post a UnitOfMeasure", func() {
			err := ts.PostUnitOfMeasure(&model.UnitOfMeasure{
				Name:        "EA",
				Description: "EA",
			})
			So(err, ShouldBeNil)
		})

		Convey("Should Put a UnitOfMeasure", nil)

		Convey("Should Delete a UnitOfMeasure", nil)

		Convey("Should Get a UnitOfMeasure", func() {
			unitOfMeasure := &model.UnitOfMeasure{
				Name:        "EA",
				Description: "EA",
			}
			ts.PostUnitOfMeasure(unitOfMeasure)

			tenant2, err := ts.GetUnitOfMeasure(unitOfMeasure.ID)

			So(err, ShouldBeNil)
			So(tenant2.Name, ShouldEqual, unitOfMeasure.Name)

		})

		Convey("Should Get a UnitOfMeasure List", func() {
			ts.PostUnitOfMeasure(&model.UnitOfMeasure{
				Name:        "EA",
				Description: "EA",
			})
			res, err := ts.GetUnitOfMeasureList()

			So(err, ShouldBeNil)
			So(len(res), ShouldEqual, 1)
		})

	})

}
