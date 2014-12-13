package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/MattAitchison/akero/datastore"
	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate/migrate"

	_ "github.com/lib/pq"
)

func NewDatastore(db *sqlx.DB) datastore.Datastore {

	return struct {
		*Tenantstore
		*Contactstore
		*ItemGroupstore
		*UnitOfMeasurestore
		*Itemstore
	}{
		NewTenantstore(db),
		NewContactstore(db),
		NewItemGroupstore(db),
		NewUnitOfMeasurestore(db),
		NewItemstore(db),
	}
}

// Connect is a helper function that establishes a new
// database connection and auto-generates the database
// schema. If the database already exists, it will perform
// and update as needed.
func Connect(driver, dsn string) (*sqlx.DB, error) {
	url := fmt.Sprintf("%s://%s", driver, dsn)

	// use synchronous versions of migration functions ...
	allErrs, ok := migrate.UpSync(url, "../migrate")
	if !ok {
		log.Println(allErrs)
		return nil, errors.New("Migration Error")
	}

	return sqlx.Open(driver, url)
}

// MustConnect is a helper function that creates a
// new database connection and auto-generates the
// database schema. An error causes a panic.
func MustConnect(driver, dsn string) *sqlx.DB {
	db, err := Connect(driver, dsn)
	if err != nil {
		panic(err)
	}
	return db
}

// mustConnectTest is a helper function that creates a
// new database connection using environment variables.
func mustConnectTest() *sqlx.DB {
	var (
		driver = os.Getenv("TEST_DRIVER")
		dsn    = os.Getenv("TEST_DSN")
	)
	if len(driver) == 0 {
		driver = "postgres"
		dsn = "Matt@localhost/akero_test?sslmode=disable"
	}

	migrate.DownSync(driver+"://"+dsn, "../migrate")

	db, err := Connect(driver, dsn)
	if err != nil {
		panic(err)
	}
	return db
}
