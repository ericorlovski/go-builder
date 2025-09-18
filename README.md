# go-builder

A lightweight code generation tool that provides the **Builder pattern** for Go structs.  
Inspired by Java's Lombok `@Builder`, but implemented natively for Go using `go generate`.

---

## Features
- Automatically generates builder structs for your Go types.
- Chainable `With<Field>()` methods for each struct field.
- `Build()` method to assemble the final struct.
- Supports basic Go types: `int`, `string`, `bool`, `float64`, slices, maps, and pointers.
- Works with `//go:generate` directive.
- Simple CLI tool.

### Supported Struct Tags
- `default:"..."` → assign default values.
- `required:"true"` → enforce presence at build time.
- `omitempty:"true"` → assign only if non-zero.
- `validate:"..."` → simple validation rules:
    - `min=<value>` → minimum numeric value
    - `max=<value>` → maximum numeric value
    - `email` → must contain `@`

---

## Installation
```bash
go install github.com/ericorlovski/go-builder/cmd/gobuilder@latest