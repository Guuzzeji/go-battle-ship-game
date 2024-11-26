package game

// GetKeys returns a slice of all keys in the given map.
// The order of the keys is not defined.
func GetKeys[T comparable, V any](m map[T]V) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
