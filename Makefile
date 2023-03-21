.DEFAULT_GOAL := build
BINARY_NAME=mygrep

fmt:
	go fmt 
.PHONY:fmt

lint: fmt
		golint 
.PHONY:lint

vet: fmt 
		go vet 
.PHONY:vet

build: vet
		go build -o ${BINARY_NAME} 
.PHONY:build

test:
	go test -v ./...

coverage:
	go test -coverprofile=coverage.out ./... ;    
	go tool cover -func=coverage.out

report:
	go test -coverprofile=coverage.out ./... ;
	go tool cover -html=coverage.out

clean:
	go clean
	if [ -f ${BINARY_NAME} ]; then \
			rm -f ${BINARY_NAME}; \
	fi