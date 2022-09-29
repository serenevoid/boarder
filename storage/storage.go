package storage

import (
	"boarder/models"
	"boarder/util"
	"os"
	"strings"
)

func Store_posts_in_md(entry string, posts []models.Post) {
    content := models.Format_posts_to_string(posts)
    entry_elements := strings.Split(entry, "_")
    board := entry_elements[0]
    thread := entry_elements[1]
    file_name := "archive" + string(os.PathSeparator) + board + string(os.PathSeparator) + thread + string(os.PathSeparator) + "data.md"
    f, err := os.Create(file_name)
    util.CheckErr(err)
    defer f.Close()
    f.WriteString(content)
}

func Store_post_images(posts []models.Post) {

}
