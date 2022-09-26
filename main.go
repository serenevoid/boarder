package main

import (
	"boarder/storage"
	"boarder/tui"
)

func main() {
	db, err := storage.SetupDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
    tui.Setup_UI()
}

