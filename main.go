package main

import (
	"github.com/MattAitchison/akero/datastore/database"
	_ "github.com/lib/pq"
)

func main() {

	ts := database.NewTennantstore()
	ts.GetTennant(10)
}
