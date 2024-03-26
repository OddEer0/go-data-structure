package container

type (
	Container interface {
		Size() int
		Clear()
		IsEmpty() bool
		String() string
	}

	JSONSerializer interface {
		ToJSON() ([]byte, error)
		MarshalJSON() ([]byte, error)
	}

	JSONDeserializer interface {
		FromJSON([]byte) error
		UnmarshalJSON([]byte) error
	}
)
