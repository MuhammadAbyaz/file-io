package main

import (
	"files/utils"
	"fmt"
)

func main() {
	var path string
	fmt.Print("Enter file path: ")
	fmt.Scan(&path)
	content := utils.Reader(path)
	fmt.Println(content)
	var writePath, contentToWrite string
	fmt.Print("Enter path of file to write on: ")
	fmt.Scan(&writePath)
	fmt.Print("Enter content: ")
	fmt.Scan(&contentToWrite)
	utils.Writer(writePath, contentToWrite)
}
