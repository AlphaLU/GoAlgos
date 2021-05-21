package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				fmt.Println("[+]", path, info.Size())
				return nil
			}
			fmt.Println("\t", path, info.Size())

			return nil
		})
	if err != nil {
		log.Println(err)
	}

}
