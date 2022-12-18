package number

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | float32 | float64
}

func Add[T Number](x, y T) T {
	return x + y
}

func Subtract[T Number](x, y T) T {
	return x - y
}

func Multiply[T Number](x, y T) T {
	return x * y
}

func Divide[T Number](x, y T) T {
	return x / y
}
