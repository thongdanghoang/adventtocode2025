package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/tabwriter"
)

type Row map[string]string

type Table struct {
	Name    string
	Headers []string
	Rows    []Row
}

func main() {
	csvMode := flag.Bool("csv", false, "Output in CSV format")
	filePath := flag.String("file", "day00/input.txt", "Path to input file")
	flag.Parse()

	contentBytes, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	content := string(contentBytes)

	tables := parseTables(content)

	for _, tbl := range tables {
		if *csvMode {
			printCSV(tbl)
		} else {
			printASCII(tbl)
		}
		fmt.Println()
	}
}

func parseTables(content string) []Table {
	var tables []Table

	reTable := regexp.MustCompile(`(?s)<TABLE name="([^"]+)".*?>(.*?)</TABLE>`)
	matches := reTable.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		tableName := match[1]
		innerXML := match[2]

		innerXML = strings.ReplaceAll(innerXML, `<?xml version="1.0"?>`, "")
		innerXML = strings.ReplaceAll(innerXML, "<ROWSET>", "")
		innerXML = strings.ReplaceAll(innerXML, "</ROWSET>", "")

		validXML := "<ROOT>" + innerXML + "</ROOT>"

		rows, headers := parseRowsFromXML(validXML)

		if len(rows) > 0 {
			tables = append(tables, Table{
				Name:    tableName,
				Headers: headers,
				Rows:    rows,
			})
		}
	}
	return tables
}

func parseRowsFromXML(xmlContent string) ([]Row, []string) {
	decoder := xml.NewDecoder(strings.NewReader(xmlContent))
	var rows []Row
	headerMap := make(map[string]bool)
	var headers []string

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "ROW" {
				currentRow := make(Row)
				for {
					token, _ := decoder.Token()
					if token == nil {
						break
					}
					switch subSe := token.(type) {
					case xml.StartElement:
						tagName := subSe.Name.Local

						var value string
						nextToken, _ := decoder.Token()
						if cd, ok := nextToken.(xml.CharData); ok {
							value = string(cd)
							decoder.Token()
						} else if _, ok := nextToken.(xml.EndElement); ok {
							value = ""
						}

						currentRow[tagName] = strings.TrimSpace(value)

						if !headerMap[tagName] {
							headerMap[tagName] = true
							headers = append(headers, tagName)
						}
					case xml.EndElement:
						if subSe.Name.Local == "ROW" {
							goto EndRow
						}
					}
				}
			EndRow:
				rows = append(rows, currentRow)
			}
		}
	}
	sort.Strings(headers)
	return rows, headers
}

func printASCII(t Table) {
	fmt.Printf("TABLE: %s\n", t.Name)
	if len(t.Rows) == 0 {
		fmt.Println("(No rows)")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	for _, h := range t.Headers {
		fmt.Fprintf(w, "%s\t", h)
	}
	fmt.Fprintln(w)

	for _, r := range t.Rows {
		for _, h := range t.Headers {
			val := r[h]
			if len(val) > 50 {
				val = val[:47] + "..."
			}
			fmt.Fprintf(w, "%s\t", val)
		}
		fmt.Fprintln(w)
	}
	w.Flush()
	fmt.Println(strings.Repeat("-", 80))
}

func printCSV(t Table) {
	fmt.Printf("--- CSV FOR TABLE: %s ---\n", t.Name)
	fmt.Println(strings.Join(t.Headers, ","))
	for _, r := range t.Rows {
		var line []string
		for _, h := range t.Headers {
			val := strings.ReplaceAll(r[h], "\"", "\"\"")
			if strings.Contains(val, ",") || strings.Contains(val, "\n") {
				val = "\"" + val + "\""
			}
			line = append(line, val)
		}
		fmt.Println(strings.Join(line, ","))
	}
}
