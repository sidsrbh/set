# **Go Generic Set Library**

A simple, efficient, and idiomatic **generic Set implementation for Go 1.18+**
Supports any `comparable` type—integers, strings, structs, and more.

---

## **Features**

* Generic: works with any `comparable` type (`int`, `string`, custom struct, etc.)
* Idiomatic Go API: `Add`, `Remove`, `Contains`, `Len`, `Clear`, `Clone`, `Elements`
* Set operations: `Union`, `Intersection`, `Difference`, `IsSubset`, `IsSuperSet`, `Equals`
* Implements `fmt.Stringer` for pretty printing
* Supports JSON marshalling/unmarshalling as a slice/array
* 100% Go: no dependencies, lightweight, ready for production

---

## **Installation**

```sh
go get github.com/sidsrbh/set
```
---

---

## **Usage**

```go
package main

import (
	"fmt"
	"github.com/yourusername/set"
)

func main() {
	// Create a new Set of integers
	intSet := set.New[int]()
	intSet.Add(10)
	intSet.Add(20)
	intSet.Add(30)
	fmt.Println("Set:", intSet) // { 10, 20, 30 }

	// Remove, check, and clone
	intSet.Remove(20)
	fmt.Println("After Remove(20):", intSet) // { 10, 30 }
	fmt.Println("Contains 10?", intSet.Contains(10)) // true

	clone := intSet.Clone()
	fmt.Println("Clone:", clone) // { 10, 30 }

	// Set operations
	another := set.New[int]()
	another.Add(30)
	another.Add(40)

	union := intSet.Union(another)
	fmt.Println("Union:", union) // { 10, 30, 40 }

	intersection := intSet.Intersection(another)
	fmt.Println("Intersection:", intersection) // { 30 }

	diff := intSet.Difference(another)
	fmt.Println("Difference:", diff) // { 10 }

	// Subset, Superset, Equals
	fmt.Println("IsSubset:", intersection.IsSubset(union)) // true
	fmt.Println("IsSuperSet:", union.IsSuperSet(intersection)) // true
	fmt.Println("Equals:", intSet.Equals(clone)) // true

	// JSON Marshalling
	b, _ := json.Marshal(intSet)
	fmt.Println("JSON:", string(b)) // [10,30]
	var s set.Set[int]
	json.Unmarshal(b, &s)
	fmt.Println("Decoded:", &s) // { 10, 30 }
}
```

---

## **API Reference**

### **Set Methods**

* `New[T comparable]() *Set[T]` — create a new set
* `Add(elem T)` — add an element
* `Remove(elem T)` — remove an element
* `Contains(elem T) bool` — check if element exists
* `Len() int` — number of elements
* `Clear()` — remove all elements
* `Clone() *Set[T]` — deep copy of set
* `Elements() []T` — get slice of elements

### **Set Operations**

* `Union(other *Set[T]) *Set[T]` — set union
* `Intersection(other *Set[T]) *Set[T]` — set intersection
* `Difference(other *Set[T]) *Set[T]` — set difference
* `IsSubset(other *Set[T]) bool` — subset check
* `IsSuperSet(other *Set[T]) bool` — superset check
* `Equals(other *Set[T]) bool` — set equality

### **Interfaces**

* Implements `fmt.Stringer` (pretty print)
* Implements `json.Marshaler` and `json.Unmarshaler`

---

## **Supported Types**

Any Go type that is `comparable`:

* Built-ins: `int`, `float64`, `string`, `bool`, etc.
* Structs (no map/slice fields)

**Example:**

```go
type User struct {
    ID   int
    Name string
}

userSet := set.New[User]()
userSet.Add(User{1, "Alice"})
userSet.Add(User{2, "Bob"})
fmt.Println("UserSet:", userSet)
```

---

## **License**

[MIT License](LICENSE)

---

## **Contributing**

* Pull requests, issues, and suggestions welcome!
* Please write tests for any new features or bug fixes.

---

## **Author**

Created by [Saurabh Siddhartha](https://github.com/sidsrbh)

---

**Enjoy a simpler, faster, and more idiomatic set for Go!**

---