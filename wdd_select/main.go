package main

import "flag"
import "os"
import "encoding/csv"
import "code.google.com/p/go.net/html"
import "github.com/webdevdata/webdevdata-tools/webdevdata"
import "regexp"

func main() {
  attrs    := flag.String("attrs","","CSV list of attributes to print")
  flag.Parse()
  selector := flag.Arg(0)
  file     := flag.Arg(1)

  attrList := regexp.MustCompile(",").Split(*attrs, -1)
  csv := csv.NewWriter(os.Stdout)
  webdevdata.ProcessMatchingTags(file, selector, func (node *html.Node) {
    content := []string{file, node.Data}
    for _, attr := range attrList {
      if attr != "" {
        content = append(content, webdevdata.GetAttr(attr, node.Attr))
      }
    }
    csv.Write(content)
  })

  csv.Flush()
}

