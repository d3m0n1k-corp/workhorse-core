# Contributing to workhorse-core

Thank you for your interest in contributing to workhorse-core! We welcome contributions and improvements to this project. This guide will help you get started and explain our file structure, workflow, and code standards—especially when creating new converters.

## Table of Contents
- [Contributing to workhorse-core](#contributing-to-workhorse-core)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
  - [Project Structure](#project-structure)
  - [Creating a New Converter](#creating-a-new-converter)
  - [Code Standards](#code-standards)
  - [Testing](#testing)
  - [Submission Guidelines](#submission-guidelines)

## Getting Started

- **Fork the Repository:** Begin by forking the repository and cloning your fork locally.
- **Install Dependencies:** Ensure you have [Go 1.24](https://golang.org/dl/) installed along with required Go modules.
- **Run Tests Locally:** Use the command `go test ./...` to run the tests and verify your environment.

## Project Structure

Below is an overview of the key directories and files:

```
workhorse-core/
├── .github/
│   ├── workflows/
│   │   ├── go.yml             # CI/CD for building & releasing
│   │   └── cov.yml            # Code coverage configuration
├── cmd/
│   ├── cli/
│   │   └── main.go            # CLI entrypoint
│   └── workhorse.wasm/
│       ├── main.go            # WASM entrypoint
│       ├── logging/
│       │   └── writer.go
│       └── operations/
│           ├── execute_converter.go
│           └── list_converters.go
├── app/
│   ├── list.go               # Converter listing functionality
│   ├── execute.go            # Converter execution functionality
│   └── registrations.go      # Auto-register all converters by import
├── internal/
│   ├── common/
│   │   ├── linked_list/
│   │   │   ├── list.go
│   │   │   └── list_test.go
│   │   └── types/
│   │       └── base.go
│   └── converters/
│       ├── registry.go       # Register, list and create converters
│       ├── registry_test.go
│       ├── json_prettifier/
│       │   ├── converter.go
│       │   ├── converter_test.go
│       │   ├── config.go
│       │   ├── config_test.go
│       │   └── register.go
│       ├── json_to_yaml/
│       │   ├── converter.go
│       │   ├── converter_test.go
│       │   ├── config.go
│       │   ├── config_test.go
│       │   └── register.go
│       └── yaml_to_json/
│           ├── converter.go
│           ├── converter_test.go
│           ├── config.go
│           ├── config_test.go
│           └── register.go
├── go.mod
├── go.sum
└── Makefile
```

## Creating a New Converter

When adding a new converter, please follow these guidelines:

1. **Directory Structure:**
   - Create a new directory under `internal/converters` named after your converter (e.g., `my_converter`).
   - Add the following files:
     - `converter.go` – Implements the converter logic using the `BaseConverter` interface.
     - `config.go` – Defines a configuration struct implementing the `BaseConfig` interface and a `Validate()` method.
     - `register.go` – Registers your converter by calling `converters.Register(&converters.Registration{ ... })`.  
     - Tests for both, the converter and config (`converter_test.go` and `config_test.go`).

2. **Registration:**
   - In your `register.go` file, create a registration that includes:
     - `Name`: Unique name of your converter.
     - `DemoInput`: An example input value.
     - `Description`: A short description.
     - `Config`: The reflect type of your config struct.
     - `InputType` and `OutputType`: Valid values defined in `internal/common/types/base.go`.
     - `Constructor`: A function that takes a `BaseConfig` and returns your converter instance.

   *Example snippet from an existing registration:*
   ```go
   // In register.go
   package my_converter

   import (
       "reflect"
       "workhorse-core/internal/common/types"
       "workhorse-core/internal/converters"
   )

   var _ = converters.Register(&converters.Registration{
       Name:        "my_converter",
       DemoInput:   []byte(`{"example": "data"}`),
       Description: "MyConverter converts input from X to Y.",
       Config:      reflect.TypeOf(MyConverterConfig{}),
       InputType:   types.JSON,
       OutputType:  types.YAML,
       Constructor: func(config converters.BaseConfig) converters.BaseConverter {
           return &MyConverter{config: config.(MyConverterConfig)}
       },
   })
   ```

3. **Code Organization:**
   - **Converters should implement:**
     - `Apply(input any) (any, error)`
     - `InputType() string`
     - `OutputType() string`
   - **Configurations:**
     - Must implement a `Validate() error` method.
     - Use proper struct tags for JSON and validations (using `github.com/go-playground/validator/v10`).

## Code Standards

- **Formatting:** Use Go’s standard formatting (`gofmt`) and lint with `golangci-lint`.
- **Error Handling:** Always check for errors and return helpful messages.
- **Tests:** Provide unit tests for both the converter logic and the configuration validation.
- **Documentation:** Include comments where applicable and update this CONTRIBUTING.md if guidelines change.

## Testing

- Run unit tests with:
  ```
  go test ./...
  ```
- CI workflows are configured in `.github/workflows`. Please ensure your changes pass all tests locally before submitting a pull request.
- You can use `make cov` to generate a cobertura coverage file that can be then used using vscode's `coverage gutters` extension.

## Submission Guidelines

- Open an issue if you plan to implement a major new feature or refactor.
- Fork the repository and create a new branch for your changes.
- Create a pull request describing your contribution and reference any related issues.
- Follow the code standards and file structure guidelines detailed above.

Thank you for contributing to workhorse-core. Your efforts help improve the quality and usability of the project for everyone!