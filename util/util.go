package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
   Load the config file threads.txt from directoy of running the program

   @return ([]string, error) - (array of threads, error)
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
                if len(entry_elements) != 2 {
                    return nil, fmt.Errorf("entry not in proper format: %s", entry)
                }
				board := entry_elements[0]
				thread := entry_elements[1]
				err := create_folder_structure(board, thread)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to read threads.txt")
	}
	return entry_list, nil
}

/*
   Create all the folders and subfolders required for storing data from threads

   @return error
*/
func create_folder_structure(board string, thread string) error {
	sep := string(os.PathSeparator)
	err := os.MkdirAll("archive"+sep+board+sep+thread+sep+"media", 0777)
	if err != nil {
		return fmt.Errorf("unable to create folder structure")
	}
	return nil
}
