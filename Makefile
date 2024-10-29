# Default values for directory
dir := .

# Target to run the Go application
run:
	@echo "Running main.go in $(dir)..."
	@cd $(dir) && go run main.go

# Target to run tests in the specified directory
test:
	@echo "Running tests in $(dir)..."
	@cd $(dir) && go test -v

# Target to list files in the specified directory
list:
	@echo "Listing files in $(dir):"
	@ls -l $(dir)

# Target to clean up any generated files
clean:
	@echo "Cleaning up in $(dir)..."
	@cd $(dir) && rm -f *.out  # Adjust according to your needs

# Help target to display usage instructions
help:
	@echo "Makefile commands:"
	@echo "  make run dir=directory         # Run the Go application in the specified directory"
	@echo "  make test dir=directory        # Run tests in the specified directory"
	@echo "  make list dir=directory        # List files in the specified directory"
	@echo "  make clean dir=directory       # Clean up generated files in the specified directory"
	@echo "  make help                      # Display this help message"

# Allow default commands when no arguments are given
.DEFAULT_GOAL := help
