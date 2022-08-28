export GO111MODULE=on
export GOSUMDB=off

dep:
	go mod tidy

build: dep
	go build -o flakeid .
