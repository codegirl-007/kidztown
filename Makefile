.PHONY: format lint

format:
	gofmt -w .

lint:
	@files=$$(gofmt -l .); \
	if [ -n "$$files" ]; then \
		echo "gofmt needed:"; \
		echo "$$files"; \
		exit 1; \
	fi
	go vet ./...
