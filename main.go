package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]

	for _, fPath := range args {
		dir, filename := filepath.Split(fPath)

		if dir != "" {
			err := os.MkdirAll(dir, os.ModePerm)

			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		if filename != "" {
			file, err := os.Create(fPath)
			defer file.Close()

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
