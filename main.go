package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	parse()
}

func generateFilename() string {
	t := time.Now()
	return fmt.Sprintf("%d%d%d-%d%d%d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func createSnippet(s string) {
	titleEnd := strings.Index(s, ":")
	if titleEnd == -1 {
		abort("Title required")
	}
	title := s[:titleEnd]

	links := make([]string, 0)
	for _, sub := range strings.Split(s, "[[") {
		end := strings.Index(sub, "]]")
		if end != -1 {
			links = append(links, sub[:end])
		}
	}


	fmt.Println("parsed title", title)
	fmt.Println("parsed links", links)

	err := os.WriteFile(title, []byte(s[titleEnd+1:]), 0777)
	if err != nil {
		fmt.Println("Error creating snippet", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func help() {
	fmt.Println("help")
	os.Exit(0)
}

func abort(s string) {
	fmt.Println(s)
	os.Exit(1)
}

func new() {
	args := os.Args[2:]

	if len(args) == 0 {
		fmt.Println("args required for new snippet")
		os.Exit(0)
	}

	s := strings.Join(args, " ")
	createSnippet(s)
}

func search() {
	fmt.Println("search")
}

func parse() {
	if len(os.Args) < 2 {
		help()
	}

	arg := strings.ReplaceAll(os.Args[1], "-", "")
	if arg == "n" || arg == "new" {
		new()
	} else if arg == "s" || arg == "search" {
		search()
	} else {
		help()
	}
}
