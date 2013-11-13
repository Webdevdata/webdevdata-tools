package main

import "flag"
import "os"
import "encoding/csv"
import "code.google.com/p/go.net/html"
import "github.com/ernesto-jimenez/webdevdata-tools/webdevdata"

func main() {
  flag.Parse()
  file     := flag.Arg(0)

  csv := csv.NewWriter(os.Stdout)
  webdevdata.ProcessTags(file, func (token html.Token) {
    if (token.Data == "meta") {
      name    := webdevdata.GetAttr("name", token.Attr)
      if name != "" {
        csv.Write([]string{file, name})
      }
    }
  })

  csv.Flush()
}

