package main

import (
	"encoding/json"
	"fmt"

	"github.com/sidsrbh/set"
)

func main() {
	// Create a new set of integers
	intSet := set.New[int]()
	intSet.Add(10)
	intSet.Add(20)
	intSet.Add(30)
	intSet.Add(20) // Duplicate, will be ignored

	fmt.Println("IntSet:", intSet) // Uses Stringer interface

	// Remove an element
	intSet.Remove(10)
	fmt.Println("After Remove(10):", intSet)

	// Check membership
	fmt.Println("Contains 20?", intSet.Contains(20)) // true
	fmt.Println("Contains 10?", intSet.Contains(10)) // false

	// Set Length
	fmt.Println("Length:", intSet.Len())

	// Get all elements as a slice (order is not guaranteed)
	fmt.Println("Elements:", intSet.Elements())

	// Clone the set
	cloneSet := intSet.Clone()
	fmt.Println("Cloned Set:", cloneSet)

	// Create another set for operations
	anotherSet := set.New[int]()
	anotherSet.Add(20)
	anotherSet.Add(40)
	anotherSet.Add(50)

	// Union
	union := intSet.Union(anotherSet)
	fmt.Println("Union:", union)

	// Intersection
	intersection := intSet.Intersection(anotherSet)
	fmt.Println("Intersection:", intersection)

	// Difference (elements in intSet but not in anotherSet)
	difference := intSet.Difference(anotherSet)
	fmt.Println("Difference:", difference)

	// Subset & Superset
	fmt.Println("IsSubset?", intersection.IsSubset(union))
	fmt.Println("IsSuperSet?", union.IsSuperSet(intersection))

	// Equality
	fmt.Println("intSet == cloneSet?", intSet.Equals(cloneSet))
	fmt.Println("intSet == anotherSet?", intSet.Equals(anotherSet))

	// Marshal set to JSON
	b, err := json.Marshal(intSet)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
	} else {
		fmt.Println("JSON Marshal:", string(b))
	}

	// Unmarshal set from JSON
	var s set.Set[int]
	err = json.Unmarshal(b, &s)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	} else {
		fmt.Println("JSON Unmarshal (decoded):", &s)
	}

	// Clear all elements
	intSet.Clear()
	fmt.Println("After Clear:", intSet)

	//-------------------------------------------------
	// Example: Set of strings (works with any comparable type)
	strSet := set.New[string]()
	strSet.Add("apple")
	strSet.Add("banana")
	strSet.Add("cherry")
	strSet.PowerSet()
	fmt.Println("String Set:", strSet)
}
