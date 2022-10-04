package models

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
   Thread struct used to store the thread details
*/
type thread struct {
	No            int `json:"no"`
	Last_modified int `json:"last_modified"`
	Replies       int `json:"replies"`
}

/*
   Page struct used to store page details
*/
type page struct {
	Page    uint8    `json:"page"`
	Threads []thread `json:"threads"`
}

/*
   Post struct used to store post details
*/
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
	Size         string
	Resto        int    `json:"resto"`
	Bumplimit    int    `json:"bumplimit"`
	Imagelimit   int    `json:"imagelimit"`
	Semantic_url string `json:"semantic_url"`
	Replies      int    `json:"replies"`
	Images       int    `json:"images"`
	Unique_ips   int    `json:"unique_ips"`
}

/*
   Post array type to store array of posts
*/
type post_array struct {
	Posts []Post `json:"posts"`
}

/*
   File struct to store file data
*/
type File struct {
	File_name     string
	Media_URL     string
}

/*
   Extract the thread IDs from the json byte array.

   @param []byte - the body of the response from http req

   @return ([]int, error) - (array of thread IDs, error)
*/
func Get_threads_from_json(body []byte) ([]int, error) {
	var dat []page
	var threads_list []int

	err := json.Unmarshal(body, &dat)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(dat); i++ {
		page := dat[i]
		threads := page.Threads
		for j := 0; j < len(threads); j++ {
			thread_data := threads[j]
			thread_id := thread_data.No
			threads_list = append(threads_list, thread_id)
		}
	}

	return threads_list, nil
}

/*
   Extract the thread IDs from the json byte array.

   @param []byte - the body of the response from http req

   @return ([]int, error) - (array of thread IDs, error)
*/
func Get_posts_from_json(body []byte) ([]Post, error) {
	var media_list []Post
	var dat post_array

	err := json.Unmarshal(body, &dat)
	if err != nil {
		return nil, err
	}

	post_list := dat.Posts
	for i := 0; i < len(post_list); i++ {
		post_content := post_list[i]
		post_content.Size = ByteCountDecimal(post_content.Fsize)
		if post_content.Tim != 0 {
			media_list = append(media_list, post_content)
		}
	}
	return media_list, nil
}

/*
   Convert post details to a string form to store in md file.

   @param []Post - array of all posts present in a thread

   @return string - all posts combined to a single string
*/
func Format_posts_to_string(posts []Post) string {
	var content string
	for num, post := range posts {
		content = fmt.Sprintf("%s%s", content, "\n\nPost "+fmt.Sprint(num)+"\n-")
		if post.No != 0 {
			content = fmt.Sprintf("%s%s", content, "\nNo: "+fmt.Sprint(post.No))
		}
		if post.Now != "" {
			content = fmt.Sprintf("%s%s", content, "\\\nNow: "+post.Now)
		}
		if post.Name != "" {
			content = fmt.Sprintf("%s%s", content, "\\\nName: "+post.Name)
		}
		if post.Sub != "" {
			content = fmt.Sprintf("%s%s", content, "\\\nSub: "+post.Sub)
		}
		if post.Com != "" {
			content = fmt.Sprintf("%s%s", content, "\\\nCom: "+post.Com)
		}
		if post.Filename != "" {
			content = fmt.Sprintf("%s%s", content, "\\\nFilename: "+post.Filename)
		}
		if post.Ext != "" {
			content = fmt.Sprintf("%s%s", content, "\\\nExt: "+post.Ext)
		}
		if post.W != 0 {
			content = fmt.Sprintf("%s%s", content, "\\\nW: "+fmt.Sprint(post.W))
		}
		if post.H != 0 {
			content = fmt.Sprintf("%s%s", content, "\\\nH: "+fmt.Sprint(post.H))
		}
		if post.Tn_w != 0 {
			content = fmt.Sprintf("%s%s", content, "\\\nTn_w: "+fmt.Sprint(post.Tn_w))
		}
		if post.Tn_h != 0 {
			content = fmt.Sprintf("%s%s", content, "\\\nTn_h: "+fmt.Sprint(post.Tn_h))
		}
		if post.Tim != 0 {
			content = fmt.Sprintf("%s%s", content, "\\\nTim: "+fmt.Sprint(post.Tim))
		}
		if post.Time != 0 {
			content = fmt.Sprintf("%s%s", content, "\\\nTime: "+fmt.Sprint(post.Time))
		}
		if post.Md5 != "" {
			content = fmt.Sprintf("%s%s", content, "\\\nMd5: "+post.Md5)
		}
		if post.Fsize != 0 {
			content = fmt.Sprintf("%s%s", content, "\\\nFsize: "+fmt.Sprint(post.Fsize))
		}
		if post.Resto != 0 {
			content = fmt.Sprintf("%s%s", content, "\\\nResto: "+fmt.Sprint(post.Resto))
		}
		if post.Bumplimit != 0 {
			content = fmt.Sprintf("%s%s", content, "\\\nBumplimit: "+fmt.Sprint(post.Bumplimit))
		}
		if post.Imagelimit != 0 {
			content = fmt.Sprintf("%s%s", content, "\\\nImagelimit: "+fmt.Sprint(post.Imagelimit))
		}
		if post.Semantic_url != "" {
			content = fmt.Sprintf("%s%s", content, "\\\nSemantic_url: "+post.Semantic_url)
		}
		if post.Replies != 0 {
			content = fmt.Sprintf("%s%s", content, "\\\nReplies: "+fmt.Sprint(post.Replies))
		}
		if post.Images != 0 {
			content = fmt.Sprintf("%s%s", content, "\\\nImages: "+fmt.Sprint(post.Images))
		}
		if post.Unique_ips != 0 {
			content = fmt.Sprintf("%s%s", content, "\\\nUnique_ips: "+fmt.Sprint(post.Unique_ips))
		}
	}
	return content
}

/*
   Get all media URLs from the list of posts for downloading them.

   @param (string, []Post) - (board and thread ID, array of all posts)

   @return ([]File, error) - (list of all files, error)
*/
func Get_media_urls_from_posts(entry string, posts []Post) ([]File, error) {
	var file_list []File
	entry_elements := strings.Split(entry, "_")
	board := entry_elements[0]
	thread := entry_elements[1]

	for _, post := range posts {
		var media_file File
		if post.Filename != "" {
			sep := string(os.PathSeparator)
			media_file.File_name = "archive" + sep + board + sep + thread + sep + "media" + sep + post.Filename + post.Ext
		}
		if post.Tim != 0 {
			media_file.Media_URL = fmt.Sprintf("https://i.4cdn.org/%s/%s%s", board, fmt.Sprint(post.Tim), post.Ext)
		}
		var thumbnail_file File
		if post.Filename != "" {
			sep := string(os.PathSeparator)
			thumbnail_file.File_name = "archive" + sep + board + sep + thread + sep + "thumbnails" + sep + post.Filename + "s.jpg"
		}
		if post.Tim != 0 {
			thumbnail_file.Media_URL = fmt.Sprintf("https://i.4cdn.org/%s/%ss.jpg", board, fmt.Sprint(post.Tim))
		}
		file_list = append(file_list, media_file, thumbnail_file)
	}

	return file_list, nil
}

/*
   Convert size from bytes to a readable format

   @param (int) - (size byte int)

   @return (string) - (readable string format)
*/
func ByteCountDecimal(b int) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
