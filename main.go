package main

import (
	"acc/cmd"
	"acc/db"
	"os"
)

func main() {
	var err error
	if err = db.SetUp("."); err != nil {
		handleError(err)
	}
	if err = cmd.Execute(); err != nil {
		handleError(err)
	}
}

func handleError(e error) {
	os.Exit(1)
}
