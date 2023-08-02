package main

import (
	"asciiArt"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "regexp"
)

func main() {
	Str := os.Args[1:]
	if len(Str) > 2 || len(Str) < 1 {
		if len(Str) == 0 {
			fmt.Println("Error: no arguments :).")
			os.Exit(0)
		} else if len(Str) > 2 {
			fmt.Println("Error: so many Args.")
			os.Exit(1)
		}
	} else {
		var fileLines []string
		input := os.Args[1]
		if len(Str) == 1 {
			if len(Str) == 1 {
				file, err := os.Open("standard.txt")
				if err != nil {
					log.Fatal(err)
				}
				fileScanner := bufio.NewScanner(file)
				fileScanner.Split(bufio.ScanLines)
				for fileScanner.Scan() {
					fileLines = append(fileLines, fileScanner.Text())
				}
				file.Close()
			}
		} else {
			fileType := os.Args[2]
			shadow := "shadow"
			standard := "standard"
			thinkertoy := "thinkertoy"
			if fileType == standard {
				file, err := os.Open("standard.txt")
				if err != nil {
					log.Fatal(err)
				}
				fileScanner := bufio.NewScanner(file)
				fileScanner.Split(bufio.ScanLines)
				for fileScanner.Scan() {
					fileLines = append(fileLines, fileScanner.Text())
				}
				file.Close()
			} else if fileType == shadow {
				file, err := os.Open("shadow.txt")
				if err != nil {
					log.Fatal(err)
				}
				fileScanner := bufio.NewScanner(file)
				fileScanner.Split(bufio.ScanLines)
				for fileScanner.Scan() {
					fileLines = append(fileLines, fileScanner.Text())
				}
				file.Close()
			} else if fileType == thinkertoy {
				file, err := os.Open("thinkertoy.txt")
				if err != nil {
					log.Fatal(err)
				}
				fileScanner := bufio.NewScanner(file)
				fileScanner.Split(bufio.ScanLines)
				for fileScanner.Scan() {
					fileLines = append(fileLines, fileScanner.Text())
				}
				file.Close()
			} else {
				fmt.Println("Error: too many Args lol.")
				os.Exit(2)
			}
		}
		var fileload []int
		result := strings.ReplaceAll(input, "\\t", "   ")
		result = strings.ReplaceAll(result, "\\n", " \\n ")
		words := strings.Split(result, " \\n ")
		for j := 0; j < len(words); j++ {
			if words[j] != "" {
				fileload = asciiArt.ConvertToASCIIArt(words[j])
				fileload1 := asciiArt.PrintLines(fileload, fileLines)
				for i := 0; i < len(fileload1); i++ {
					if strings.Join(fileload1[i], "") != "" {
						fmt.Println(strings.Join(fileload1[i], ""))
					}
				}
			}  else if j != 1 && words[j] == "" {
				fmt.Println()
			}
		}
	}
}
