lint:
	@echo "# Linting..."
	@echo "## Go mod tidy..."
	@go mod tidy
	@echo "## Fixing whitespaces..."
	@wsl --allow-cuddle-declarations --force-err-cuddling --fix ./...
	@echo "## Running golangci-lint..."
	@golangci-lint run
