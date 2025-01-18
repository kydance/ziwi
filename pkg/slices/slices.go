package slices

import (
	"math/rand"
	"sort"

	"golang.org/x/exp/constraints"
)

// Clonable defines a constraint of types having Clone() T method.
type Clonable[T any] interface {
	Clone() T
}

// Filter 函数接受一个集合和一个谓词函数，返回一个由集合中满足谓词函数条件的元素组成的新集合。
func Filter[T any, Slice ~[]T](slice Slice, predicate func(item T, index int) bool) Slice {
	ret := make(Slice, 0, len(slice))

	for idx, item := range slice {
		if predicate(item, idx) {
			ret = append(ret, item)
		}
	}

	return ret
}

// Map manipulates a slice and transforms it to a slice of another type.
// 操作切片并将其转换为另一种类型的切片
func Map[From any, To any](slice []From, iteratee func(item From, index int) To) []To {
	ret := make([]To, len(slice))

	for idx, item := range slice {
		ret[idx] = iteratee(item, idx)
	}

	return ret
}

// FilterMap returns a slice which obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
//
// 返回使用给定回调函数进行过滤和映射后获得的切片。
// 回调函数应返回两个值：
//   - 映射操作的结果和
//   - 是否应包含结果元素
func FilterMap[From any, To any](slice []From, callback func(item From, index int) (To, bool)) []To {
	ret := []To{}

	for idx, item := range slice {
		if r, ok := callback(item, idx); ok {
			ret = append(ret, r)
		}
	}

	return ret
}

// FlatMap manipulates a slice and transforms and flattens it to a slice of another type.
// The transform function can either return a slice or a `nil`, and in the `nil` case
// no value is added to the final slice.
// 操作切片并将其转换并展平为另一种类型的切片。
// 变换函数可以返回切片或“nil”，在“nil”情况下没有值添加到最终切片。
func FlatMap[From any, To any](slice []From, iteratee func(item From, index int) []To) []To {
	ret := make([]To, 0, len(slice))

	for idx, item := range slice {
		ret = append(ret, iteratee(item, idx)...)
	}

	return ret
}

// Reduce reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
// The initial value serves as the first argument to the first call of the accumulator.
//
// Example:
//
//	sum := Reduce([]int{1, 2, 3}, func(agg int, item int, index int) int {
//	    return agg + item
//	}, 0) // Returns 6
func Reduce[From any, To any](slice []From, accumulator func(agg To, item From, index int) To, initial To) To {
	if len(slice) == 0 {
		return initial
	}

	result := initial
	for i := range slice {
		result = accumulator(result, slice[i], i)
	}

	return result
}

// ReduceRight is like Reduce except that it iterates over elements of collection from right to left.
// This is particularly useful when the order of operations matters, such as function composition or
// string operations.
//
// Example:
//
//	sentence := ReduceRight([]string{"!", "World", " ", "Hello"}, func(agg string, item string, index int) string {
//	    return agg + item
//	}, "") // Returns "Hello World!"
func ReduceRight[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	if len(collection) == 0 {
		return initial
	}

	result := initial
	for i := len(collection) - 1; i >= 0; i-- {
		result = accumulator(result, collection[i], i)
	}

	return result
}

// ForEach iterates over elements of slice and invokes iteratee for each element.
// The iteration is performed in order, from index 0 to len(slice)-1.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	ForEach(numbers, func(item int, index int) {
//	    fmt.Printf("Item at %d: %d\n", index, item)
//	})
func ForEach[T any](slice []T, iteratee func(item T, index int)) {
	if len(slice) == 0 {
		return
	}

	for idx := range slice {
		iteratee(slice[idx], idx)
	}
}

// ForEachWhile iterates over elements of slice and invokes iteratee for each element.
// The iteration continues until either all elements are processed or iteratee returns false.
// Returns true if all elements were processed, false if iteration was stopped early.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	ForEachWhile(numbers, func(item int, index int) bool {
//	    if item > 3 {
//	        return false // stop iteration
//	    }
//	    fmt.Printf("Item at %d: %d\n", index, item)
//	    return true // continue iteration
//	})
func ForEachWhile[T any](slice []T, iteratee func(item T, index int) bool) bool {
	if len(slice) == 0 {
		return true
	}

	for idx, item := range slice {
		if !iteratee(item, idx) {
			return false
		}
	}
	return true
}

// Times invokes the iteratee function count times and returns the resulting slice.
// If count is negative, returns an empty slice.
//
// Example:
//
//	// Generate sequence
//	numbers := Times(3, func(idx int) int {
//	    return idx * 2
//	}) // Returns []int{0, 2, 4}
//
//	// Generate string slice
//	labels := Times(2, func(idx int) string {
//	    return fmt.Sprintf("Item %d", idx)
//	}) // Returns []string{"Item 0", "Item 1"}
func Times[T any](count int, iteratee func(idx int) T) []T {
	if count <= 0 {
		return make([]T, 0)
	}

	ret := make([]T, count)
	for i := range ret {
		ret[i] = iteratee(i)
	}
	return ret
}

// Uniq returns a duplicate-free version of an array, in which only the first occurrence of each
// element is kept. The order of result values is determined by the order they occur in the array.
//
//	Usage:
//		Uniq([]int{1, 2, 2, 1}) -> []int{1, 2}
func Uniq[T comparable, Slice ~[]T](slice Slice) Slice {
	ret := make(Slice, 0, len(slice))
	seen := make(map[T]struct{}, len(slice))

	for _, item := range slice {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		ret = append(ret, item)
	}

	return ret
}

// UniqBy returns a duplicate-free version of an array, using the iteratee function to determine uniqueness.
// The iteratee is invoked with one argument: (value). The first occurrence of each unique key is kept.
// The order of result values is determined by the order they occur in the array.
//
// Example:
//
//	// Remove duplicates by a specific field
//	people := []Person{{Name: "Alice", Age: 25}, {Name: "Bob", Age: 25}}
//	unique := UniqBy(people, func(p Person) int { return p.Age })
//	// Returns [{Name: "Alice", Age: 25}]
//
//	// Custom uniqueness criteria
//	numbers := []int{1, -1, 2, -2}
//	unique := UniqBy(numbers, func(n int) int { return abs(n) })
//	// Returns [1, 2]
func UniqBy[T any, U comparable, Slice ~[]T](slice Slice, iteratee func(item T) U) Slice {
	if len(slice) == 0 {
		return make(Slice, 0)
	}

	ret := make(Slice, 0, len(slice))
	seen := make(map[U]struct{}, len(slice))

	for _, item := range slice {
		key := iteratee(item)
		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		ret = append(ret, item)
	}

	return ret
}

// GroupBy creates an object composed of keys generated from running each element of slice
// through iteratee. The corresponding value of each key is a slice of elements that generated the key.
// Elements are collected in the order they appear in the input slice.
//
// Example:
//
//	// Group by length
//	words := []string{"one", "two", "three"}
//	groups := GroupBy(words, func(s string) int { return len(s) })
//	// Returns map[3:["one", "two"], 5:["three"]]
//
//	// Group by first letter
//	names := []string{"Alice", "Bob", "Amy"}
//	groups := GroupBy(names, func(s string) string { return s[:1] })
//	// Returns map["A":["Alice", "Amy"], "B":["Bob"]]
func GroupBy[T any, U comparable, Slice ~[]T](slice Slice, iteratee func(item T) U) map[U]Slice {
	if len(slice) == 0 {
		return make(map[U]Slice)
	}

	// Pre-allocate with a reasonable size
	ret := make(map[U]Slice, min(len(slice), 10))

	for _, item := range slice {
		key := iteratee(item)
		if existing, ok := ret[key]; ok {
			ret[key] = append(existing, item)
		} else {
			// Start with capacity 4 for small groups, can grow if needed
			ret[key] = append(make(Slice, 0, 4), item)
		}
	}

	return ret
}

// Chunk splits a slice into multiple slices of specified size.
// The last chunk may contain fewer elements if len(slice) is not evenly divisible by size.
// If size is less than or equal to 0, it will panic.
// Each chunk is a view into the original slice to avoid unnecessary copying.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	chunks := Chunk(numbers, 2)
//	// Returns [][]int{{1, 2}, {3, 4}, {5}}
//
//	words := []string{"a", "b", "c", "d"}
//	chunks := Chunk(words, 3)
//	// Returns [][]string{{"a", "b", "c"}, {"d"}}
func Chunk[T any, Slice ~[]T](slice Slice, size int) []Slice {
	if size <= 0 {
		panic("chunk size must be greater than 0")
	}

	if len(slice) == 0 {
		return []Slice{}
	}

	// Calculate number of chunks needed
	n := (len(slice) + size - 1) / size

	// Pre-allocate result slice
	chunks := make([]Slice, n)

	// Fill chunks
	for i := 0; i < n-1; i++ {
		start := i * size
		chunks[i] = slice[start : start+size : start+size]
	}

	// Handle last chunk separately to avoid bounds check
	start := (n - 1) * size
	chunks[n-1] = slice[start:len(slice):len(slice)]

	return chunks
}

// PartitionBy splits a slice into sub-slices based on the result of the iteratee function.
// Elements that produce the same key when passed through the iteratee are grouped together.
// The order of elements within each group is preserved, and groups appear in the order of
// their first occurrence in the input slice.
//
// Example:
//
//	// Group numbers by parity
//	numbers := []int{1, 2, 3, 4, 5}
//	groups := PartitionBy(numbers, func(n int) bool {
//	    return n%2 == 0
//	}) // Returns [][]int{{1, 3, 5}, {2, 4}}
//
//	// Group strings by length
//	words := []string{"one", "two", "three", "four"}
//	groups := PartitionBy(words, func(s string) int {
//	    return len(s)
//	}) // Returns [][]string{{"one", "two"}, {"four"}, {"three"}}
func PartitionBy[T any, U comparable, Slice ~[]T](slice Slice, iteratee func(item T) U) []Slice {
	if len(slice) == 0 {
		return []Slice{}
	}

	// Pre-allocate with a reasonable size
	ret := make([]Slice, 0, min(len(slice), 10))
	seen := make(map[U]int, min(len(slice), 10))

	// Pre-allocate the first group with a reasonable size
	ret = append(ret, make(Slice, 0, min(len(slice), 8)))

	// Process first element to initialize the first group
	firstKey := iteratee(slice[0])
	seen[firstKey] = 0
	ret[0] = append(ret[0], slice[0])

	// Process remaining elements
	for i := 1; i < len(slice); i++ {
		item := slice[i]
		key := iteratee(item)

		if retIdx, ok := seen[key]; ok {
			// Append to existing group
			ret[retIdx] = append(ret[retIdx], item)
		} else {
			// Create new group
			retIdx := len(ret)
			seen[key] = retIdx
			// Pre-allocate new group
			ret = append(ret, make(Slice, 0, min(len(slice)-i, 8)))
			ret[retIdx] = append(ret[retIdx], item)
		}
	}

	return ret
}

// Flatten combines a slice of slices into a single slice by concatenating all elements
// in order. The function preserves the order of elements both between and within slices.
// Returns an empty slice if the input is empty or contains only empty slices.
//
// Example:
//
//	numbers := [][]int{{1, 2}, {3, 4}, {5, 6}}
//	flattened := Flatten(numbers)
//	// Returns []int{1, 2, 3, 4, 5, 6}
//
//	mixed := [][]int{{1}, {}, {2, 3}, {}, {4}}
//	flattened := Flatten(mixed)
//	// Returns []int{1, 2, 3, 4}
//
//	empty := [][]int{}
//	flattened := Flatten(empty)
//	// Returns []int{}
func Flatten[T any, Slice ~[]T](slice []Slice) Slice {
	if len(slice) == 0 {
		return make(Slice, 0)
	}

	// Calculate total length and check if we have any non-empty slices
	totalLen := 0
	hasElements := false
	for _, item := range slice {
		if len(item) > 0 {
			hasElements = true
			totalLen += len(item)
		}
	}

	if !hasElements {
		return make(Slice, 0)
	}

	// Pre-allocate result slice with exact capacity
	ret := make(Slice, 0, totalLen)

	// Append all elements at once for each non-empty slice
	for _, item := range slice {
		if len(item) > 0 {
			ret = append(ret, item...)
		}
	}

	return ret
}

// Interleave combines multiple slices by alternating their elements.
// Elements are taken from each slice in order until all elements are used.
// If slices have different lengths, shorter slices are skipped once exhausted.
//
// Example:
//
//	a := []int{1, 2, 3}
//	b := []int{4, 5}
//	c := []int{6, 7, 8}
//	result := Interleave(a, b, c)
//	// Returns []int{1, 4, 6, 2, 5, 7, 3, 8}
//
//	// With empty slices
//	result := Interleave([]int{1}, []int{}, []int{2, 3})
//	// Returns []int{1, 2, 3}
//
//	// Single slice
//	result := Interleave([]int{1, 2, 3})
//	// Returns []int{1, 2, 3}
func Interleave[T any, Slice ~[]T](slices ...Slice) Slice {
	if len(slices) == 0 {
		return Slice{}
	}

	if len(slices) == 1 {
		return append(make(Slice, 0, len(slices[0])), slices[0]...)
	}

	// Calculate total length and find longest slice
	maxLen := 0
	totalLen := 0
	nonEmpty := 0
	for _, slice := range slices {
		if len(slice) > 0 {
			nonEmpty++
		}
		totalLen += len(slice)
		if len(slice) > maxLen {
			maxLen = len(slice)
		}
	}

	if totalLen == 0 {
		return Slice{}
	}

	// Pre-allocate result slice
	ret := make(Slice, totalLen)
	retIdx := 0

	// Interleave elements
	for i := 0; i < maxLen; i++ {
		for _, slice := range slices {
			if i < len(slice) {
				ret[retIdx] = slice[i]
				retIdx++
			}
		}
	}

	return ret[:retIdx]
}

// Shuffle randomly permutes the elements in the slice using the default random source.
// It modifies the original slice and returns it. This is a Fisher-Yates shuffle
// implementation with O(n) time complexity.
//
// If the slice is empty or has only one element, it returns immediately.
// The function preserves the type of the input slice.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	Shuffle(numbers)
//	// Returns []int{2, 3, 5, 1, 4}
func Shuffle[T any, Slice ~[]T](slice Slice) Slice {
	if len(slice) <= 1 {
		return slice
	}
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}

// Reverse reverses the order of elements in the slice in-place and returns it.
// The algorithm uses the two-pointer technique with O(n/2) swaps and O(1) extra space.
//
// If the slice is empty or has only one element, it returns immediately.
// The function preserves the type of the input slice.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	Reverse(numbers)
//	// Returns []int{5, 4, 3, 2, 1}
func Reverse[T any, Slice ~[]T](slice Slice) Slice {
	if len(slice) <= 1 {
		return slice
	}
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// Fill fills elements of slice with initial value.
// The function returns a new slice with the same length as the input slice,
// where each element is a clone of the initial value.
//
// The initial value must implement the Clonable interface to ensure proper copying.
// This function is useful when you need multiple independent copies of an object.
//
// Example:
//
//	type Person struct {
//	    Name string
//	}
//
//	p := Person{Name: "Alice"}
//	people := Fill([]Person{{}, {}, {}}, p)
//	// Returns []Person{{Name: "Alice"}, {Name: "Alice"}, {Name: "Alice"}}
func Fill[T Clonable[T]](slice []T, initial T) []T {
	if len(slice) == 0 {
		return []T{}
	}

	ret := make([]T, len(slice))
	for i := range ret {
		ret[i] = initial.Clone()
	}
	return ret
}

// Repeat creates a new slice with count copies of the initial value.
// Each element in the resulting slice is a clone of the initial value,
// ensuring that modifications to one element don't affect others.
//
// If count is less than or equal to 0, returns an empty slice.
// The initial value must implement the Clonable interface.
//
// Example:
//
//	p := Person{Name: "Alice"}
//	people := Repeat(3, p)
//	// Returns []Person{{Name: "Alice"}, {Name: "Alice"}, {Name: "Alice"}}
func Repeat[T Clonable[T]](count int, initial T) []T {
	if count <= 0 {
		return []T{}
	}

	ret := make([]T, count)
	for i := range ret {
		ret[i] = initial.Clone()
	}
	return ret
}

// RepeatBy creates a new slice of length count where each element is generated
// by calling the predicate function with the element's index.
//
// If count is less than or equal to 0, returns an empty slice.
// The predicate function is called exactly once for each index from 0 to count-1.
//
// Example:
//
//	numbers := RepeatBy(3, func(idx int) int { return idx * 2 })
//	// Returns []int{0, 2, 4}
func RepeatBy[T any](count int, predicate func(idx int) T) []T {
	if count <= 0 {
		return []T{}
	}

	ret := make([]T, count)
	for i := range ret {
		ret[i] = predicate(i)
	}
	return ret
}

// KeyBy creates a map from a slice using a key selector function.
// For each element in the slice, the iteratee function determines its key in the resulting map.
// If multiple elements produce the same key, later elements will overwrite earlier ones.
//
// Example:
//
//	users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
//	userMap := KeyBy(users, func(u User) int { return u.ID })
//	// Results in map[int]User{1: {ID: 1, Name: "Alice"}, 2: {ID: 2, Name: "Bob"}}
func KeyBy[K comparable, V any](slice []V, iteratee func(item V) K) map[K]V {
	if len(slice) == 0 {
		return make(map[K]V)
	}
	ret := make(map[K]V, len(slice))
	for _, item := range slice {
		ret[iteratee(item)] = item
	}
	return ret
}

// Associate transforms a slice into a map by applying a transform function to each element.
// The transform function returns both the key and value for each element in the resulting map.
// If multiple elements produce the same key, later elements will overwrite earlier ones.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	result := Associate(numbers, func(n int) (string, bool) {
//	    return fmt.Sprintf("num%d", n), n%2 == 0
//	})
//	// Results in map[string]bool{"num1": false, "num2": true, "num3": false}
func Associate[T any, K comparable, V any](slice []T, transform func(item T) (K, V)) map[K]V {
	if len(slice) == 0 {
		return make(map[K]V)
	}
	ret := make(map[K]V, len(slice))
	for _, item := range slice {
		k, v := transform(item)
		ret[k] = v
	}
	return ret
}

// SliceToMap is an alias for Associate function. It transforms a slice into a map by applying
// a transform function to each element. The transform function returns both the key and value
// for each element in the resulting map.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	result := SliceToMap(numbers, func(n int) (string, int) {
//	    return fmt.Sprintf("key%d", n), n * 2
//	})
//	// Results in map[string]int{"key1": 2, "key2": 4, "key3": 6}
func SliceToMap[T any, K comparable, V any](slice []T, transform func(item T) (K, V)) map[K]V {
	return Associate(slice, transform)
}

// Drop returns a slice with n elements dropped from the beginning.
// If n is greater than or equal to the length of the slice, returns an empty slice.
// The function preserves the type of the input slice.
//
// Example:
//
//	nums := []int{1, 2, 3, 4, 5}
//	result := Drop(nums, 2) // Returns []int{3, 4, 5}
func Drop[T any, Slice ~[]T](slice Slice, n int) Slice {
	if n <= 0 {
		ret := make(Slice, len(slice))
		copy(ret, slice)
		return ret
	}

	if len(slice) <= n {
		return make(Slice, 0)
	}

	ret := make(Slice, 0, len(slice)-n)
	return append(ret, slice[n:]...)
}

// DropRight returns a slice with n elements dropped from the end.
// If n is greater than or equal to the length of the slice, returns an empty slice.
// The function preserves the type of the input slice.
//
// Example:
//
//	nums := []int{1, 2, 3, 4, 5}
//	result := DropRight(nums, 2) // Returns []int{1, 2, 3}
func DropRight[T any, Slice ~[]T](slice Slice, n int) Slice {
	if n <= 0 {
		ret := make(Slice, len(slice))
		copy(ret, slice)
		return ret
	}

	if len(slice) <= n {
		return make(Slice, 0)
	}

	ret := make(Slice, 0, len(slice)-n)
	return append(ret, slice[:len(slice)-n]...)
}

// DropWhile returns a slice with elements dropped from the beginning while the predicate returns true.
// Once the predicate returns false for an element, that element and all subsequent elements are kept.
// The function preserves the type of the input slice.
//
// Example:
//
//	nums := []int{2, 4, 6, 7, 8, 10}
//	result := DropWhile(nums, func(n int) bool {
//	    return n%2 == 0
//	}) // Returns []int{7, 8, 10}
func DropWhile[T any, Slice ~[]T](slice Slice, predicate func(item T) bool) Slice {
	if len(slice) == 0 {
		return make(Slice, 0)
	}

	idx := 0
	for ; idx < len(slice) && predicate(slice[idx]); idx++ {
	}

	if idx == len(slice) {
		return make(Slice, 0)
	}

	ret := make(Slice, 0, len(slice)-idx)
	return append(ret, slice[idx:]...)
}

// DropRightWhile returns a slice with elements dropped from the end while the predicate returns true.
// Once the predicate returns false for an element, that element and all previous elements are kept.
// The function preserves the type of the input slice.
//
// Example:
//
//	nums := []int{1, 3, 5, 6, 8, 10}
//	result := DropRightWhile(nums, func(n int) bool {
//	    return n%2 == 0
//	}) // Returns []int{1, 3, 5}
func DropRightWhile[T any, Slice ~[]T](slice Slice, predicate func(item T) bool) Slice {
	if len(slice) == 0 {
		return make(Slice, 0)
	}

	idx := len(slice) - 1
	for ; idx >= 0 && predicate(slice[idx]); idx-- {
	}

	if idx < 0 {
		return make(Slice, 0)
	}

	ret := make(Slice, 0, idx+1)
	return append(ret, slice[:idx+1]...)
}

// DropByIndex returns a new slice with elements at the specified indexes removed.
// Negative indexes are supported and count from the end of the slice (-1 is the last element).
// Out of range indexes are ignored. Duplicate indexes are handled correctly.
//
// Example:
//
//	nums := []int{1, 2, 3, 4, 5}
//	result := DropByIndex(nums, 1, 3) // Returns []int{1, 3, 5}
//	result = DropByIndex(nums, -1, -3) // Returns []int{1, 2, 4}
func DropByIndex[T any](slice []T, indexes ...int) []T {
	if len(indexes) == 0 || len(slice) == 0 {
		return make([]T, 0)
	}

	// Convert negative indexes to positive
	validIndexes := make([]int, 0, len(indexes))
	for _, idx := range indexes {
		if idx < 0 {
			idx += len(slice)
		}
		if idx >= 0 && idx < len(slice) {
			validIndexes = append(validIndexes, idx)
		}
	}

	if len(validIndexes) == 0 {
		ret := make([]T, len(slice))
		copy(ret, slice)
		return ret
	}

	validIndexes = Uniq(validIndexes)
	sort.Ints(validIndexes)

	ret := make([]T, 0, len(slice)-len(validIndexes))
	lastIdx := 0
	for _, idx := range validIndexes {
		ret = append(ret, slice[lastIdx:idx]...)
		lastIdx = idx + 1
	}
	if lastIdx < len(slice) {
		ret = append(ret, slice[lastIdx:]...)
	}

	return ret
}

// Reject is the opposite of Filter. It returns a new slice with elements that satisfy the predicate removed.
// The predicate function takes both the element and its index as arguments.
// The function preserves the type of the input slice.
//
// Example:
//
//	nums := []int{1, 2, 3, 4, 5}
//	result := Reject(nums, func(n int, _ int) bool {
//	    return n%2 == 0
//	}) // Returns []int{1, 3, 5}
func Reject[T any, Slice ~[]T](slice Slice, predicate func(item T, index int) bool) Slice {
	if len(slice) == 0 {
		return make(Slice, 0)
	}

	ret := make(Slice, 0, len(slice))
	for idx, item := range slice {
		if !predicate(item, idx) {
			ret = append(ret, item)
		}
	}
	return ret
}

// RejectMap is the opposite of FilterMap.
// It returns a slice which obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be rejected (true) or kept (false).
//
// Example:
//
//	nums := []int{1, 2, 3, 4, 5}
//	result := RejectMap(nums, func(n int, _ int) (string, bool) {
//	    return fmt.Sprintf("num%d", n), n%2 == 0
//	}) // Returns []string{"num1", "num3", "num5"}
func RejectMap[T any, R any](slice []T, callback func(item T, index int) (R, bool)) []R {
	if len(slice) == 0 {
		return make([]R, 0)
	}

	ret := make([]R, 0, len(slice))
	for idx, item := range slice {
		if r, reject := callback(item, idx); !reject {
			ret = append(ret, r)
		}
	}
	return ret
}

// FilterReject is a combination of Filter and Reject. It returns two slices:
//   - kept: elements that satisfy the predicate
//   - rejected: elements that do not satisfy the predicate
//
// The function preserves the type of the input slice.
//
// Example:
//
//	nums := []int{1, 2, 3, 4, 5}
//	kept, rejected := FilterReject(nums, func(n int, _ int) bool {
//	    return n%2 == 0
//	}) // Returns []int{2, 4}, []int{1, 3, 5}
func FilterReject[T any, Slice ~[]T](slice Slice, predicate func(T, int) bool) (kept Slice, rejected Slice) {
	if len(slice) == 0 {
		return make(Slice, 0), make(Slice, 0)
	}

	kept = make(Slice, 0, len(slice)/2)
	rejected = make(Slice, 0, len(slice)/2)

	for idx, item := range slice {
		if predicate(item, idx) {
			kept = append(kept, item)
		} else {
			rejected = append(rejected, item)
		}
	}

	return kept, rejected
}

// Count returns the number of elements in the slice that are equal to the given value.
// Uses the == operator for comparison, so it works with any comparable type.
//
// Example:
//
//	nums := []int{1, 2, 2, 3, 2, 4, 5}
//	count := Count(nums, 2) // Returns 3
func Count[T comparable](slice []T, value T) int {
	if len(slice) == 0 {
		return 0
	}

	cnt := 0
	for _, item := range slice {
		if item == value {
			cnt++
		}
	}
	return cnt
}

// CountBy returns the number of elements in the slice that satisfy the given predicate.
//
// Example:
//
//	nums := []int{1, 2, 3, 4, 5, 6}
//	count := CountBy(nums, func(n int) bool {
//	    return n%2 == 0
//	}) // Returns 3
func CountBy[T any](slice []T, predicate func(item T) bool) int {
	if len(slice) == 0 {
		return 0
	}

	cnt := 0
	for _, item := range slice {
		if predicate(item) {
			cnt++
		}
	}
	return cnt
}

// CountValues returns a map that counts the number of occurrences of each unique value in the slice.
// The map's keys are the unique values from the slice, and the values are their counts.
//
// Example:
//
//	nums := []int{1, 2, 2, 3, 2, 4, 1}
//	counts := CountValues(nums) // Returns map[1:2 2:3 3:1 4:1]
func CountValues[T comparable](slice []T) map[T]int {
	if len(slice) == 0 {
		return make(map[T]int)
	}

	ret := make(map[T]int, len(slice))
	for _, item := range slice {
		ret[item]++
	}
	return ret
}

// CountValuesBy returns a map that counts the number of occurrences of each unique value in the slice,
// based on the given mapper function. The mapper function transforms each element into a key,
// and the map counts how many elements map to each key.
//
// Example:
//
//	nums := []int{1, 2, 3, 4, 5, 6}
//	counts := CountValuesBy(nums, func(n int) string {
//	    if n%2 == 0 { return "even" }
//	    return "odd"
//	}) // Returns map["even":3 "odd":3]
func CountValuesBy[T any, U comparable](slice []T, mapper func(item T) U) map[U]int {
	if len(slice) == 0 {
		return make(map[U]int)
	}

	ret := make(map[U]int, len(slice))
	for _, item := range slice {
		ret[mapper(item)]++
	}
	return ret
}

// SubSet returns a subset of the slice starting at offset with the given length.
// If offset is negative, it counts from the end of the slice.
// If length would extend beyond the slice bounds, it is truncated.
// If offset is beyond slice bounds (after handling negative offset), returns nil.
// The function preserves the type of the input slice.
//
// Example:
//
//	nums := []int{0, 1, 2, 3, 4, 5}
//	result := SubSet(nums, 2, 3) // Returns []int{2, 3, 4}
//	result = SubSet(nums, -2, 2) // Returns []int{4, 5}
func SubSet[T any, Slice ~[]T](slice Slice, offset int, length uint) Slice {
	if len(slice) == 0 {
		return make(Slice, 0)
	}

	size := len(slice)
	if offset < 0 {
		offset = size + offset
		if offset < 0 {
			offset = 0
		}
	}

	if offset >= size {
		return make(Slice, 0)
	}

	if length == 0 {
		return make(Slice, 0)
	}

	if length > uint(size-offset) { //nolint:gosec
		length = uint(size - offset) //nolint:gosec
	}

	ret := make(Slice, length)
	copy(ret, slice[offset:offset+int(length)]) //nolint:gosec
	return ret
}

// Slice returns a new slice containing elements from start (inclusive) to end (exclusive).
// It handles out-of-bounds indices gracefully:
//   - If start >= end, returns an empty slice
//   - If start < 0, uses 0 instead
//   - If start > len(slice), uses len(slice) instead
//   - If end < 0, uses 0 instead
//   - If end > len(slice), uses len(slice) instead
//
// Example:
//
//	nums := []int{0, 1, 2, 3, 4}
//	result := Slice(nums, 1, 3) // returns []int{1, 2}
//	empty := Slice(nums, 3, 1)  // returns []int{}
func Slice[T any, Slice ~[]T](slice Slice, start, end int) Slice {
	if len(slice) == 0 {
		return Slice{}
	}

	if start >= end {
		return Slice{}
	}

	size := len(slice)
	if start < 0 {
		start = 0
	} else if start > size {
		start = size
	}

	if end < 0 {
		end = 0
	} else if end > size {
		end = size
	}

	return slice[start:end]
}

// Replace returns a new slice with the first n occurrences of old replaced by new.
// If n < 0, all occurrences are replaced.
// The original slice remains unchanged.
//
// Example:
//
//	nums := []int{1, 2, 2, 3, 2, 4}
//	result := Replace(nums, 2, 5, 2) // returns []int{1, 5, 5, 3, 2, 4}
func Replace[T comparable, Slice ~[]T](slice Slice, old, new T, n int) Slice {
	if len(slice) == 0 || n == 0 {
		return slice
	}

	// Fast path: no replacements needed
	if n < 0 {
		hasOld := false
		for _, v := range slice {
			if v == old {
				hasOld = true
				break
			}
		}
		if !hasOld {
			return slice
		}
	}

	ret := make(Slice, len(slice))
	copy(ret, slice)

	if n < 0 {
		for i := range ret {
			if ret[i] == old {
				ret[i] = new
			}
		}
		return ret
	}

	count := n
	for i := range ret {
		if count == 0 {
			break
		}
		if ret[i] == old {
			ret[i] = new
			count--
		}
	}

	return ret
}

// ReplaceAll returns a new slice with all occurrences of old replaced by new.
// The original slice remains unchanged.
//
// Example:
//
//	nums := []int{1, 2, 2, 3, 2, 4}
//	result := ReplaceAll(nums, 2, 5) // returns []int{1, 5, 5, 3, 5, 4}
func ReplaceAll[T comparable, Slice ~[]T](slice Slice, old, new T) Slice {
	return Replace(slice, old, new, -1)
}

// Compact returns a new slice with all zero values(0, "", false) removed.
// The original slice remains unchanged.
//
// Example:
//
//	nums := []int{0, 1, 0, 2, 0, 3}
//	result := Compact(nums) // returns []int{1, 2, 3}
func Compact[T comparable, Slice ~[]T](slice Slice) Slice {
	if len(slice) == 0 {
		return slice
	}

	var zero T
	// Preallocate with capacity but zero length
	ret := make(Slice, 0, len(slice))
	for _, item := range slice {
		if item != zero {
			ret = append(ret, item)
		}
	}

	// If no non-zero elements found, return empty slice with zero capacity
	if len(ret) == 0 {
		return ret[:0]
	}
	return ret
}

// IsSorted returns true if the slice is sorted in ascending order, false otherwise.
// If the slice is empty or contains only one element, returns true.
//
// Example:
//
//	nums := []int{1, 2, 2, 3, 4, 5}
//	result := IsSorted(nums) // returns true
func IsSorted[T constraints.Ordered](slice []T) bool {
	if len(slice) <= 1 {
		return true
	}
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return false
		}
	}
	return true
}

// IsSortedByKey returns true if the slice is sorted in ascending order based on the iteratee function,
// false otherwise. If the slice is empty or contains only one element, returns true.
//
// Example:
//
//	nums := []int{1, 2, 2, 3, 4, 5}
//	result := IsSortedByKey(nums, func(n int) int { return n }) // returns true
func IsSortedByKey[T any, K constraints.Ordered](slice []T, iteratee func(item T) K) bool {
	if len(slice) <= 1 {
		return true
	}

	for i := 1; i < len(slice); i++ {
		if iteratee(slice[i-1]) > iteratee(slice[i]) {
			return false
		}
	}
	return true
}

// Splice returns a new slice with elements from the original slice inserted at the specified index.
// If idx is negative, it counts from the end of the slice.
//
// Example:
//
//	nums := []int{1, 2, 3, 4, 5}
//	result := Splice(nums, 2, 10, 20) // returns []int{1, 2, 10, 20, 3, 4, 5}
func Splice[T any, Slice ~[]T](slice Slice, idx int, elements ...T) Slice {
	sizeSlice := len(slice)
	sizeElems := len(elements)

	if sizeElems == 0 {
		return slice
	}

	// Normalize negative index
	if idx < 0 {
		idx = sizeSlice + idx
		if idx < 0 {
			idx = 0
		}
	}
	if idx > sizeSlice {
		idx = sizeSlice
	}

	// Preallocate result slice with exact capacity needed
	ret := make(Slice, 0, sizeSlice+sizeElems)
	ret = append(ret, slice[:idx]...)
	ret = append(ret, elements...)
	ret = append(ret, slice[idx:]...)

	return ret
}

// Equal checks if two slices are equal.
// For basic types (integers, floats, strings), it performs a direct comparison.
// For complex types (structs, pointers), it uses the == operator which may panic
// if the type does not support direct comparison.
//
// Example:
//
//	s1 := []int{1, 2, 3}
//	s2 := []int{1, 2, 3}
//	result := Equal(s1, s2) // returns true
func Equal[T comparable](v1, v2 []T) bool {
	// Quick length check
	if len(v1) != len(v2) {
		return false
	}

	// Empty  slices are equal
	if len(v1) == 0 {
		return true
	}

	// For small slices (< 100 elements), direct comparison is faster
	if len(v1) < 100 {
		for i := range v1 {
			if v1[i] != v2[i] {
				return false
			}
		}
		return true
	}

	// For larger slices, check boundaries first to potentially fail fast
	if v1[0] != v2[0] || v1[len(v1)-1] != v2[len(v1)-1] {
		return false
	}

	// Use a single loop for the remaining elements
	for i := 1; i < len(v1)-1; i++ {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true
}
