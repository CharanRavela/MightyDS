# Lint target for formatting and static analysis
lint: fmt vet revive staticcheck

# Formatting check
fmt:
	@echo "Running go fmt..."
	go fmt ./...

# Vet check
vet:
	@echo "Running go vet..."
	go vet ./...

# Revive linting
revive:
	@echo "Running revive linting..."
	revive -formatter stylish -config .revive.toml ./...

# Staticcheck linting
staticcheck:
	@echo "Running staticcheck linting..."
	staticcheck ./...
