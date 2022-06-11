package lists

type MapKeyType interface {
	string | int | int32 | int64 | interface{}
}

type KeyGetter[T any] func(i T) MapKeyType

type arrays[T any] struct {
	list []T
}

func (arrays *arrays[T]) Map(keyGetter func(i T) MapKeyType) map[MapKeyType]T {
	result := make(map[MapKeyType]T)
	for _, item := range arrays.list {
		result[keyGetter(item)] = item
	}
	return result
}

func (arrays *arrays[T]) Filter(predicate func(i T) bool) []T {
	result := make([]T, 0, len(arrays.list))
	for _, i := range arrays.list {
		if predicate(i) {
			result = append(result, i)
		}
	}
	return result
}

func (arrays *arrays[T]) GroupBy(keyGetter func(i T) MapKeyType) map[MapKeyType][]T {
	result := make(map[MapKeyType][]T)
	for _, item := range arrays.list {
		key := keyGetter(item)
		group, _ := result[key]
		result[key] = append(group, item)
	}
	return result
}

func Arrays[T any](list []T) *arrays[T] {
	return &arrays[T]{list: list}
}
