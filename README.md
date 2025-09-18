# go-builder

A lightweight code generation tool that provides the **Builder pattern** for Go structs.  
Inspired by Java's Lombok `@Builder`, but implemented natively for Go using `go generate`.

---

## Features
- Automatically generates builder structs for your Go types.
- Chainable `<Field>()` methods for each struct field.
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
## Examples

### Basic Example
```go
type User struct {
    ID    int    `required:"true" default:"1"`
    Name  string `omitempty:"true"`
    Email string `validate:"email"`
}

// go:generate ../bin/gobuilder -type=User -file=user.go -output=user_builder.go

func main() {
    u := NewUserBuilder().
        ID(1).
        Name("Alice").
        Email("alice@example.com").
        Build()

    fmt.Println(u)
    // Output: {ID:1 Name:"Alice" Email:"alice@example.com"}
}
```

---

## Installation
```bash
go install github.com/ericorlovski/go-builder/cmd/gobuilder@latest