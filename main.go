package main

import (
	"context"
	"os"
	"oss.terrastruct.com/d2/d2compiler"
	"oss.terrastruct.com/d2/d2format"
	"oss.terrastruct.com/d2/d2lib"
	"path/filepath"
	"strings"
)

func main() {
	emptyCanvas()
	twoNodes()
	twoNodesUsingD2lib()
}

func emptyCanvas() {
	// Create an empty graph/canvas
	graph, _, _ := d2compiler.Compile("", strings.NewReader(""), nil)

	// Get the Abstract Syntax Tree
	ast := d2format.Format(graph.AST)
	data := []byte(ast)

	// Save it to a d2 file
	_ = os.WriteFile(filepath.Join("results/emptyCanvas.d2"), data, 0600)
}

func twoNodes() {
	// Create a graph/canvas with 2 Nodes
	graph, _, _ := d2compiler.Compile("", strings.NewReader("A,B"), nil)

	// Get the Abstract Syntax Tree
	ast := d2format.Format(graph.AST)
	data := []byte(ast)

	// Save it to a d2 file
	_ = os.WriteFile(filepath.Join("results/twoNodes.d2"), data, 0600)
}

func twoNodesUsingD2lib() {
	ctx := context.Background()

	// Create a graph/canvas with 2 Nodes
	_, graph, _ := d2lib.Compile(ctx, "A,B", nil, nil)

	// Get the Abstract Syntax Tree
	ast := d2format.Format(graph.AST)
	data := []byte(ast)

	// Save it to a d2 file
	_ = os.WriteFile(filepath.Join("results/twoNodesUsingD2lib.d2"), data, 0600)
}
