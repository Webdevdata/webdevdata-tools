package main

import "flag"
import "os"
import "encoding/csv"
import "golang.org/x/net/html"
import "github.com/webdevdata/webdevdata-tools/webdevdata"

func main() {
	flag.Parse()
	file := flag.Arg(0)

	csv := csv.NewWriter(os.Stdout)
	selector := "meta[name]" // All meta tags with name attribute
	webdevdata.ProcessMatchingTags(file, selector, func(node *html.Node) {
		name := webdevdata.GetAttr("name", node.Attr)
		csv.Write([]string{file, name})
	})

	csv.Flush()
}
