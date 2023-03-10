package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) >= 2 {
		args = args[1:]
	} else {
		args = []string{"help"}
	}
	switch args[0] {
	case "compress":
		if len(args) >= 3 {
			os.WriteFile(args[2], []byte(compress(listDirTree(args[1]))), os.FileMode(0666))
		}
	case "decompress":
		if len(args) >= 2 {
			file, err := os.ReadFile(args[1])
			if err == nil {
				decompress(loadStream(string(file)))
			}
		}
	default:
		fmt.Printf("HELP:\ncompress [directory to compress] [output filename]\ndecompress [compressed file]\n")
	}
}
