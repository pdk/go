package maps

// Realloc rebuilds a map and returns the new map/copy. Since a map can never
// shrink, if you want to reclaim memory for deleted mappings, you need to
// rebuild the map.
func Realloc[T1 comparable, T2 any](m map[T1]T2) map[T1]T2 {

	newMap := make(map[T1]T2,len(m))

	for k,v := range m {
		newMap[k] = v
	}

	return newMap
}

// Apply returns a new map containing key, value pairs returned by the function f.
func Apply[T1 comparable, T2 any](m map[T1]T2, f func(T1,T2) (T1,T2)) map[T1]T2 {

	newMap := make(map[T1]T2)

	for k,v := range m {
		k, v = f(k,v)
		newMap[k] = v
	}

	return newMap

}

// Filter returns a new map containing only those key, value pairs where the
// filter function returns true.
func Filter[T1 comparable, T2 any](m map[T1]T2, filter func(T1,T2) bool) map[T1]T2 {

	newMap := make(map[T1]T2)

	for k,v := range m {
		if filter(k,v) {
			newMap[k] = v
		}
	}

	return newMap

}

// Count returns the number of entries where the function match returns true.
func Count[T1 comparable, T2 any](m map[T1]T2, match func(T1,T2) bool) int {

	c := 0

	for k,v := range m {
		if match(k,v) {
			c++
		}
	}

	return c
}