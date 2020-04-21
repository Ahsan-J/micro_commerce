package db

import (
	"cart/helper"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// ConnectToDatabse established connection to database
func ConnectToDatabse() *sqlx.DB {
	config := helper.GetConfiguration()
	db, err := sqlx.Connect("postgres", config.PostgresURL)
	helper.FailOnError(err, "Failed connecting to Database")
	fmt.Println("Connected to database ", config.PostgresURL)
	return db
}
