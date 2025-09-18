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

	if !contains(code, "ID") || !contains(code, "Name") {
		t.Errorf("Generated code missing methods:\n%s", code)
	}

	_ = os.WriteFile(os.TempDir()+"/test_output.go", []byte(code), 0644)
}

func TestGenerateDefaults(t *testing.T) {
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

func TestGenerateOmitempty(t *testing.T) {
	meta := &model.StructMeta{
		PackageName: "examples",
		Name:        "User",
		Fields: []model.Field{
			{Name: "ID", Type: "int"},
			{Name: "Name", Type: "string", Omitempty: true},
		},
	}

	code, err := generator.Generate(meta)
	if err != nil {
		t.Fatalf("Error generating code: %v", err)
	}

	if !strings.Contains(code, "if b.name != \"\"") {
		t.Errorf("Omitempty not generated for Name field:\n%s", code)
	}
}

func TestGenerateCollections(t *testing.T) {
	meta := &model.StructMeta{
		PackageName: "examples",
		Name:        "User",
		Fields: []model.Field{
			{Name: "Tags", Type: "[]string", Omitempty: true},
			{Name: "Data", Type: "map[string]int", Omitempty: true},
			{Name: "Ref", Type: "*int", Omitempty: true},
		},
	}

	code, err := generator.Generate(meta)
	if err != nil {
		t.Fatalf("Error generating code: %v", err)
	}

	checks := []string{
		"if len(b.tags) > 0 {",
		"if b.data != nil && len(b.data) > 0 {",
		"if b.ref != nil {",
	}

	for _, c := range checks {
		if !strings.Contains(code, c) {
			t.Errorf("Missing check: %s\n%s", c, code)
		}
	}
}

func TestGenerateValidation(t *testing.T) {
	meta := &model.StructMeta{
		PackageName: "examples",
		Name:        "User",
		Fields: []model.Field{
			{Name: "Age", Type: "int", Validate: strPtr("min=18,max=99")},
			{Name: "Email", Type: "string", Validate: strPtr("email")},
		},
	}

	code, err := generator.Generate(meta)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	checks := []string{
		"if b.age < 18 || b.age > 99 { panic(\"validation failed for b.age\") }",
		"if !strings.Contains(b.email, \"@\") { panic(\"validation failed for b.email\") }",
	}

	for _, c := range checks {
		if !strings.Contains(code, c) {
			t.Errorf("Missing validation: %s\n%s", c, code)
		}
	}
}

func strPtr(s string) *string {
	return &s
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
