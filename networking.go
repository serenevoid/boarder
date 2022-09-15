package main

import (
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

func get_user_input() string {
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
	checkErr(err)

	req.Header.Set("User-Agent", "linux:go-postgrabber:v0.1")

	resp, err := client.Do(req)
	checkErr(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

    return body
}

func get_threads(board string) []int {
	url := fmt.Sprintf("https://a.4cdn.org/%s/threads.json", board)
    body := get_response(url)
	threads := list_threads(body)
	return threads
}

func get_posts(board string, thread int) []string {
	url := fmt.Sprintf("https://a.4cdn.org/%s/thread/%s.json", board, fmt.Sprint(thread))
    body := get_response(url)
	posts := list_posts(body)
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

func downloadMultipleFiles(urls []string) {
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
