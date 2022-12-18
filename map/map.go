package map

import "golang.org/x/exp/constraints"

func Map[T constraints.Ordered](values []T, mapFunc func (T) T) []T {
	values []T
	for _, v := range values {
		val := mapFunc(v)
		values = append(values, val)
	}
	return values
}
