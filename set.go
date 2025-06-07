package set

import (
	"encoding/json"
	"fmt"
)

type Set[T comparable] struct {
	data map[T]struct{}
}

// New Empty Set
func New[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

// Add new Element
func (s *Set[T]) Add(elem T) {
	s.data[elem] = struct{}{}
}

// remove an element from set
func (s *Set[T]) Remove(elem T) {
	delete(s.data, elem)
}

// returns true if the element exists in set
func (s *Set[T]) Contains(elem T) bool {
	_, exists := s.data[elem]
	return exists
}

// returns length of a Set
func (s *Set[T]) Len() int {
	return len(s.data)
}

// returns a slice of all elements in the set
func (s *Set[T]) Elements() []T {
	elems := make([]T, 0, s.Len())
	for key, _ := range s.data {
		elems = append(elems, key)
	}
	return elems
}

//remove all elements from the set

func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}

// returns a clone of set
func (s *Set[T]) Clone() *Set[T] {
	newSet := New[T]()
	for key, _ := range s.data {
		newSet.Add(key)
	}
	return newSet
}

// Other Set Operations
// Union AUB
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := s.Clone()
	for key, _ := range other.data {
		result.Add(key)
	}
	return result
}

// intersection A()B
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := s.Clone()
	for key, _ := range result.data {
		if !other.Contains(key) {
			result.Remove(key)
		}
	}
	return result
}

// Difference => Elements in A but not in B
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	results := New[T]()
	for key, _ := range s.data {
		if !other.Contains(key) {
			results.Add(key)
		}
	}
	return results
}

// IsSubset returns true if A is subset of B, else return false
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for key, _ := range s.data {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// IsSuperSet returns true if A is Super Set of B, else return false
func (s *Set[T]) IsSuperSet(other *Set[T]) bool {
	return other.IsSubset(s)
}

func (s *Set[T]) Equals(other *Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for key, _ := range other.data {
		if !s.Contains(key) {
			return false
		}
	}
	return true
}

func (s *Set[T]) String() string {
	elems := s.Elements()
	setString := ""
	for _, val := range elems {
		setString += fmt.Sprintf("%v, ", val)
	}
	return "{ " + setString + " }"
}

// JSON interfaces
func (s *Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Elements())
}
func (s *Set[T]) UnmarshalJSON(data []byte) error {
	var elems []T
	if err := json.Unmarshal(data, &elems); err != nil {
		return err
	}
	s.Clear()
	for _, e := range elems {
		s.Add(e)
	}
	return nil
}
