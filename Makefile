.PHONY: release wasm clean lint test ci cov

release : clean wasm

wasm:
	@echo "Building WASM for $(OS)"
ifeq ($(OS),Windows_NT)
	@if not exist out\wasm @mkdir out\wasm
	@set "GOOS=js" && set "GOARCH=wasm" && go build -o out/wasm/ ./api/workhorse.wasm/...
else
	@mkdir -p out/wasm
	@GOOS=js GOARCH=wasm go build -o out/wasm/  ./api/workhorse.wasm/...
endif

build: clean
	@echo "Building All for $(OS)"
ifeq ($(OS),Windows_NT)
	@if not exist out @mkdir out
	@go build -o out/ ./...
else
	@mkdir -p out/wasm
	@go build -o out/  ./...
endif

clean:
	@echo "Cleaning WASM for $(OS)"
ifeq ($(OS),Windows_NT)
	@if exist out rmdir /s /q out
else
	@rm -rf out
endif

lint:
	@echo "Linting WASM for $(OS)"
	@golangci-lint run ./...

test:
	@echo "Testing WASM for $(OS)"
	@go test -coverprofile=coverage.out -covermode=atomic ./...

ci:
	@echo "Testing WASM for $(OS)"
	@go test -coverprofile="coverage.txt" ./...

cov: test
	@echo "Generating coverage report"
	@go install github.com/boumenot/gocover-cobertura@latest
	@gocover-cobertura < coverage.out > cov.xml
