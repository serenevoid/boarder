package main

import (
	"boarder/models"
	"boarder/networking"
	"boarder/storage"
	"boarder/util"
	"fmt"
)

func main() {

	var file_list []models.File

	fmt.Println("BOARDER V0.1.0")

	entry_list, err := util.Load_config()
	if err != nil {
		fmt.Print("Error: ", err)
		return
	}

	fmt.Println("Collecting thread data...")
	for _, entry := range entry_list {
		posts, err := networking.Get_posts_from_thread(entry)
		if err != nil {
			fmt.Print("Error: ", err)
			return
		}
        if posts == nil {
            continue
        }

		err = storage.Store_posts_as_json(entry, posts)
		if err != nil {
			fmt.Print("Error: ", err)
			return
		}

        err = storage.Create_html_page(entry, posts)
		if err != nil {
			fmt.Print("Error: ", err)
			return
		}

		thread_file_list, err := models.Get_media_urls_from_posts(entry, posts)
		if err != nil {
			fmt.Print("Error: ", err)
			return
		}

		file_list = append(file_list, thread_file_list...)
	}

	fmt.Println("Downloading media content...")
	networking.Download_media(file_list)
}
