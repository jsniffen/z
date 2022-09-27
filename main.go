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
  fn := generateFilename()

  tags := make([]string, 0)

  for _, word := range strings.Split(s, " ") {
    if strings.HasPrefix(word, "@") {
      tags = append(tags, word)
    }
  }

  fmt.Println("parsed tags", tags)

  err := os.WriteFile(fn, []byte(s), 0777)
  if err != nil {
    fmt.Println("Error creating snippet", err)
    os.Exit(1)
  }
  fmt.Println("created snippet at", fn)
  os.Exit(0)
}

func help() {
  fmt.Println("help")
  os.Exit(0)
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

