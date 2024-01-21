.PHONY: parse

parse:
	@echo "Parsing currencies..."
	@go run main.go
	@echo "Done."