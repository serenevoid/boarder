package models

import (
	"boarder/common"
	"encoding/json"
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

type Post struct {
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
	Fsize        int    `json:"fsize"`
	Resto        int    `json:"resto"`
	Bumplimit    int    `json:"bumplimit"`
	Imagelimit   int    `json:"imagelimit"`
	Semantic_url string `json:"semantic_url"`
	Replies      int    `json:"replies"`
	Images       int    `json:"images"`
	Unique_ips   int    `json:"unique_ips"`
}

type post_array struct {
	Posts []Post `json:"posts"`
}

func Get_threads_from_json(body []byte) []int {
	var dat []page
	var threads_list []int

	err := json.Unmarshal(body, &dat)
	common.CheckErr(err)

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

func Get_posts_from_json(body []byte) []Post {
	var media_list []Post
	var dat post_array

	err := json.Unmarshal(body, &dat)
	common.CheckErr(err)

	post_list := dat.Posts
	for i := 0; i < len(post_list); i++ {
		post_content := post_list[i]
		if post_content.Tim != 0 {
			media_list = append(media_list, post_content)
		}
	}
	return media_list
}
