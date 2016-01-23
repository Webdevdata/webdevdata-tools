package main

import "flag"
import "os"
import "encoding/csv"
import "golang.org/x/net/html"
import "github.com/webdevdata/webdevdata-tools/webdevdata"
import "strconv"

func main() {
	flag.Parse()
	filesChan := make(chan string)

	go webdevdata.GetFiles(filesChan, 0)
	csv := csv.NewWriter(os.Stdout)

	for file := range filesChan {
		tags := make(map[string]int)
		webdevdata.ProcessTags(file, func(token html.Token) {
			tags[token.Data]++
		})

		for tag, count := range tags {
			csv.Write([]string{file, tag, strconv.Itoa(count)})
		}
	}
	csv.Flush()
}
