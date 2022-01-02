package slices

import (
	"constraints"
)

// Equal compares two slices of the same type with ==. If any elements are not
// ==, returns false.
func Equal[T comparable](values1, values2 []T) bool {
	if len(values1) != len(values2) {
		return false
	}

	for i, v := range values1 {
		if v != values2[i] {
			return false
		}
	}

	return true
}

// Compare returns true if every comparison with the compare func returns true.
// This provides a way to check if two slices are equal when the elements are
// not comparable, by passing in an equality function.
func Compare[T any](values1, values2 []T, compare func(T,T) bool) bool {
	if len(values1) != len(values2) {
		return false
	}

	for i, v := range values1 {
		if !compare(v, values2[i]) {
			return false
		}
	}

	return true
}

// Every returns true if f is true for every element.
func Every[T any](values []T, f func(T) bool) bool {

	for _, e := range values {
		if !f(e) {
			return false
		}
	}

	return true
}

// Any returns true if f is true for any element. (Shortcut for IndexF() != -1.)
func Any[T any](values []T, f func(T) bool) bool {
	return IndexF(values, f) != -1
}

// Apply creates a new slice having the values produced by f applied to the
// elements of values.
func Apply[T1 any, T2 any](values []T1, f func(T1) T2) []T2 {

	to := make([]T2, len(values), len(values))

	for i, e := range values {
		to[i] = f(e)
	}

	return to
}

// Map is another name for Apply.
func Map[T1 any, T2 any](values []T1, f func(T1) T2) []T2 {
	return Apply(values, f)
}

// Reduce combines all elements of a slice to produce a single value.
func Reduce[T any](values []T, reduce func(T,T) T) T {
	
	var t T

	if len(values) == 0 {
		return t // whatever the zero value is
	}

	t = values[0]

	for i := 1; i<len(values) ; i++ {
		t = reduce(t, values[i])
	}

	return t
}

// Sum produces the sum of any Integer type slice elements.
func Sum[T constraints.Integer](values []T) T {
	return Reduce(values, func(a,b T) T {
		return a+b
	})
}

// Copy creates a new copy of the input slice.
func Copy[T any](values []T) []T {
	ns := make([]T, len(values), 0)
	return append(ns, values...)
}

// Filter returns a new slice containing only those values where the function filter
// returns true.
func Filter[T any](values []T, filter func(T) bool) []T {

	var n []T

	for _, e := range values {
		if filter(e) {
			n = append(n, e)
		}
	}

	return n
}

// Index returns the first index where the element equals (==) the item. Returns
// -1 if there is no matching element.
func Index[T comparable](values []T, item T) int {

	for i, e := range values {
		if e == item {
			return i
		}
	}

	return -1
}

// IndexF returns the first index where function match returns true. Returns -1
// if there is no matching element.
func IndexF[T any](values []T, match func(T) bool) int {

	for i, e := range values {
		if match(e) {
			return i
		}
	}

	return -1
}

// Count returns the count of elements in values where the function match
// returns true. Returns 0 if none match.
func Count[T any](values []T, match func(T) bool) int {
	
	c:=0
	
	for _, e := range values {
		if match(e) {
			c++
		}
	}

	return c
}

// HeadTail splits a slice into the first element, and a slice containing the
// rest of the elements.
func HeadTail[T any](values []T) (T, []T) {
	if len(values) == 0 {
		var t T
		return t, []T(nil) // return zero values if empty
	}

	return values[0], values[1:]
}