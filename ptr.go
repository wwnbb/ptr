package ptr

func Ptr[T any](value T) (pointer *T) { return &value }
