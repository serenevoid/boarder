package main

import "io/ioutil"

func checkErr(err error) {
    if err != nil {
        panic(err.Error())
    }
}

func createFiles(byte_array [][]byte, name_array []string) {
    for i := 0; i < len(name_array); i++ {
        err := ioutil.WriteFile(name_array[i], byte_array[i], 0644)
        checkErr(err)
    }
}
