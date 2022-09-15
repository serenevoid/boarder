package main

import "fmt"

func main() {
	board := get_user_input()
	threads := get_threads(board)
	for i := 0; i < len(threads); i++ {
		posts := get_posts(board, threads[i])
		for j := 0; j < len(posts); j++ {
			fmt.Printf("https://i.4cdn.org/%s/%s\n", board, posts[j])
		}
		// downloadMultipleFiles(image_list)
	}
}
