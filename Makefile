.PHONY: test
test:
	go test ./...

# Usage: export DAY=1 && make build-darwin && dist/day$DAY
.PHONY: build-darwin
build-darwin:
	go build -o dist/day$(DAY) ./cmd/day$(DAY) 
