release : clean wasm
wasm: clean
	@echo "Building WASM for $(OS)"
ifeq ($(OS),Windows_NT)
	@if not exist out\wasm @mkdir out\wasm
	@set "GOOS=js" && set "GOARCH=wasm" && go build -o out/wasm/ ./api/workhorse.wasm/...
else
	@mkdir -p out/wasm
	@GOOS=js GOARCH=wasm go build -o out/wasm/  ./api/workhorse.wasm/...
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
	@go test ./...

ci:
	@echo "Testing WASM for $(OS)"
	@go test -coverprofile="coverage.txt" ./...