package main

import "flag"
import "os"
import "encoding/csv"
import "golang.org/x/net/html"
import "github.com/webdevdata/webdevdata-tools/webdevdata"
import "regexp"
import "fmt"

func main() {
	attrs := flag.String("attrs", "", "CSV list of attributes to print")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "You need to provide a CSS selector and one file")
		os.Exit(-1)
	}

	selector := flag.Arg(0)
	filesChan := make(chan string)

	go webdevdata.GetFiles(filesChan, 1)

	attrList := regexp.MustCompile(",").Split(*attrs, -1)
	csv := csv.NewWriter(os.Stdout)

	for file := range filesChan {
		process(file, selector, attrList, csv)
	}

	csv.Flush()
}

func process(file string, selector string, attrList []string, csv *csv.Writer) {
	webdevdata.ProcessMatchingTags(file, selector, func(node *html.Node) {
		content := []string{file, node.Data}
		for _, attr := range attrList {
			if attr != "" {
				content = append(content, webdevdata.GetAttr(attr, node.Attr))
			}
		}
		csv.Write(content)
	})
}
