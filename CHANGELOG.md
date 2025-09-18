# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [v0.1.0] - 2025-09-18
### Added
- Initial release of `go-builder`.
- Support for generating Builder pattern for Go structs:
    - Automatically creates `With<Field>()` methods for each struct field.
    - Supports basic types (`int`, `string`, `bool`, `float64`).
    - Generates a `Build()` method for assembling the final struct.
- Added support for the `//go:generate` directive for code generation.
- CLI tool `gobuilder` with arguments:
    - `-type` — struct name,
    - `-file` — path to the input `.go` file,
    - `-output` — path for the generated builder file.
- Examples provided in the `examples/` directory.

---

## [Unreleased]
### Planned
- Support for default values via struct tag `default:"..."`.
- Validation of required fields in `Build()`.
- Support for `omitempty`.
- Deep builder generation for nested structs.
- Test coverage for more field types (`[]string`, `map`, pointers).

---

## [v0.2.0] - 2025-09-19
### Added
- Support for default values via struct tag `default:"..."`.
    - Example:
      ```go
      type User struct {
          ID    int    `default:"1"`
          Name  string `default:"Anonymous"`
      }
      ```
    - Generates:
      ```go
      func NewUserBuilder() *UserBuilder {
          return &UserBuilder{
              id:   1,
              name: "Anonymous",
          }
      }
      ```
- Improved lowercase handling (`ID` → `id`, `UUID` → `uuid`).

### Fixed
- Removed duplicate `New<Type>Builder()` constructor from generated code.
