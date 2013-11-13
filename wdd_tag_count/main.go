package main

import "flag"
import "os"
import "encoding/csv"
import "code.google.com/p/go.net/html"
import "github.com/webdevdata/webdevdata-tools/webdevdata"
import "strconv"

func main() {
  flag.Parse()
  file := flag.Arg(0)
  tags := make(map[string]int)

  webdevdata.ProcessTags(file, func (token html.Token) {
    tags[token.Data]++
  })

  csv := csv.NewWriter(os.Stdout)
  for tag, count := range tags {
    csv.Write([]string{file, tag, strconv.Itoa(count)})
  }
  csv.Flush()
}

