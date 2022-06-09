package goset

// Set is a set type providing operations like Add, Discard, Union, and
// Disjoint, implemented as a map with empty-struct values.
type Set[T comparable] struct {
	underlyingMap map[T]struct{}
}

func NewSet[T comparable](initialValues... T) Set[T] {
	set := Set[T]{}
	set.underlyingMap = make(map[T]struct{})
	set.Add(initialValues...)
	return set
}

// nothing is the empty struct, a value consuming no memory, which is used
// for all values in the map underlying the set.
var nothing struct{}

// Size returns the number of elements in the set.
func (s Set[T]) Size() int {
	return len(s.underlyingMap)
}

// Add adds each of its arguments to the set.
func (s Set[T]) Add(v... T) {
	for _, e := range v {
		s.underlyingMap[e] = nothing
	}
}

// Has returns true if v is in the set, false if it is not.
func (s Set[T]) Has(v T) bool {
	_, ok := s.underlyingMap[v]
	return ok
}

// Discard removes v from the set.
func (s Set[T]) Discard(v... T) {
	for _, e := range v {
		delete(s.underlyingMap, e)
	}
}

// Union returns a new set containing all elements in the set and all elements
// in other.
func (s Set[T]) Union(other Set[T]) Set[T] {
	new := NewSet[T]()
	for key := range other.underlyingMap {
		new.Add(key)
	}
	for key := range s.underlyingMap {
		new.Add(key)
	}
	return new
}

// Intersection returns a new set containing the elements present in both the
// set and other.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	new := NewSet[T]()
	for key := range s.underlyingMap {
		if other.Has(key) {
			new.Add(key)
		}
	}
	return new
}

// Disjoints returns a new set containing elements exclusive to the set
// and elements exclusive to other, but not elements present in both.
func (s Set[T]) Disjoint(other Set[T]) Set[T] {
	new := NewSet[T]()
	for key := range s.underlyingMap {
		if !other.Has(key) {
			new.Add(key)
		}
	}
	for key := range other.underlyingMap {
		if !s.Has(key) {
			new.Add(key)
		}
	}
	return new
}

// Subtract modifies the set, removing all elements present in other.
func (s Set[T]) Subtract(other Set[T]) {
	for key := range other.underlyingMap {
		s.Discard(key)
	}
}

// Values returns a slice of all values in the set in no particular order.
func (s Set[T]) Values() []T {
	result := make([]T, len(s.underlyingMap))
	index := 0
	for value := range s.underlyingMap {
		result[index] = value
		index++
	}
	return result
}

// EqualTo reports whether the set is equal to `other`. Two sets are equal if
// they contain the same values.
func (s Set[T]) EqualTo(other Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}
	return s.SubsetOf(other)
}

// SubsetOf reports whether the set is a subset of `other`.
func (s Set[T]) SubsetOf(other Set[T]) bool {
	for value := range s.underlyingMap {
		if !other.Has(value) {
			return false
		}
	}
	return true
}

// SupersetOf reports whether the set is a superset of `other`.
func (s Set[T]) SupersetOf(other Set[T]) bool {
	return other.SubsetOf(s)
}

// Extend adds each element in other to the set.
func (s Set[T]) Extend(other Set[T]) {
	for value := range other.underlyingMap {
		s.Add(value)
	}
}

// Clear removes every element from the set.
func (s Set[T]) Clear() {
	for value := range s.underlyingMap {
		s.Discard(value)
	}
}