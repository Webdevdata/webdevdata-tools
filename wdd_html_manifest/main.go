package main

import "flag"
import "os"
import "encoding/csv"
import "code.google.com/p/go.net/html"
import "github.com/webdevdata/webdevdata-tools/webdevdata"

func main() {
  flag.Parse()
  file     := flag.Arg(0)

  csv := csv.NewWriter(os.Stdout)
  selector := "html[manifest]" // all html tags with manifest attribute
  webdevdata.ProcessMatchingTags(file, selector, func (node *html.Node) {
    manifest    := webdevdata.GetAttr("manifest", node.Attr)
    if manifest != "" {
      csv.Write([]string{file, manifest})
    }
  })

  csv.Flush()
}

