package main

import (
	"boarder/models"
	"boarder/networking"
	"boarder/storage"
	"boarder/util"
	"fmt"
)

func main() {
    fmt.Println("BOARDER V0.1.0")
	entry_list := util.Load_config()
    var file_list []models.File
    fmt.Println("Collecting thread data")
	for _, entry := range entry_list {
		posts := networking.Get_posts_from_thread(entry)
		storage.Store_posts_in_md(entry, posts)
        thread_file_list := models.Get_media_urls_from_posts(entry, posts)
        file_list = append(file_list, thread_file_list...)
	}
    networking.Download_media(file_list)
}
