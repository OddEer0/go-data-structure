rbtree:
	go test ./tree/redblacktree/... -v -coverpkg=./tree/redblacktree/... -coverprofile=c.out
	go tool cover -html="c.out"
	rm c.out

go:
	go run main.go