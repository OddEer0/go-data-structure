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

	EnumWithKey[T any, K any, Q any] interface {
		Each(func(key T, val K))
		EachLast(func(key T, val K))
		Some(func(key T, val K) bool) bool
		Every(func(key T, val K) bool) bool
		Map(func(key T, val K) (T, K)) Q
		Filter(func(key T, val K) bool) Q
		Concat(others ...Q) Q
	}

	EnumWitIndex[K any, Q any] interface {
		Each(func(index int, val K))
		EachLast(func(index int, val K))
		Some(func(index int, val K) bool) bool
		Every(func(index int, val K) bool) bool
		Map(func(index int, val K) K) Q
		Filter(func(index int, val K) bool) Q
		Concat(others ...Q) Q
	}
)
