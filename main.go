package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lukasjarosch/go-docx"
)

func main() {
	content, err := os.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	headers := strings.Fields(strings.TrimSpace(lines[0]))

	for _, line := range lines[1:] {
		replaceMap := docx.PlaceholderMap{}
		fields := strings.Fields(line)
		for i, header := range headers {
			replaceMap[header] = fields[i]
		}

		// read and parse the template docx
		doc, err := docx.Open("template.docx")
		if err != nil {
			panic(err)
		}

		// replace the keys with values from replaceMap
		err = doc.ReplaceAll(replaceMap)
		if err != nil {
			panic(err)
		}

		// write out a new file
		fileName := fmt.Sprintf("%s-%s.docx", fields[0], fields[1])
		err = doc.WriteToFile(fileName)
		if err != nil {
			panic(err)
		}
	}
}
