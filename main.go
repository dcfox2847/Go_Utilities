package main

import "fmt"

// Store variables used for functions outside of 'main' here
var srcFilePath string = "C:\\Users\\dcfox\\Desktop\\Code\\go_prac\\src_test\\move.txt"
var dstFilePath string = "C:\\Users\\dcfox\\Desktop\\Code\\go_prac\\dst_test\\moved.txt"

// Function created for individual services begin here
// Will be tested here and then moved into separate file for increased functionality and readability

// Pull logs all logs from the system recursively from given directory, then store them in a

// Begin the execution of the 'main' function
func main() {

	var err error

	// Attempt to run the single file puller function from 'filePuller.go'
	if err = filePuller(srcFilePath, dstFilePath); err != nil {
		fmt.Println(err)
	}

	// Attempt to run the directory copy / puller function from 'dirPuller.go'
	
}
