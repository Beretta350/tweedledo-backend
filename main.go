package main

import (
	"os"

	"github.com/tweedledo/config/db"
)

func main() {
	db.ConnectDB(os.Getenv("env"))
}
