package utils

import "os"

func Reader(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("Can't read the file")
	} else {
		return string(data)
	}
}
