PHONY: test

test: dep gen
	go test -v ./test --count 1

gen:
	cd src;go generate

dep:
	go mod tidy
