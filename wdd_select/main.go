package main

import "flag"
import "os"
import "encoding/csv"
import "code.google.com/p/go.net/html"
import "github.com/webdevdata/webdevdata-tools/webdevdata"
import "regexp"
import "fmt"
import "bufio"

func main() {
  attrs    := flag.String("attrs","","CSV list of attributes to print")
  flag.Parse()

  if flag.NArg() == 0 {
    fmt.Fprintln(os.Stderr, "You need to provide a CSS selector and one file")
    os.Exit(-1)
  }

  selector := flag.Arg(0)
  filesChan:= make(chan string)

  go getFiles(filesChan)

  attrList := regexp.MustCompile(",").Split(*attrs, -1)
  csv := csv.NewWriter(os.Stdout)

  run(filesChan, selector, attrList, csv)

  csv.Flush()
}

func run(filesChan chan string, selector string, attrList []string, csv *csv.Writer) {
  for file := range filesChan {
    process(file, selector, attrList, csv)
  }
}

func getFiles(filesChan chan string) {
  if flag.NArg() > 1 {
    files := flag.Args()
    for i, file := range files {
      if i == 0 { continue }
      filesChan <- file
    }
  } else {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
      filesChan <- scanner.Text()
    }
    if err := scanner.Err(); err != nil {
      fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
  }
  close(filesChan)
}

func process(file string, selector string, attrList []string, csv *csv.Writer){
  webdevdata.ProcessMatchingTags(file, selector, func (node *html.Node) {
    content := []string{file, node.Data}
    for _, attr := range attrList {
      if attr != "" {
        content = append(content, webdevdata.GetAttr(attr, node.Attr))
      }
    }
    csv.Write(content)
  })
}

