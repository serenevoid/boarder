package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
 Check error and panic if not nil.

 @param {error} err - error variable
*/
func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

/*
    Load the config file threads.txt from directoy of running the program

    @return []string - array of threads
*/
func Load_config() ([]string, error) {
	f, err := os.Open("threads.txt")
    if err != nil {
        return nil, fmt.Errorf("file threads.txt does not exist")
    }
    defer f.Close()
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
    if err = scanner.Err(); err != nil {
        return nil, fmt.Errorf("unable to read threads.txt %s", err)
    }
    return entry_list, nil
}

func create_folder_structure(board string, thread string) {
    sep := string(os.PathSeparator)
    err := os.MkdirAll("archive"+sep+board+sep+thread+sep+"media", 0777)
    CheckErr(err)
}
