package datastructures

type flagError string

func (e flagError) Error() string {
	return string(e)
}

const (
	OutOfBoundsError = flagError("index out of bounds")
)
