package networking

import (
	"boarder/util"
	"boarder/models"
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func Get_user_input() string {
	var i string
	fmt.Print("Enter Board ID: ")
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

func Get_posts_from_threads(board string, thread int) []models.Post {
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

func Download_media(urls []string) {
	var waiter sync.WaitGroup
	waiter.Add(len(urls))
	for _, URL := range urls {
		go func(URL string, waiter *sync.WaitGroup) {
			b, err := downloadFile(URL)
			url_comp := strings.Split(URL, "/")
			file_name := url_comp[len(url_comp)-1]
			fmt.Println("Started - " + file_name)
			if err != nil {
				return
			}
			ioutil.WriteFile("r_"+file_name, b, 0777)
			waiter.Done()
			fmt.Println("Downloaded - " + file_name)
		}(URL, &waiter)
	}
	waiter.Wait()
}
