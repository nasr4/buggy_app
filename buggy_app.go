package main

import "os"

func main() {
	for true {
		file, err := os.Open("./config_file")
		if err != nil {
			continue
		} else {
			file.Close()
			break
		}
	}
}
