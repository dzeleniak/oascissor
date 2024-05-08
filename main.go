package main

import (
	"oascissor/graph"
	"os"
)

type OpenAPI struct {
	Components struct {
		Schemas map[string]interface{} `json:"schemas"`
	} `json:"components"`
}

type Graph map[string][]string

func main() {
	//// Read OpenAPI specification file

	if len(os.Args) != 2 {
		panic("No file provided...")
	}

	filePath := os.Args[1]

	b, err := os.ReadFile(filePath)

	if err != nil {
		panic("Couldn't read file...")
	}

	g := graph.New(b)
	g.Dump()
}
