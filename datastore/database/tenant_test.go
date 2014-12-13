package database

import (
	"testing"

	"github.com/MattAitchison/akero/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTenant(t *testing.T) {
	db := mustConnectTest()
	ts := NewTenantstore(db)
	defer db.Close()
	Convey("Tenantstore", t, func() {
		// Delete all before each test.
		db.Exec("DELETE FROM tenant")
		Convey("Should Post a Tenant", func() {
			err := ts.PostTenant(&model.Tenant{
				Name:        "Lanciv",
				Description: "Lanciv LLC",
			})
			So(err, ShouldBeNil)
		})

		Convey("Should Put a Tenant", nil)

		Convey("Should Delete a Tenant", nil)

		Convey("Should Get a Tenant", func() {
			tenant := &model.Tenant{
				Name:        "Lanciv",
				Description: "Lanciv LLC",
			}
			ts.PostTenant(tenant)

			tenant2, err := ts.GetTenant(tenant.ID)

			So(err, ShouldBeNil)
			So(tenant2.Name, ShouldEqual, tenant.Name)

		})

		Convey("Should Get a Tenant List", func() {
			ts.PostTenant(&model.Tenant{
				Name:        "Lanciv",
				Description: "Lanciv LLC",
			})
			res, err := ts.GetTenantList()

			So(err, ShouldBeNil)
			So(len(res), ShouldEqual, 1)
		})

	})

}
