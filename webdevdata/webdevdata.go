package webdevdata

import "code.google.com/p/go.net/html"
import "os"
import "io"
import "fmt"

func ProcessTags(file string, process func(html.Token)) {
  html_reader, err := os.Open(file)
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
  d := html.NewTokenizer(html_reader)
  for {
    // token type
    tokenType := d.Next()
    if tokenType == html.ErrorToken {
      err := d.Err()
      if err != io.EOF {
        fmt.Println(err)
        os.Exit(-1)
      } else {
        return
      }
    }
    token := d.Token()
    switch tokenType {
      case html.StartTagToken, html.SelfClosingTagToken: // <tag>
      process(token)
    }
  }
  return
}

func GetAttr(key string, attrs []html.Attribute) (string) {
  for _, attr := range attrs {
    if attr.Key == key {
      return attr.Val
    }
  }
  return ""
}

