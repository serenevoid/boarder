package main

import (
	"encoding/json"
	"fmt"
)

type thread struct {
	No            int `json:"no"`
	Last_modified int `json:"last_modified"`
	Replies       int `json:"replies"`
}

type page struct {
	Page    uint8    `json:"page"`
	Threads []thread `json:"threads"`
}

type post struct {
	No           int    `json:"no"`
	Now          string `json:"now"`
	Name         string `json:"name"`
	Sub          string `json:"sub"`
	Com          string `json:"Com"`
	Filename     string `json:"filename"`
	Ext          string `json:"ext"`
	W            int    `json:"w"`
	H            int    `json:"h"`
	Tn_w         int    `json:"tn_w"`
	Tn_h         int    `json:"tn_h"`
	Tim          int64  `json:"tim"`
	Time         int64  `json:"time"`
	Md5          string `json:"md5"`
	Fsize        string `json:"fsize"`
	Resto        int    `json:"resto"`
	Bumplimit    int    `json:"bumplimit"`
	Imagelimit   int    `json:"imagelimit"`
	Semantic_url string `json:"semantic_url"`
	Replies      int    `json:"replies"`
	Images       int    `json:"images"`
	Unique_ips   int    `json:"unique_ips"`
}

func list_threads(body []byte) []int {
	var dat []page
	var threads_list []int

	err := json.Unmarshal(body, &dat)
	checkErr(err)

	for i := 0; i < len(dat); i++ {
		page := dat[i]
		threads := page.Threads
		for j := 0; j < len(threads); j++ {
			thread_data := threads[j]
			thread_id := thread_data.No
			threads_list = append(threads_list, thread_id)
		}
	}

	return threads_list
}

func list_posts(body []byte) []string {
	var post_list []string
	var dat map[string]interface{}

	err := json.Unmarshal(body, &dat)
	checkErr(err)
	return post_list
}
