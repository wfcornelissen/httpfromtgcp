package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()
	content := make([]byte, 1000)
	length, err := file.Read(content)
	if err != nil {
		fmt.Printf("Error getting file info: %v", err)
	}

	section := make([]byte, 8)
	for i := 0; i < length; i += 8 {
		if i+8 > length {
			remaining := length - i
			lastSection := make([]byte, remaining)
			_, err = file.ReadAt(lastSection, int64(i))
			fmt.Printf("read: %s\n", string(lastSection))
			os.Exit(0)
		} else {
			_, err = file.ReadAt(section, int64(i))
			if err != nil {
				fmt.Printf("Error reading file: %v", err)
			}
		}
		fmt.Printf("read: %s\n", string(section))
	}

}
