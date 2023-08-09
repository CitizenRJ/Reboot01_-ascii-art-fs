package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	if !isValid(args[0]) {
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
	printBanners(argsArr, arr)
}

// Check for valid of characters by runes from 32 to 126
func isValid(s string) bool {
	for _, ch := range s {
		if ch < ' ' && ch != 10 || ch > '~' {
			return false
		}
	}
	return true
}

func OnlyContains(s, sep string) bool {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "\\" {
			if string(s[i+1]) == "n" {
				if i != len(s)-3 {
					i++
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// Print the full outcome
func printBanners(banners, arr []string) {
	num := 0
	for _, ch := range banners {
		num = num + 1
		if ch == "" {
			if num < len(banners) {
			fmt.Println()
			continue
			} else {
				continue
			}
		}
		for i := 0; i < 8; i++ {
			for _, j := range ch {
				n := (j-32)*9 + 1
				fmt.Print(arr[int(n)+i])

			}
				fmt.Println()

		}
		fmt.Println()

	}
}
