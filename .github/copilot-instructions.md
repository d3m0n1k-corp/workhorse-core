# Workhorse Core - Copilot Instructions

## Project Intent & Vision

Workhorse Core is designed to be a **robust, extensible data conversion system** for browser environments with the following key characteristics:

- **Browser-native**: WebAssembly compilation enables seamless integration into Chrome extensions and web applications
- **Type-safe**: Strong typing with validation ensures data integrity throughout conversion chains
- **Extensible**: Plugin-based architecture allows easy addition of new converters without modifying core system
- **Reliable**: Comprehensive testing with high coverage and fail-fast behavior for production readiness
- **Chainable**: Linked-list based pipeline enables complex multi-step data transformations
- **Developer-friendly**: Clear patterns, auto-registration, and comprehensive error handling

The system serves as the backend engine for data format conversions (JSON ↔ YAML ↔ XML) in browser-based tools, providing both single converter execution and complex conversion chains.

## Architecture Overview

Workhorse is a Chrome extension backend that converts data between formats (JSON, YAML, XML) using a plugin-based converter system. The core is built in Go and compiles to WebAssembly for browser execution.

### Key Components

- **WASM API (`cmd/workhorse.wasm/`)**: Browser interface exposing three global functions: `list_converters`, `execute_converter`, `chain_execute`
- **Converter Registry (`internal/converters/`)**: Plugin system where each converter implements `BaseConverter` interface
- **Chain Execution (`internal/chain/`)**: Linked-list based pipeline for chaining multiple converters
- **Auto-registration (`app/registrations.go`)**: Import-based converter registration using blank imports

## Core Patterns

### Converter Implementation

All converters follow this exact pattern (see `internal/converters/json_prettifier/`):

1. **Struct**: Implements `BaseConverter` with `Apply()`, `InputType()`, `OutputType()` methods
2. **Config**: Separate config struct implementing `BaseConfig` with `Validate()` method
3. **Registration**: `register.go` file using `converters.Register()` with reflection-based constructor
4. **Tests**: Table-driven tests covering `InputType()`, `OutputType()`, `Apply()` success/error cases

### Type Constants

Use predefined constants from `internal/common/types/base.go`: `JSON`, `YAML`, `XML`, `JSON_STRINGIFIED`

### Registry Pattern

- Converters auto-register via blank imports in `app/registrations.go`
- Registration panics on duplicate names (intentional fail-fast behavior)
- Constructor uses reflection: `reflect.New(reg.Config).Interface()`

### Chain Validation

Chain execution validates input/output type compatibility between adjacent converters. The `ConverterList.Validate()` method ensures the output type of converter N matches the input type of converter N+1.

## Development Workflow

### Building

```bash
make wasm      # Build WASM binary to out/wasm/
make clean     # Remove out/ directory
make release   # clean + wasm
```

### Testing

```bash
make test      # Run tests with coverage
make cov       # Generate XML coverage report
make lint      # Run golangci-lint
```

### Adding New Converters

1. Create directory under `internal/converters/[name]/`
2. Implement: `converter.go`, `config.go`, `register.go`, `converter_test.go`, `config_test.go`
3. Add blank import to `app/registrations.go`
4. Follow existing naming: `[format]_to_[format]` or `[format]_[operation]`

## Testing Conventions

### Test Organization

- **Unit Tests**: Files with `_test.go` suffix in the same package/directory as the code they test
  - Examples: `converter_test.go`, `config_test.go` in converter directories
  - Focus on single component/function behavior
- **Integration/System Tests**: Located in `tests/` directory for project-level testing
  - Examples: `tests/chain_execute/`, `tests/converter_list/`, `tests/benchmarks/`
  - Include multi-component integration tests, black-box tests, benchmarks, and end-to-end tests

### Test Implementation

- Use `testify/assert` and `testify/require`
- Mock external dependencies (e.g., `mockableJsonMarshalIndent` in json_prettifier)
- Test both success and error paths for `Apply()` method
- Validate type methods return correct constants
- Benchmark tests should include memory allocation analysis (`-benchmem`)

## Error Handling

- Converters return `(any, error)` from `Apply()`
- Chain execution stops on first error, returning partial results
- Registration errors cause panics (expected behavior)
- Input validation happens in converter `Apply()` method, not in config

## Key Files to Reference

- `internal/converters/json_prettifier/`: Complete converter example
- `internal/chain/converter_list.go`: Chain execution logic
- `cmd/workhorse.wasm/main.go`: WASM entry point and function registration
- `tests/benchmarks/`: Performance benchmarks and optimization validation
- `CONTRIBUTING.md`: Detailed contributor guide with examples
