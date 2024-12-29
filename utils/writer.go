package utils

import "os"

func Writer(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0666)
	if err != nil {
		panic("Unable to write to file")
	}
}
