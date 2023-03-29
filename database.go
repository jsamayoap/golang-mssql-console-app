package main

import (
	"database/sql"
	"fmt"
)

var database *sql.DB

var (
	port     = 1433
	password = "misDatos2023!"
	user     = "SA"
)

var connectionString = fmt.Sprintf("user id=%s;password=%s;port=%d", user, password, port)
