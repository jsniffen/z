package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	LinkFile = ".links"
	TagFile  = ".tags"
)

func main() {
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

func help() {
	fmt.Println("help")
	os.Exit(0)
}

func abort(s string) {
	fmt.Println(s)
	os.Exit(1)
}

func writeTags(title string, tags []string) {
	f, err := os.OpenFile(TagFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var b bytes.Buffer
	for _, tag := range tags {
		b.WriteString(fmt.Sprintf("%s=%s\n", title, tag))
	}

	_, err = f.Write(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}

func writeLinks(title string, links []string) {
}

func createFileIfNotExist(fn string) error {
	_, err := os.Stat(fn)
	if errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(fn)
		if err != nil {
			return err
		}
		f.Close()
	}
	return nil
}

func parseFileContent(fc string) ([]string, []string) {
	tags := make([]string, 0)
	links := make([]string, 0)

	for _, sub := range strings.Split(fc, "[[") {
		end := strings.Index(sub, "]]")
		if end != -1 {
			links = append(links, sub[:end])
		}
	}

	for _, word := range strings.Fields(fc) {
		if strings.HasPrefix(word, "#") {
			tags = append(tags, word[1:])
		}
	}

	return tags, links
}

func new() {
	args := os.Args[2:]

	if len(args) == 0 {
		fmt.Println("args required for new snippet")
		os.Exit(0)
	}

	title := strings.Join(args, " ")
	err := createFileIfNotExist(title)
	if err != nil {
		abort("error creating file")
	}

	fc, err := edit(title)
	if err != nil {
		abort("Error editing file")
	}

	tags, links := parseFileContent(fc)

	for _, link := range links {
		createFileIfNotExist(link)
	}

	writeLinks(title, links)
	writeTags(title, tags)

	fmt.Println("parsed tags:", strings.Join(tags, ", "))
	fmt.Println("parsed links:", strings.Join(links, ", "))
	fmt.Println(string(fc))
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
