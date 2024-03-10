test:
	go test ./...

cover:
	go test ./... -v -coverpkg=./... -coverprofile=c.out
	go tool cover -html="c.out"
	rm c.out

go:
	go run main.go