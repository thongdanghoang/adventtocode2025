package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

/**
* go run gen.go 05
* create day05/main.go, day05/input.txt and day05/README.md
 */

const mainTemplate = `package main

import (
	"adventofcode2025/utils"
	"fmt"
	"time"
)

func main() {
	// lines := utils.ReadLines("input.txt") // line by line
	// grid := utils.ReadGrid("input.txt")   // map 2D [][]rune
	
	input := utils.ReadLines("example-input.txt")
	// input := utils.ReadLines("input.txt")
	
	start := time.Now()
	fmt.Printf("Part 1: %v (took %v)\n", Part1(input), time.Since(start))

	start = time.Now()
	fmt.Printf("Part 2: %v (took %v)\n", Part2(input), time.Since(start))
}

func Part1(input []string) int {
	return 0
}

func Part2(input []string) int {
	return 0
}
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run gen.go <day>")
		fmt.Println("Example: go run gen.go 01")
		os.Exit(1)
	}

	day := os.Args[1]
	if len(day) == 1 {
		day = "0" + day
	}

	dirName := fmt.Sprintf("day%s", day)

	if err := os.MkdirAll(dirName, 0755); err != nil {
		panic(err)
	}

	inputFile := filepath.Join(dirName, "input.txt")
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		os.WriteFile(inputFile, []byte(""), 0644)
		fmt.Printf("Created %s\n", inputFile)
	}
	mainTemplate := strings.Replace(mainTemplate, "utils.ReadLines(\"input.txt\")", "utils.ReadLines(\""+dirName+"/input.txt\")", -1)
	exampleInputFile := filepath.Join(dirName, "example-input.txt")
	if _, err := os.Stat(exampleInputFile); os.IsNotExist(err) {
		os.WriteFile(exampleInputFile, []byte(""), 0644)
		fmt.Printf("Created %s\n", exampleInputFile)
	}
	mainTemplate = strings.Replace(mainTemplate, "utils.ReadLines(\"example-input.txt\")", "utils.ReadLines(\""+dirName+"/example-input.txt\")", -1)

	readmeFile := filepath.Join(dirName, "README.md")
	if _, err := os.Stat(readmeFile); os.IsNotExist(err) {
		os.WriteFile(readmeFile, []byte(""), 0644)
		fmt.Printf("Created %s\n", readmeFile)
	}

	mainFile := filepath.Join(dirName, "main.go")
	if _, err := os.Stat(mainFile); os.IsNotExist(err) {
		tmpl, _ := template.New("main").Parse(mainTemplate)
		f, _ := os.Create(mainFile)
		defer f.Close()
		tmpl.Execute(f, nil)
		fmt.Printf("Created %s\n", mainFile)
	} else {
		fmt.Printf("%s already exists, skipping overwrite.\n", mainFile)
	}

	fmt.Printf("Ready for Day %s! 🚀\n", day)
}
