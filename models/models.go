package models

import (
	"boarder/util"
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

type File struct {
    File_name string
    URL string
}

func Get_threads_from_json(body []byte) []int {
	var dat []page
	var threads_list []int

	err := json.Unmarshal(body, &dat)
	util.CheckErr(err)

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
	util.CheckErr(err)

	post_list := dat.Posts
	for i := 0; i < len(post_list); i++ {
		post_content := post_list[i]
		if post_content.Tim != 0 {
			media_list = append(media_list, post_content)
		}
	}
	return media_list
}

func Format_posts_to_string(posts []Post) string {
    var content string
    for num, post := range(posts) {
        content = fmt.Sprintf("%s%s", content, "\n\nPost " + fmt.Sprint(num) + "\n-")
        if post.No != 0 {
            content = fmt.Sprintf("%s%s", content, "\nNo: " + fmt.Sprint(post.No))
        }
        if post.Now != "" {
            content = fmt.Sprintf("%s%s", content, "\\\nNow: " + post.Now)
        }
        if post.Name != "" {
            content = fmt.Sprintf("%s%s", content, "\\\nName: " + post.Name)
        }
        if post.Sub != "" {
            content = fmt.Sprintf("%s%s", content, "\\\nSub: " + post.Sub)
        }
        if post.Com != "" {
            content = fmt.Sprintf("%s%s", content, "\\\nCom: " + post.Com)
        }
        if post.Filename != "" {
            content = fmt.Sprintf("%s%s", content, "\\\nFilename: " + post.Filename)
        }
        if post.Ext != "" {
            content = fmt.Sprintf("%s%s", content, "\\\nExt: " + post.Ext)
        }
        if post.W != 0 {
            content = fmt.Sprintf("%s%s", content, "\\\nW: " + fmt.Sprint(post.W))
        }
        if post.H != 0 {
            content = fmt.Sprintf("%s%s", content, "\\\nH: " + fmt.Sprint(post.H))
        }
        if post.Tn_w != 0 {
            content = fmt.Sprintf("%s%s", content, "\\\nTn_w: " + fmt.Sprint(post.Tn_w))
        }
        if post.Tn_h != 0 {
            content = fmt.Sprintf("%s%s", content, "\\\nTn_h: " + fmt.Sprint(post.Tn_h))
        }
        if post.Tim != 0 {
            content = fmt.Sprintf("%s%s", content, "\\\nTim: " + fmt.Sprint(post.Tim))
        }
        if post.Time != 0 {
            content = fmt.Sprintf("%s%s", content, "\\\nTime: " + fmt.Sprint(post.Time))
        }
        if post.Md5 != "" {
            content = fmt.Sprintf("%s%s", content, "\\\nMd5: " + post.Md5)
        }
        if post.Fsize != 0 {
            content = fmt.Sprintf("%s%s", content, "\\\nFsize: " + fmt.Sprint(post.Fsize))
        }
        if post.Resto != 0 {
            content = fmt.Sprintf("%s%s", content, "\\\nResto: " + fmt.Sprint(post.Resto))
        }
        if post.Bumplimit != 0 {
            content = fmt.Sprintf("%s%s", content, "\\\nBumplimit: " + fmt.Sprint(post.Bumplimit))
        }
        if post.Imagelimit != 0 {
            content = fmt.Sprintf("%s%s", content, "\\\nImagelimit: " + fmt.Sprint(post.Imagelimit))
        }
        if post.Semantic_url != "" {
            content = fmt.Sprintf("%s%s", content, "\\\nSemantic_url: " + post.Semantic_url)
        }
        if post.Replies != 0 {
            content = fmt.Sprintf("%s%s", content, "\\\nReplies: " + fmt.Sprint(post.Replies))
        }
        if post.Images != 0 {
            content = fmt.Sprintf("%s%s", content, "\\\nImages: " + fmt.Sprint(post.Images))
        }
        if post.Unique_ips != 0 {
            content = fmt.Sprintf("%s%s", content, "\\\nUnique_ips: " + fmt.Sprint(post.Unique_ips))
        }
    }
    return content
}

func Get_media_urls_from_posts(entry string, posts []Post) []File {
    var file_list []File
    entry_elements := strings.Split(entry, "_")
    board := entry_elements[0]
    thread := entry_elements[1]
    for _, post := range(posts) {
        var new_file File
        if post.Filename != "" {
            sep := string(os.PathSeparator)
            new_file.File_name = "archive" + sep + board + sep + thread + sep + "media" + sep + post.Filename + post.Ext
        }
        if post.Tim != 0 {
            new_file.URL = fmt.Sprintf("https://i.4cdn.org/%s/%s%s", board, fmt.Sprint(post.Tim), post.Ext)
        }
        file_list = append(file_list, new_file)
    }
    return file_list
}
