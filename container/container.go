package container

type (
	Container interface {
		Size() int
		Clear()
		IsEmpty() bool
		String() string
	}
)
