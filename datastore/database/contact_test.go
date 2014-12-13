package database

import (
	"testing"

	"github.com/MattAitchison/akero/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestContact(t *testing.T) {
	db := mustConnectTest()
	ts := NewContactstore(db)
	defer db.Close()
	Convey("Contactstore", t, func() {
		// Delete all before each test.
		db.Exec("DELETE FROM contact")
		Convey("Should Post a Contact", func() {
			err := ts.PostContact(&model.Contact{
				Name: "Lanciv",
			})
			So(err, ShouldBeNil)
		})

		Convey("Should Put a Contact", nil)

		Convey("Should Delete a Contact", nil)

		Convey("Should Get a Contact", func() {
			contact := &model.Contact{
				Name: "Lanciv",
			}
			ts.PostContact(contact)

			contact2, err := ts.GetContact(contact.ID)

			So(err, ShouldBeNil)
			So(contact2.Name, ShouldEqual, contact.Name)

		})

		Convey("Should Get a Contact List", func() {
			ts.PostContact(&model.Contact{
				Name: "Lanciv",
			})
			res, err := ts.GetContactList()

			So(err, ShouldBeNil)
			So(len(res), ShouldEqual, 1)
		})
		Convey("Should Get a Contact Supplier List", nil)
		Convey("Should Get a Contact Vendor List", nil)

	})

}
