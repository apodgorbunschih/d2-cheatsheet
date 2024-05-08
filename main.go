package main

import (
	"os"
	"oss.terrastruct.com/d2/d2compiler"
	"oss.terrastruct.com/d2/d2format"
	"oss.terrastruct.com/d2/d2oracle"
	"path/filepath"
	"strings"
)

func main() {
	emptyCanvas()
	twoNodes()
	threeNodes()
	threeNodesGroup()
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
	graph, _, _ := d2compiler.Compile("", strings.NewReader("A;B"), nil)

	// Get the Abstract Syntax Tree
	ast := d2format.Format(graph.AST)
	data := []byte(ast)

	// Save it to a d2 file
	_ = os.WriteFile(filepath.Join("results/twoNodes.d2"), data, 0600)
}

func threeNodes() {
	// Create an empty graph/canvas
	graph, _, _ := d2compiler.Compile("", strings.NewReader(""), nil)

	// Create Nodes
	graph, _, _ = d2oracle.Create(graph, nil, "A")
	graph, _, _ = d2oracle.Create(graph, nil, "B")
	graph, _, _ = d2oracle.Create(graph, nil, "C")

	// Create Edges
	graph, _, _ = d2oracle.Create(graph, nil, "A->B")
	graph, _, _ = d2oracle.Create(graph, nil, "A->C")

	// Get the Abstract Syntax Tree
	ast := d2format.Format(graph.AST)
	data := []byte(ast)

	// Save it to a d2 file
	_ = os.WriteFile(filepath.Join("results/threeNodes.d2"), data, 0600)
}

func threeNodesGroup() {
	// Create an empty graph/canvas
	graph, _, _ := d2compiler.Compile("", strings.NewReader(""), nil)

	// Create Nodes
	graph, _, _ = d2oracle.Create(graph, nil, "A")

	// Create Group
	graph, _, _ = d2oracle.Create(graph, nil, "D")

	// Grouping
	graph, _, _ = d2oracle.Create(graph, nil, "D.B")
	graph, _, _ = d2oracle.Create(graph, nil, "D.C")

	// Create Edges
	graph, _, _ = d2oracle.Create(graph, nil, "A->D.B")
	graph, _, _ = d2oracle.Create(graph, nil, "A->D.C")

	// Get the Abstract Syntax Tree
	ast := d2format.Format(graph.AST)
	data := []byte(ast)

	// Save it to a d2 file
	_ = os.WriteFile(filepath.Join("results/threeNodesGroup.d2"), data, 0600)
}
