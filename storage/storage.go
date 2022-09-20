package storage

import (
	"boarder/models"
	"encoding/json"
	"fmt"

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
	fmt.Println("DB Setup Done")
	return db, nil
}

func Store_data(db *bolt.DB, board string, thread int, posts []models.Post) error {
    existing_post_id_list := get_existing_post_id_list(db, thread)
	var new_post_id_list []int
	for i := 0; i < len(posts); i++ {
		new_post_id_list = append(new_post_id_list, posts[i].No)
	}

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
	new_value, err := json.Marshal(updated_post_id_list)
	err = db.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("THREAD")).Put([]byte(fmt.Sprint(thread)), new_value)
		if err != nil {
			return fmt.Errorf("could not set config: %v", err)
		}
		return nil
	})
	for i := 0; i < len(posts); i++ {
		post_data, err := json.Marshal(posts[i])
		err = db.Update(func(tx *bolt.Tx) error {
			err = tx.Bucket([]byte("POST")).Put([]byte(fmt.Sprint(board, "_", posts[i].No)), post_data)
			if err != nil {
				return fmt.Errorf("could not set config: %v", err)
			}
			return nil
		})
	}
	return err
}

func get_existing_post_id_list(db *bolt.DB, thread int) []int {
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
	if err != nil {
		return fmt.Errorf("cannot view existing key: %v", err)
	}
}
