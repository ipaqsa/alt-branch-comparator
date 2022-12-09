build:
	go build -o ./bin/altcomparator ./cmd/main.go

install: build
	chmod u+s ./bin/altcomparator