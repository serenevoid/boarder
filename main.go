package main

import (
	"boarder/networking"
	"boarder/storage"
	"fmt"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := storage.SetupDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	board := networking.Get_user_input()
	threads := networking.Get_threads_from_board(board)
    for i := 0; i < 5; i++ {
		posts := networking.Get_posts_from_threads(board, threads[i])
		err = storage.Store_data(db, board, threads[i], posts)
		if err != nil {
			panic(err)
		}
	}
	err = db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("THREAD"))

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})
	if err != nil {
		panic(err)
	}
}
