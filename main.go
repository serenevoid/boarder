package main

import (
	"boarder/networking"
	"boarder/storage"
	"fmt"
	"sync"

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
	var wg sync.WaitGroup
	for i := 1; i < 5; i++ {
		wg.Add(1)
        i := i
		go func(wg *sync.WaitGroup) {
			posts := networking.Get_posts_from_threads(board, threads[i])
			err = storage.Store_data(db, board, threads[i], posts)
			if err != nil {
				panic(err)
			}
            wg.Done()
		}(&wg)
	}
    fmt.Println("reached after for loop")
    wg.Wait()
    fmt.Println("reached after wait")
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
