package networking

import (
	"boarder/models"
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
)

/*
    Get all response data from an http req.

    @param string - the URL to request data from

    @return ([]byte, error) - (response body data, error)
*/
func get_response(url string) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{},
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "linux:go-postgrabber:v0.1")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, nil
    }

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

/*
    Get all posts from a thread

    @param string - board and thread ID

    @return ([]Posts, error) - (list of posts present in thread, error)
*/
func Get_posts_from_thread(entry string) ([]models.Post, error) {
	entry_elements := strings.Split(entry, "_")
	board := entry_elements[0]
	thread := entry_elements[1]
	url := fmt.Sprintf("https://a.4cdn.org/%s/thread/%s.json", board, fmt.Sprint(thread))
	body, err := get_response(url)
	if err != nil {
		return nil, fmt.Errorf("%s %s", err, entry)
	}
    if body == nil {
        fmt.Println("Unable to reach thread ", entry)
        return nil, nil
    }
	posts, err := models.Get_posts_from_json(body)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

/*
    Download specified file from URL.

    @param string - URL of the file

    @return ([]byte, error) - bytes recieved from the request
*/
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

/*
    Download a files from a given list

    @param []File - list of all files to be downloaded
*/
func Download_media(file_list []models.File) {
	var waiter sync.WaitGroup
	waiter.Add(len(file_list))
	for _, file := range file_list {
		go func(file models.File, waiter *sync.WaitGroup) {
			_, err := os.Stat(file.File_name)
			if errors.Is(err, os.ErrNotExist) {
				b, err := downloadFile(file.Media_URL)
				if err != nil {
					fmt.Println("Download of " + file.File_name + " failed")
				}
				ioutil.WriteFile(file.File_name, b, 0777)
			}
			waiter.Done()
		}(file, &waiter)
	}
	waiter.Wait()
}
