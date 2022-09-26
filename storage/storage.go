package storage

import (
	"boarder/models"
	"boarder/networking"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/boltdb/bolt"
)

func SetupDB() (*bolt.DB, error) {
	db, err := bolt.Open("data.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("THREAD"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte("POST"))
		if err != nil {
			return fmt.Errorf("could not create weight bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	return db, nil
}

func Update_db(db *bolt.DB, board string, threads []int) {
	var wg sync.WaitGroup
	for i := 1; i < len(threads); i++ {
		wg.Add(1)
        i := i
		go func(wg *sync.WaitGroup) {
			posts := networking.Get_posts_from_threads(board, threads[i])
            err := store_data(db, board, threads[i], posts)
			if err != nil {
				panic(err)
			}
            wg.Done()
		}(&wg)
	}
    wg.Wait()
}

func store_data(db *bolt.DB, board string, thread int, posts []models.Post) error {
	existing_post_id_list, err := get_existing_post_id_list(db, thread)
	if err != nil {
		return fmt.Errorf("cannot view existing key: %v", err)
	}
	var new_post_id_list []int
	for i := 0; i < len(posts); i++ {
		new_post_id_list = append(new_post_id_list, posts[i].No)
	}
	updated_post_id_list := update_posts_list(existing_post_id_list, new_post_id_list)
	new_value, err := json.Marshal(updated_post_id_list)
	if err != nil {
		return fmt.Errorf("could not marshal data: %v", err)
	}
	err = update_bucket(db, "THREAD", []byte(fmt.Sprint(thread)), new_value)
	if err != nil {
		return fmt.Errorf("could not update post bucket: %v", err)
	}
	var wg sync.WaitGroup
	for i := 0; i < len(posts); i++ {
		wg.Add(1)
		i := i
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			post_data, err := json.Marshal(posts[i])
			if err != nil {
				panic(err)
			}
			err = update_bucket(db, "POST", []byte(fmt.Sprint(board, "_", posts[i].No)), post_data)
			if err != nil {
				panic(err)
			}
		}(&wg)
	}
	wg.Wait()
	return err
}

func get_existing_post_id_list(db *bolt.DB, thread int) ([]int, error) {
	var existing_post_id_list []int
	err := db.View(func(tx *bolt.Tx) error {
		existing_post_id_bytes := tx.Bucket([]byte("THREAD")).Get([]byte(fmt.Sprint(thread)))
		if existing_post_id_bytes != nil {
			err := json.Unmarshal(existing_post_id_bytes, &existing_post_id_list)
			if err != nil {
				return fmt.Errorf("could not parse data from threads table, %v", err)
			}
		}
		return nil
	})
	return existing_post_id_list, err
}

func update_posts_list(existing_post_id_list []int, new_post_id_list []int) []int {
	var updated_post_id_list = existing_post_id_list
	for i := 0; i < len(new_post_id_list); i++ {
		post_exists := false
		for j := 0; j < len(updated_post_id_list); j++ {
			if updated_post_id_list[j] == new_post_id_list[i] {
				post_exists = true
				break
			}
		}
		if !post_exists {
			updated_post_id_list = append(updated_post_id_list, new_post_id_list[i])
		}
	}
	return updated_post_id_list
}

func update_bucket(db *bolt.DB, bucket string, key []byte, value []byte) error {
	err := db.Batch(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte(bucket)).Put(key, value)
		if err != nil {
			return fmt.Errorf("could not set key and value: %v", err)
		}
		return nil
	})
	return err
}

func read_threads(db *bolt.DB) {
    err := db.View(func(tx *bolt.Tx) error {
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
