package helper

func Present[T comparable](value T) bool {
	return value == *new(T)
}

func Empty[T comparable](value []T) bool {
	return value[0] == *new(T)
}
