package main

import (
	"log"
	"strings"

	"github.com/doraiot/battery/elastic"
)

func main() {
	log.SetFlags(0)
	// elastic.IndexRead()
	elastic.Bulk()
	log.Println(strings.Repeat("▁", 66))
}
