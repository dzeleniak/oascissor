package graph

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Graph map[string][]string

func New(byteValue []byte) Graph {
	// Parse JSON into OpenAPI struct
	var openAPI OpenAPI
	err := json.Unmarshal(byteValue, &openAPI)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	// Create graph
	graph := make(Graph)

	// Iterate through schemas and build graph
	for schemaName, schema := range openAPI.Components.Schemas {
		refs := findReferences(schema)
		graph[schemaName] = refs
	}

	return graph
}

// Recursive function to find references within a schema
func findReferences(schema interface{}) []string {
	refs := make([]string, 0)

	switch v := schema.(type) {
	case map[string]interface{}:
		if ref, ok := v["$ref"].(string); ok {
			trimmedRef := strings.TrimPrefix(ref, "#/components/schemas/")
			refs = append(refs, trimmedRef)
		} else {
			for _, value := range v {
				refs = append(refs, findReferences(value)...)
			}
		}
	case []interface{}:
		for _, value := range v {
			refs = append(refs, findReferences(value)...)
		}
	}

	return refs
}

func (g Graph) DFS(start string) {
	visited := make(map[string]bool)

	var traverse func(s string)

	traverse = func(s string) {
		visited[s] = true

		fmt.Println(s)

		for _, v := range g[s] {
			val, exists := visited[v]
			if !exists || !val {
				traverse(v)
			}
		}
	}

	traverse(start)

}

func (g Graph) Dump() {
	// Print the graph
	fmt.Println("Graph of references:")
	for schemaName, refs := range g {
		fmt.Printf("%s -> [\n", schemaName)
		for _, ref := range refs {
			fmt.Printf("\t%v\n", ref)
		}
		fmt.Printf("]\n\n")
	}
}
