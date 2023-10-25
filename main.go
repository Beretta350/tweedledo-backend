package main

import (
	"os"

	"github.com/tweedledo/infrastructure/db"
)

func main() {
	db.ConnectDB(os.Getenv("env"))
}
