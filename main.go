package main

import (
	"os"

	"github.com/tweedledo/configs/db"
)

func main() {
	db.ConnectDB(os.Getenv("env"))
}
