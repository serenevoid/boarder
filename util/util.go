package util

import (
	"bufio"
	"os"
	"strings"
)

/*
 Check error and panic if not nil.

 @param {error} err - error variable

 @example checkErr(err)
*/
func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func Load_config() []string {
	f, err := os.Open("threads.txt")
	CheckErr(err)
	scanner := bufio.NewScanner(f)

    var entry_list []string
	for scanner.Scan() {
		// do something with a line
		entry := scanner.Text()
		if len(entry) > 0 {
			if !strings.HasPrefix(entry, "//") {
                entry_list = append(entry_list, entry)
				entry_elements := strings.Split(entry, "_")
				board := entry_elements[0]
				thread := entry_elements[1]
				create_folder_structure(board, thread)
			}
		}
	}
	CheckErr(scanner.Err())
    return entry_list
}

func create_folder_structure(board string, thread string) {
    sep := string(os.PathSeparator)
	os.MkdirAll("archive"+sep+board+sep+thread+sep+"media", 0777)
}
