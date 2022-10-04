package storage

import (
	"boarder/models"
	"fmt"
	"html/template"
	"os"
	"strings"
)

/*
   Store data of posts in a md file on the correct folder

   @param (string, []Posts) - (details of board and thread, array of posts in thread)

   @return error - error from creation of file
*/
func Store_posts_in_md(entry string, posts []models.Post) error {
    content := models.Format_posts_to_string(posts)
    entry_elements := strings.Split(entry, "_")
    board := entry_elements[0]
    thread := entry_elements[1]
    file_name := "archive" + string(os.PathSeparator) + board + string(os.PathSeparator) + thread + string(os.PathSeparator) + "data.md"
    f, err := os.Create(file_name)
    if err != nil {
        return err
    }
    defer f.Close()
    f.WriteString(content)
    Create_html_page(entry, posts)
    
    return nil
}

func Create_html_page(entry string, posts []models.Post) error {
    t := template.Must(template.ParseFiles("./storage/template.gohtml"))
    entry_elements := strings.Split(entry, "_")
    board := entry_elements[0]
    thread := entry_elements[1]

    data := struct {
        Title string
        Posts []models.Post
    }{
        Title: entry,
        Posts: posts,
    }

    file_name := "archive" + string(os.PathSeparator) + board + string(os.PathSeparator) + thread + string(os.PathSeparator) + "index.html"
    f, err := os.Create(file_name)
    if err != nil {
        return err
    }
    defer f.Close()

    err = t.Execute(f, data)
    if err != nil {
        return err
    }

    return nil
}
