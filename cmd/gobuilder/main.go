package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ericorlovski/go-builder/internal/generator"
	"github.com/ericorlovski/go-builder/internal/parser"
)

func main() {
	structName := flag.String("type", "", "Struct name to generate builder for")
	inputFile := flag.String("file", "", "Path to Go file containing the struct")
	outputFile := flag.String("output", "", "Path to output generated file")
	flag.Parse()

	if *structName == "" || *inputFile == "" || *outputFile == "" {
		fmt.Println("Usage: gobuilder -type=StructName -file=path/to/file.go -output=builder.go")
		os.Exit(1)
	}

	meta, err := parser.ParseStruct(*inputFile, *structName)
	if err != nil {
		fmt.Println("Error parsing struct:", err)
		os.Exit(1)
	}

	code, err := generator.Generate(meta)
	if err != nil {
		fmt.Println("Error generating builder:", err)
		os.Exit(1)
	}

	if err := os.WriteFile(*outputFile, []byte(code), 0644); err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}

	fmt.Println("Builder generated at", *outputFile)
}
