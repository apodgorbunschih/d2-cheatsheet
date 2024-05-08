package main

import (
	"os"
	"oss.terrastruct.com/d2/d2compiler"
	"oss.terrastruct.com/d2/d2format"
	"oss.terrastruct.com/d2/d2oracle"
	"oss.terrastruct.com/util-go/go2"
	"path/filepath"
	"strings"
)

func main() {
	emptyCanvas()
	twoNodes()
	threeNodes()
	threeNodesGroup()
	cloudsExample()
	shapes()
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

func cloudsExample() {
	// Create an empty graph/canvas
	graph, _, _ := d2compiler.Compile("", strings.NewReader(""), nil)

	// Create Groups
	graph, _, _ = d2oracle.Create(graph, nil, "clouds")
	graph, _, _ = d2oracle.Create(graph, nil, "clouds.aws")
	graph, _, _ = d2oracle.Create(graph, nil, "clouds.gcloud")

	graph, _, _ = d2oracle.Create(graph, nil, "clouds.aws.load_balancer -> clouds.aws.api")
	graph, _, _ = d2oracle.Create(graph, nil, "clouds.aws.api -> clouds.aws.db")

	graph, _, _ = d2oracle.Create(graph, nil, "clouds.gcloud.auth -> clouds.gcloud.db")

	graph, _, _ = d2oracle.Create(graph, nil, "clouds.gcloud -> clouds.aws")

	graph, _, _ = d2oracle.Create(graph, nil, "users-> clouds.aws.load_balancer")
	graph, _, _ = d2oracle.Create(graph, nil, "users-> clouds.gcloud.auth")

	graph, _, _ = d2oracle.Create(graph, nil, "ci.deploy-> clouds")

	// Add labels
	graph, _ = d2oracle.Set(graph, nil, "clouds.aws.label", nil, go2.Pointer("AWS"))
	graph, _ = d2oracle.Set(graph, nil, "clouds.gcloud.label", nil, go2.Pointer("Gcloud"))

	// Get the Abstract Syntax Tree
	ast := d2format.Format(graph.AST)
	data := []byte(ast)

	// Save it to a d2 file
	_ = os.WriteFile(filepath.Join("results/cloudsExample.d2"), data, 0600)
}

func shapes() {
	// Create an empty graph/canvas
	graph, _, _ := d2compiler.Compile("", strings.NewReader(""), nil)

	// Create Nodes and shapes
	graph, _, _ = d2oracle.Create(graph, nil, "clouds")
	graph, _ = d2oracle.Set(graph, nil, "clouds.shape", nil, go2.Pointer("cloud"))

	graph, _, _ = d2oracle.Create(graph, nil, "diamond")
	graph, _ = d2oracle.Set(graph, nil, "diamond.shape", nil, go2.Pointer("diamond"))

	graph, _, _ = d2oracle.Create(graph, nil, "circle")
	graph, _ = d2oracle.Set(graph, nil, "circle.shape", nil, go2.Pointer("circle"))

	// Get the Abstract Syntax Tree
	ast := d2format.Format(graph.AST)
	data := []byte(ast)

	// Save it to a d2 file
	_ = os.WriteFile(filepath.Join("results/shapes.d2"), data, 0600)
}
