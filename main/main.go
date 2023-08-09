package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"asciiArtFs"
)

const (
	fileLen = 855
)

// check amount of arguments
func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Not a valid amount of arguments")
		return
	}
	args := os.Args[1:]
	// bol := asciiArtFs.IsValid(args[0])
	if !(asciiArtFs.IsValid(args[0])) {
		fmt.Println("Not a valid character")
		return
	}

	text := args[0]    // "hello" == [0]
	font := "standard" //base font
	if len(args) == 2 {
		switch args[1] {
		case "shadow":
			font = "shadow"
		case "thinkertoy":
			font = "thinkertoy"
		case "standard":
			font = "standard"
		default:
			fmt.Println("Not a valid font")
			return
		}
	}

	// Read the content of the file
	argsArr := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")
	arr := []string{}
	readFile, err := os.Open("fonts/" + font + ".txt")
	defer readFile.Close()

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}

	if len(arr) != fileLen {
		fmt.Println("File is corrupted")
		return
	}
	larg := len(argsArr)
	if larg >= 2 {
		if argsArr[larg-1] == "" && argsArr[larg-2] != "" {
			argsArr = argsArr[:larg-1]
		}
	}
	asciiArtFs.PrintBanners(argsArr, arr)
}
