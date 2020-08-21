package main

import (
	"github.com/mimiiiz/timescale-testing/database"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()
}
