package numberconstraints

import (
	"golang.org/x/exp/constraints"
)

func Add[T constraints.Ordered](x, y T) T {
	return x + y
}

func Subtract[T constraints.Ordered](x, y T) T {
	return x - y
}

func Multiply[T constraints.Ordered](x, y T) T {
	return x * y
}

func Divide[T constraints.Ordered](x, y T) T {
	return x / y
}
