package goutils

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Keys() []K {
	keys := []K{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m Map[K, V]) Vals() []V {
	vals := []V{}
	for _, v := range m {
		vals = append(vals, v)
	}
	return vals
}
