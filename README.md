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

## Roadmap

### Planned for v0.6.0
- **Custom validation functions**
    - Allow developers to plug in their own validation logic (e.g. `validate:"custom=IsValidUsername"`).
- **Error reporting instead of panic**
    - Option to return `error` from `Build()` instead of panicking (configurable).
- **Better type support**
    - Cover `time.Time`, `uuid.UUID`, and other common types.
- **CLI improvements**
    - Add flags to configure generation (e.g. `--no-panic`, `--error-return`).

---

## Installation
```bash
go install github.com/ericorlovski/go-builder/cmd/gobuilder@latest