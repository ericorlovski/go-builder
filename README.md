# go-builder

A lightweight code generation tool that provides the **Builder pattern** for Go structs.  
Inspired by Java's Lombok `@Builder`, but implemented natively for Go using `go generate`.

---

## Features
- Automatically generates builder structs for your Go types.
- Chainable `With<Field>()` methods for each struct field.
- `Build()` method to assemble the final struct.
- Supports basic Go types: `int`, `string`, `bool`, `float64`.
- Works with `//go:generate` directive.
- Simple CLI tool.
- **New in v0.2.0:** Default values via struct tags (`default:"..."`).
- **New in v0.3.0:**
    - Required fields via struct tags (`required:"true"`).
    - Omitempty support via struct tags (`omitempty:"true"`).

---

## Installation
```bash
go install github.com/ericorlovski/go-builder/cmd/gobuilder@latest