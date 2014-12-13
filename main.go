package main

import (
	"fmt"

	"github.com/MattAitchison/akero/datastore/database"
	_ "github.com/lib/pq"
)

func main() {

	db := database.MustConnect("postgres", "Matt@localhost/akero_test?sslmode=disable")

	ds := database.NewDatastore(db)

	xs, _ := ds.GetTenantList()
	fmt.Printf("%#v", xs)

}
