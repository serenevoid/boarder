package storage

import (
	"boarder/models"
	"encoding/json"
	"html/template"
	"os"
	"strings"
)

/*
   Store data of posts in a md file on the correct folder

   @param (string, []Posts) - (details of board and thread, array of posts in thread)

   @return error - error from creation of file
*/
func Store_posts_as_json(entry string, posts []models.Post) error {
    content, err := json.Marshal(posts)
    if err != nil {
        return err
    }
    entry_elements := strings.Split(entry, "_")
    board := entry_elements[0]
    thread := entry_elements[1]
    sep := string(os.PathSeparator)
    file_name := "archive" + sep + board + sep + thread + sep + entry + ".json"
    f, err := os.Create(file_name)
    if err != nil {
        return err
    }
    defer f.Close()
    f.WriteString(string(content))
    
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

    sep := string(os.PathSeparator)
    file_name := "archive" + sep + board + sep + thread + sep + "index.html"
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
