

rbtree:
	go test ./pkg/rb_tree/... -v -coverpkg=./pkg/rb_tree/... -coverprofile=c.out
	go tool cover -html="c.out"
	rm c.out