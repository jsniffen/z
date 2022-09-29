package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const folder = "z"

func main() {
	parse()
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

	title := strings.Join(args, " ")
	data, err := edit(title)
	if err != nil {
		abort("Error editing file")
	}

	fmt.Println(string(data))
}

func search() {
	fmt.Println("search")
}

func edit(title string) (string, error) {
	cmd := exec.Command("vim", title)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	b, err := os.ReadFile(title)
	if err != nil {
		return "", err
	}

	return string(b), nil
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
