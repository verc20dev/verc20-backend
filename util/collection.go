package util

func GroupingBy[K comparable, T any](vs []T, f func(T) K) map[K][]T {
	group := make(map[K][]T)
	for _, v := range vs {
		k := f(v)
		group[k] = append(group[k], v)
	}
	return group
}


func Keys[K comparable, T any](x map[K]T) []K {
	res := make([]K, len(x))
	i := 0
	for k := range x {
		res[i] = k
		i += 1
	}
	return res
}

func Values[K comparable, T any](x map[K]T) []T {
	res := make([]T, len(x))
	i := 0
	for _, v := range x {
		res[i] = v
		i += 1
	}
	return res
}

func MapValues[K comparable, T any, R any](x map[K]T, f func(T) R) map[K]R {
	res := make(map[K]R)
	for k, v := range x {
		res[k] = f(v)
	}
	return res
}

func FindFirst[T any](x []T, f func(T) bool) (T, bool) {
	for _, v := range x {
		if f(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

func IsEmpty[T any](x []T) bool {
	return x != nil && len(x) == 0
}

func Map[T any, R any](x []T, f func(T) R) []R {
	res := make([]R, len(x))
	for i, v := range x {
		res[i] = f(v)
	}
	return res
}

