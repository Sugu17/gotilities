package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Check if cmd args are provided
	if len(os.Args) > 1 {
		for _, filePath := range os.Args[1:] {
			file, err := os.Open(filePath)
			if err != nil {
				log.Println("Invalid file path !")
				continue
			}
			if _, err := io.Copy(os.Stdout, file); err != nil {
				log.Println("Error in copying file !", err)
			}
			file.Close()
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}
