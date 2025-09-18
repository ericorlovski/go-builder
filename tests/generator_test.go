package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/ericorlovski/go-builder/internal/generator"
	"github.com/ericorlovski/go-builder/internal/model"
)

func TestGenerate(t *testing.T) {
	meta := &model.StructMeta{
		PackageName: "examples",
		Name:        "User",
		Fields: []model.Field{
			{Name: "ID", Type: "int"},
			{Name: "Name", Type: "string"},
		},
	}

	code, err := generator.Generate(meta)
	if err != nil {
		t.Fatalf("Error generating code: %v", err)
	}

	if !contains(code, "WithID") || !contains(code, "WithName") {
		t.Errorf("Generated code missing methods:\n%s", code)
	}

	_ = os.WriteFile(os.TempDir()+"/test_output.go", []byte(code), 0644)
}

func TestGenerateWithDefaults(t *testing.T) {
	def := "42"
	meta := &model.StructMeta{
		PackageName: "examples",
		Name:        "User",
		Fields: []model.Field{
			{Name: "ID", Type: "int", Default: &def},
			{Name: "Name", Type: "string"},
		},
	}

	code, err := generator.Generate(meta)
	if err != nil {
		t.Fatalf("Error generating code: %v", err)
	}

	if !strings.Contains(code, "id: 42") {
		t.Errorf("Default int value not generated:\n%s", code)
	}
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
