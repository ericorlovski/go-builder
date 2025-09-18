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
		Name: "User",
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

	// Дополнительно можно записать во временный файл и прогнать go vet
	_ = os.WriteFile("test_output.go", []byte(code), 0644)
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
