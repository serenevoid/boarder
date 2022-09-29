package networking

import (
	"boarder/models"
	"boarder/util"
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	progressbar "github.com/schollz/progressbar/v3"
)

func Get_board_from_user() string {
	var i string
	fmt.Print("Enter Board ID: ")
	fmt.Scan(&i)
	return i
}

func Get_thread_from_user() int {
	var i int
	fmt.Print("Enter thread ID: ")
	fmt.Scan(&i)
	return i
}

func get_response(url string) []byte {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{},
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	util.CheckErr(err)

	req.Header.Set("User-Agent", "linux:go-postgrabber:v0.1")

	resp, err := client.Do(req)
	util.CheckErr(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	util.CheckErr(err)

	return body
}

func Get_threads_from_board(board string) []int {
	url := fmt.Sprintf("https://a.4cdn.org/%s/threads.json", board)
	body := get_response(url)
	threads := models.Get_threads_from_json(body)
	return threads
}

func Get_posts_from_thread(entry string) []models.Post {
	entry_elements := strings.Split(entry, "_")
	board := entry_elements[0]
	thread := entry_elements[1]
	url := fmt.Sprintf("https://a.4cdn.org/%s/thread/%s.json", board, fmt.Sprint(thread))
	body := get_response(url)
	posts := models.Get_posts_from_json(body)
	return posts
}

func downloadFile(URL string) ([]byte, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}
	var data bytes.Buffer
	_, err = io.Copy(&data, response.Body)
	if err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}

func Download_media(file_list []models.File) {
	var waiter sync.WaitGroup
	waiter.Add(len(file_list))
	bar := progressbar.NewOptions(len(file_list),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("Downloading media of posts..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	for _, file := range file_list {
		go func(file models.File, waiter *sync.WaitGroup) {
			_, err := os.Stat(file.File_name)
			if errors.Is(err, os.ErrNotExist) {
				b, err := downloadFile(file.URL)
				if err != nil {
					util.CheckErr(err)
				}
				ioutil.WriteFile(file.File_name, b, 0777)
			}
			waiter.Done()
			bar.Add(1)
		}(file, &waiter)
	}
	waiter.Wait()
}
