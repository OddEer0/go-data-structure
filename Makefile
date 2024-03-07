rbtree:
	go test ./pkg/redblacktree/... -v -coverpkg=./pkg/redblacktree/... -coverprofile=c.out
	go tool cover -html="c.out"
	rm c.out