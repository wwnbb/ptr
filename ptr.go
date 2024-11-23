package ptr

func ptr[T any](value T) (pointer *T) { return &value }
