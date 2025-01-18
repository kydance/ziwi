package slices

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test filtering integers
	vInt := []int{1, 2, 3, 4, 5}
	vRetInt := Filter(vInt, func(item int, _ int) bool { return item%2 == 0 })
	as.Equal([]int{2, 4}, vRetInt)

	// Test filtering with empty result
	emptyFilter := Filter(vInt, func(item int, index int) bool { return item > 10 })
	as.Empty(emptyFilter)

	// Test filtering with all items matching
	allMatchFilter := Filter(vInt, func(item int, index int) bool { return item < 10 })
	as.Equal(vInt, allMatchFilter)

	// Test filtering strings
	strings := []string{"", "apple", "banana", "cherry"}
	filteredStrings := Filter(strings, func(item string, _ int) bool { return len(item) > 5 })
	as.Equal([]string{"banana", "cherry"}, filteredStrings)

	type StringSlice []string
	ss := StringSlice{"", "apple", "banana", "cherry"}
	nonEmpty := Filter(ss, func(item string, _ int) bool { return len(item) > 0 })
	as.IsType(nonEmpty, ss, "type preserved")
}

func TestMap(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test mapping integers to strings
	ints := []int{-1, 0, 1, 2, 3}
	stringMap := Map(ints, func(item int, index int) string { return fmt.Sprintf("Number %d", item) })
	expectedStrings := []string{"Number -1", "Number 0", "Number 1", "Number 2", "Number 3"}
	as.Equal(expectedStrings, stringMap)

	// Test mapping strings to integers
	strings := []string{"one", "two", "three"}
	intMap := Map(strings, func(item string, index int) int { return len(item) })
	expectedInts := []int{3, 3, 5}
	as.Equal(expectedInts, intMap)

	// Test mapping with empty input
	emptySlice := []int{}
	emptyResult := Map(emptySlice, func(item int, index int) int { return item * 2 })
	as.Empty(emptyResult)

	// Test type preservation
	type MyInt int
	myInts := []MyInt{10, 20, 30}
	mappedMyInts := Map(myInts, func(item MyInt, index int) MyInt { return item + 10 })
	as.IsType(mappedMyInts, myInts, "type preserved")
}

func TestFilterMap(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test filtering and mapping integers to strings
	ints := []int{1, 2, 3, 4, 5}
	filteredMappedInts := FilterMap(ints, func(item int, index int) (string, bool) {
		return fmt.Sprintf("Number %d", item), item%2 == 0
	})
	expectedStrings := []string{"Number 2", "Number 4"}
	as.Equal(expectedStrings, filteredMappedInts)

	// Test filtering and mapping with no matching items
	noneMatching := FilterMap(ints, func(item int, index int) (string, bool) {
		return fmt.Sprintf("Number %d", item), item > 5
	})
	as.Empty(noneMatching)

	// Test filtering and mapping with all items matching
	allMatching := FilterMap(ints, func(item int, index int) (string, bool) {
		return fmt.Sprintf("Number %d", item), item < 6
	})
	expectedAllMatching := []string{"Number 1", "Number 2", "Number 3", "Number 4", "Number 5"}
	as.Equal(expectedAllMatching, allMatching)

	// Test filtering and mapping stringSlice to integers
	stringSlice := []string{"zero", "one", "two", "three"}
	filteredMappedStrings := FilterMap(stringSlice, func(item string, index int) (int, bool) {
		return len(item), len(item) > 3
	})
	expectedInts := []int{4, 5}
	as.Equal(expectedInts, filteredMappedStrings)

	// Test type preservation
	type MyInt int
	myInts := []MyInt{10, 20, 30}
	mappedMyInts := FilterMap(myInts, func(item MyInt, index int) (MyInt, bool) {
		return item + 10, item > 15
	})
	expectedMappedMyInts := []MyInt{30, 40}
	as.Equal(expectedMappedMyInts, mappedMyInts)
	as.IsType(mappedMyInts, myInts, "type preserved")

	//
	r1 := FilterMap([]int64{1, 2, 3, 4}, func(x int64, _ int) (string, bool) {
		if x%2 == 0 {
			return strconv.FormatInt(x, 10), true
		}
		return "", false
	})
	r2 := FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", true
		}
		return "", false
	})

	as.Equal(len(r1), 2)
	as.Equal(len(r2), 2)
	as.Equal(r1, []string{"2", "4"})
	as.Equal(r2, []string{"xpu", "xpu"})
}

func TestFlatMap(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test flattening nested integers
	numbers := []int{1, 2, 3}
	doubled := FlatMap(numbers, func(item int, index int) []int {
		return []int{item, item}
	})
	as.Equal([]int{1, 1, 2, 2, 3, 3}, doubled)

	// Test flattening strings to characters
	words := []string{"go", "rust"}
	chars := FlatMap(words, func(item string, index int) []string {
		return strings.Split(item, "")
	})
	as.Equal([]string{"g", "o", "r", "u", "s", "t"}, chars)

	// Test with empty input slice
	empty := []int{}
	emptyResult := FlatMap(empty, func(item int, index int) []string {
		return []string{fmt.Sprint(item)}
	})
	as.Empty(emptyResult)

	// Test with empty result slices
	noResults := FlatMap(numbers, func(item int, index int) []string {
		return []string{}
	})
	as.Empty(noResults)

	// Test type preservation
	type MyInt int
	myInts := []MyInt{10, 20}
	mappedMyInts := FlatMap(myInts, func(item MyInt, index int) []MyInt {
		return []MyInt{item, item + 1}
	})
	expectedMyInts := []MyInt{10, 11, 20, 21}
	as.Equal(expectedMyInts, mappedMyInts)
	as.IsType(mappedMyInts, myInts, "type preserved")
}

func TestReduce(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test sum reduction
	numbers := []int{1, 2, 3, 4, 5}
	sum := Reduce(numbers, func(agg int, item int, index int) int {
		return agg + item
	}, 0)
	as.Equal(15, sum)

	// Test string concatenation
	words := []string{"Hello", " ", "World", "!"}
	sentence := Reduce(words, func(agg string, item string, index int) string {
		return agg + item
	}, "")
	as.Equal("Hello World!", sentence)

	// Test slice building
	doubled := Reduce(numbers, func(agg []int, item int, index int) []int {
		return append(agg, item*2)
	}, make([]int, 0, len(numbers)))
	as.Equal([]int{2, 4, 6, 8, 10}, doubled)

	// Test with empty slice
	empty := []int{}
	emptySum := Reduce(empty, func(agg int, item int, index int) int {
		return agg + item
	}, 0)
	as.Equal(0, emptySum)

	// Test type preservation
	type MyInt int
	myInts := []MyInt{10, 20, 30}
	mySum := Reduce(myInts, func(agg MyInt, item MyInt, index int) MyInt {
		return agg + item
	}, MyInt(0))
	as.Equal(MyInt(60), mySum)
}

func TestReduceRight(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test right-to-left string concatenation
	words := []string{"!", "World", " ", "Hello"}
	sentence := ReduceRight(words, func(agg string, item string, index int) string {
		return agg + item
	}, "")
	as.Equal("Hello World!", sentence)

	// Test right-to-left array operations
	numbers := []int{1, 2, 3, 4}
	// Demonstrating order matters with division
	divisionResult := ReduceRight(numbers, func(agg float64, item int, index int) float64 {
		if agg == 0 {
			return float64(item)
		}
		return float64(item) / agg
	}, float64(0))
	// 4 -> 3/4 -> 2/(3/4) -> 1/(2/(3/4)) ≈ 0.375
	as.InDelta(0.375, divisionResult, 0.0001)

	// Test with empty slice
	empty := []int{}
	emptyResult := ReduceRight(empty, func(agg int, item int, index int) int {
		return agg + item
	}, 0)
	as.Equal(0, emptyResult)

	// Test type preservation
	type MyInt int
	myInts := []MyInt{10, 20, 30}
	mySum := ReduceRight(myInts, func(agg MyInt, item MyInt, index int) MyInt {
		return agg + item
	}, MyInt(0))
	as.Equal(MyInt(60), mySum)
}

func TestForEach(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic iteration
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	ForEach(numbers, func(item int, index int) {
		sum += item
		as.Equal(numbers[index], item)
	})
	as.Equal(15, sum)

	// Test with empty slice
	empty := []int{}
	emptyCount := 0
	ForEach(empty, func(item int, index int) {
		emptyCount++
	})
	as.Equal(0, emptyCount)

	// Test with custom type
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
	}
	totalAge := 0
	ForEach(people, func(p Person, index int) {
		totalAge += p.Age
	})
	as.Equal(55, totalAge)
}

func TestForEachWhile(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test early termination
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	completed := ForEachWhile(numbers, func(item int, index int) bool {
		if item > 3 {
			return false
		}
		sum += item
		return true
	})
	as.Equal(6, sum) // 1 + 2 + 3
	as.Equal(false, completed)

	// Test complete iteration
	allSum := 0
	completed = ForEachWhile(numbers, func(item int, index int) bool {
		allSum += item
		return true
	})
	as.Equal(15, allSum) // 1 + 2 + 3 + 4 + 5
	as.Equal(true, completed)

	// Test with empty slice
	empty := []int{}
	emptyCount := 0
	completed = ForEachWhile(empty, func(item int, index int) bool {
		emptyCount++
		return true
	})
	as.Equal(0, emptyCount)
	as.Equal(true, completed)

	// Test index accuracy
	indices := make([]int, 0)
	ForEachWhile(numbers, func(item int, index int) bool {
		indices = append(indices, index)
		return index < 2
	})
	as.Equal([]int{0, 1, 2}, indices)
}

func TestTimes(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test generating sequence
	numbers := Times(3, func(idx int) int {
		return idx * 2
	})
	as.Equal([]int{0, 2, 4}, numbers)

	// Test with strings
	labels := Times(2, func(idx int) string {
		return fmt.Sprintf("Item %d", idx)
	})
	as.Equal([]string{"Item 0", "Item 1"}, labels)

	// Test with custom type
	type Point struct {
		X, Y int
	}
	points := Times(2, func(idx int) Point {
		return Point{X: idx, Y: idx * 2}
	})
	as.Equal([]Point{{X: 0, Y: 0}, {X: 1, Y: 2}}, points)

	// Test with zero count
	empty := Times(0, func(idx int) int {
		return idx
	})
	as.Empty(empty)

	// Test with negative count
	negative := Times(-1, func(idx int) int {
		return idx
	})
	as.Empty(negative)
}

func TestUniq(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test with integers
	numbers := []int{1, 2, 2, 3, 3, 3, 4}
	uniqueNums := Uniq(numbers)
	as.Equal([]int{1, 2, 3, 4}, uniqueNums)

	// Test with strings
	words := []string{"a", "b", "b", "c", "c", "c"}
	uniqueWords := Uniq(words)
	as.Equal([]string{"a", "b", "c"}, uniqueWords)

	// Test with empty slice
	empty := []int{}
	uniqueEmpty := Uniq(empty)
	as.Empty(uniqueEmpty)

	// Test with all duplicates
	allDups := []int{1, 1, 1, 1}
	uniqueDups := Uniq(allDups)
	as.Equal([]int{1}, uniqueDups)

	// Test with custom type
	type MyInt int
	myInts := []MyInt{1, 1, 2, 2, 3}
	uniqueMyInts := Uniq(myInts)
	as.Equal([]MyInt{1, 2, 3}, uniqueMyInts)
	as.IsType(uniqueMyInts, myInts, "type preserved")

	// Test order preservation
	ordered := []int{3, 1, 2, 1, 3}
	uniqueOrdered := Uniq(ordered)
	as.Equal([]int{3, 1, 2}, uniqueOrdered)
}

func TestUniqBy(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test with custom comparison function
	numbers := []int{1, -1, 2, -2, 3, -3}
	absUnique := UniqBy(numbers, func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	})
	as.Equal([]int{1, 2, 3}, absUnique)

	// Test with struct field
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 30},
	}
	uniqueByAge := UniqBy(people, func(p Person) int {
		return p.Age
	})
	as.Equal([]Person{
		{Name: "Alice", Age: 25},
		{Name: "Charlie", Age: 30},
	}, uniqueByAge)

	// Test with empty slice
	empty := []int{}
	emptyResult := UniqBy(empty, func(n int) int { return n })
	as.Empty(emptyResult)

	// Test with all unique items
	allUnique := []string{"a", "b", "c"}
	result := UniqBy(allUnique, func(s string) string { return s })
	as.Equal([]string{"a", "b", "c"}, result)

	// Test with all duplicates
	allDups := []int{1, 1, 1, 1}
	uniqueDups := UniqBy(allDups, func(n int) int { return n })
	as.Equal([]int{1}, uniqueDups)
}

func TestGroupBy(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test grouping by length
	words := []string{"one", "two", "three", "four", "five"}
	byLength := GroupBy(words, func(s string) int {
		return len(s)
	})
	as.Equal([]string{"one", "two"}, byLength[3])
	as.Equal([]string{"four", "five"}, byLength[4])
	as.Equal([]string{"three"}, byLength[5])

	// Test grouping by first letter
	names := []string{"Alice", "Bob", "Amy", "Charlie"}
	byFirstLetter := GroupBy(names, func(s string) string {
		return s[:1]
	})
	as.Equal([]string{"Alice", "Amy"}, byFirstLetter["A"])
	as.Equal([]string{"Bob"}, byFirstLetter["B"])
	as.Equal([]string{"Charlie"}, byFirstLetter["C"])

	// Test with custom type
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 25},
	}
	byAge := GroupBy(people, func(p Person) int {
		return p.Age
	})
	as.Equal([]Person{
		{Name: "Alice", Age: 25},
		{Name: "Charlie", Age: 25},
	}, byAge[25])
	as.Equal([]Person{
		{Name: "Bob", Age: 30},
	}, byAge[30])

	// Test with empty slice
	empty := []int{}
	emptyResult := GroupBy(empty, func(n int) int { return n })
	as.Empty(emptyResult)

	// Test with single key
	sameAge := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 25},
	}
	sameAgeGroups := GroupBy(sameAge, func(p Person) int {
		return p.Age
	})
	as.Equal(1, len(sameAgeGroups))
	as.Equal(2, len(sameAgeGroups[25]))
}

func TestChunk(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test even division
	numbers := []int{1, 2, 3, 4}
	chunks := Chunk(numbers, 2)
	as.Equal([][]int{{1, 2}, {3, 4}}, chunks)

	// Test with remainder
	numbers = []int{1, 2, 3, 4, 5}
	chunks = Chunk(numbers, 2)
	as.Equal([][]int{{1, 2}, {3, 4}, {5}}, chunks)

	// Test with chunk size 1
	chunks = Chunk(numbers, 1)
	as.Equal([][]int{{1}, {2}, {3}, {4}, {5}}, chunks)

	// Test with chunk size equal to length
	chunks = Chunk(numbers, 5)
	as.Equal([][]int{{1, 2, 3, 4, 5}}, chunks)

	// Test with chunk size greater than length
	chunks = Chunk(numbers, 10)
	as.Equal([][]int{{1, 2, 3, 4, 5}}, chunks)

	// Test with empty slice
	empty := []int{}
	emptyChunks := Chunk(empty, 2)
	as.Empty(emptyChunks)

	// Test with custom type
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}
	personChunks := Chunk(people, 2)
	as.Equal([][]Person{
		{{Name: "Alice", Age: 25}, {Name: "Bob", Age: 30}},
		{{Name: "Charlie", Age: 35}},
	}, personChunks)

	// Test that chunks are views into original slice
	original := []int{1, 2, 3, 4}
	chunks = Chunk(original, 2)
	original[0] = 100
	as.Equal(100, chunks[0][0], "chunks should be views into original slice")

	// Test panic with zero size
	as.Panics(func() {
		Chunk(numbers, 0)
	})

	// Test panic with negative size
	as.Panics(func() {
		Chunk(numbers, -1)
	})

	// Verify capacity of chunks
	chunks = Chunk([]int{1, 2, 3, 4, 5}, 2)
	as.Equal(2, cap(chunks[0]), "first chunk should have capacity equal to chunk size")
	as.Equal(1, cap(chunks[2]), "last chunk should have capacity equal to remaining elements")
}

func TestPartitionBy(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test partitioning by parity
	numbers := []int{1, 2, 3, 4, 5, 6}
	groups := PartitionBy(numbers, func(n int) bool {
		return n%2 == 0
	})
	as.Equal([][]int{{1, 3, 5}, {2, 4, 6}}, groups)

	// Test partitioning strings by length
	words := []string{"one", "two", "three", "four", "five"}
	lengthGroups := PartitionBy(words, func(s string) int {
		return len(s)
	})
	as.Equal([][]string{
		{"one", "two"},   // length 3
		{"three"},        // length 5
		{"four", "five"}, // length 4
	}, lengthGroups, "groups should appear in order of first occurrence of each length")

	// Test with custom type
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 25},
		{Name: "David", Age: 30},
	}
	ageGroups := PartitionBy(people, func(p Person) int {
		return p.Age
	})
	as.Equal([][]Person{
		{{Name: "Alice", Age: 25}, {Name: "Charlie", Age: 25}},
		{{Name: "Bob", Age: 30}, {Name: "David", Age: 30}},
	}, ageGroups)

	// Test with empty slice
	empty := []int{}
	emptyGroups := PartitionBy(empty, func(n int) bool {
		return n%2 == 0
	})
	as.Empty(emptyGroups)

	// Test with single element
	single := []int{1}
	singleGroups := PartitionBy(single, func(n int) bool {
		return n%2 == 0
	})
	as.Equal([][]int{{1}}, singleGroups)

	// Test with all same key
	allSame := []int{2, 4, 6, 8}
	sameGroups := PartitionBy(allSame, func(n int) bool {
		return n%2 == 0
	})
	as.Equal([][]int{{2, 4, 6, 8}}, sameGroups)

	// Test with all different keys
	allDiff := []int{1, 2, 3, 4}
	diffGroups := PartitionBy(allDiff, func(n int) int {
		return n
	})
	as.Equal([][]int{{1}, {2}, {3}, {4}}, diffGroups)

	// Test order preservation within groups
	ordered := []int{5, 1, 3, 2, 4, 6}
	orderedGroups := PartitionBy(ordered, func(n int) bool {
		return n%2 == 0
	})
	as.Equal([][]int{{5, 1, 3}, {2, 4, 6}}, orderedGroups)

	// Verify capacity optimization
	largeSlice := make([]int, 100)
	for i := range largeSlice {
		largeSlice[i] = i
	}
	largeGroups := PartitionBy(largeSlice, func(n int) int {
		return n % 3
	})
	as.Equal(3, len(largeGroups)) // Should have 3 groups (remainders 0, 1, 2)
	as.True(cap(largeGroups[0]) >= len(largeGroups[0]), "group capacity should be sufficient")
}

func TestFlatten(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic flattening
	numbers := [][]int{{1, 2}, {3, 4}, {5, 6}}
	flattened := Flatten(numbers)
	as.Equal([]int{1, 2, 3, 4, 5, 6}, flattened)

	// Test with empty slices in between
	mixed := [][]int{{1}, {}, {2, 3}, {}, {4}}
	flattened = Flatten(mixed)
	as.Equal([]int{1, 2, 3, 4}, flattened)

	// Test with empty input
	empty := [][]int{}
	flattened = Flatten(empty)
	as.Empty(flattened)

	// Test with all empty slices
	allEmpty := [][]int{{}, {}, {}}
	flattened = Flatten(allEmpty)
	as.Empty(flattened)

	// Test with single slice
	single := [][]int{{1, 2, 3}}
	flattened = Flatten(single)
	as.Equal([]int{1, 2, 3}, flattened)

	// Test with single empty slice
	singleEmpty := [][]int{{}}
	flattened = Flatten(singleEmpty)
	as.Empty(flattened)

	// Test with custom type
	type Person struct {
		Name string
		Age  int
	}
	people := [][]Person{
		{{Name: "Alice", Age: 25}},
		{{Name: "Bob", Age: 30}, {Name: "Charlie", Age: 35}},
	}
	flatPeople := Flatten(people)
	as.Equal([]Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}, flatPeople)

	// Test with strings
	words := [][]string{{"hello"}, {"world", "!"}}
	flatWords := Flatten(words)
	as.Equal([]string{"hello", "world", "!"}, flatWords)

	// Test order preservation
	ordered := [][]int{{3, 2, 1}, {6, 5, 4}, {9, 8, 7}}
	flattened = Flatten(ordered)
	as.Equal([]int{3, 2, 1, 6, 5, 4, 9, 8, 7}, flattened)

	// Test capacity optimization
	largeSlice := make([][]int, 100)
	for i := range largeSlice {
		largeSlice[i] = []int{i}
	}
	flattened = Flatten(largeSlice)
	as.Equal(100, len(flattened))
	as.Equal(100, cap(flattened), "capacity should match exactly with total length")

	// Test that flattened slice is independent
	original := [][]int{{1, 2}, {3, 4}}
	flattened = Flatten(original)
	original[0][0] = 100
	as.Equal(1, flattened[0], "flattened slice should be independent of original")
}

func TestInterleave(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic interleaving
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := []int{7, 8, 9}
	result := Interleave[int, []int](a, b, c)
	as.Equal([]int{1, 4, 7, 2, 5, 8, 3, 6, 9}, result)

	// Test with different lengths
	a = []int{1, 2, 3}
	b = []int{4, 5}
	c = []int{6}
	result = Interleave[int, []int](a, b, c)
	as.Equal([]int{1, 4, 6, 2, 5, 3}, result)

	// Test with empty slices
	empty := []int{}
	result = Interleave[int, []int]([]int{1}, empty, []int{2, 3})
	as.Equal([]int{1, 2, 3}, result)

	// Test with all empty slices
	result = Interleave[int, []int](empty, empty, empty)
	as.Empty(result)

	// Test with no slices
	result = Interleave[int, []int]()
	as.Empty(result)

	// Test with single slice
	single := []int{1, 2, 3}
	result = Interleave[int, []int](single)
	as.Equal(single, result)

	// Test with strings
	words1 := []string{"a", "b", "c"}
	words2 := []string{"d", "e", "f"}
	strResult := Interleave[string, []string](words1, words2)
	as.Equal([]string{"a", "d", "b", "e", "c", "f"}, strResult)

	// Test with custom type
	type Person struct {
		Name string
		Age  int
	}
	people1 := []Person{{Name: "Alice", Age: 25}, {Name: "Bob", Age: 30}}
	people2 := []Person{{Name: "Charlie", Age: 35}}
	peopleResult := Interleave[Person, []Person](people1, people2)
	as.Equal([]Person{
		{Name: "Alice", Age: 25},
		{Name: "Charlie", Age: 35},
		{Name: "Bob", Age: 30},
	}, peopleResult)

	// Test with many empty slices and one non-empty
	manyEmpty := make([][]int, 10)
	for i := range manyEmpty {
		manyEmpty[i] = []int{}
	}
	manyEmpty[5] = []int{1, 2, 3}
	result = Interleave[int, []int](manyEmpty...)
	as.Equal([]int{1, 2, 3}, result)

	// Test capacity optimization
	result = Interleave[int, []int]([]int{1, 2}, []int{3, 4})
	as.Equal(4, cap(result), "capacity should match total length")

	// Test order preservation within each slice
	ordered1 := []int{3, 2, 1}
	ordered2 := []int{6, 5, 4}
	result = Interleave[int, []int](ordered1, ordered2)
	as.Equal([]int{3, 6, 2, 5, 1, 4}, result)

	// Test with nil slices
	var nilSlice []int
	result = Interleave[int, []int](nilSlice, []int{1, 2})
	as.Equal([]int{1, 2}, result)
}

func TestShuffle(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test shuffling integers
	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffled := make([]int, len(original))
	copy(shuffled, original)

	Shuffle(shuffled)

	// Verify length remains the same
	as.Equal(len(original), len(shuffled))

	// Verify all elements are present (just in different order)
	as.ElementsMatch(original, shuffled)

	// Verify the slice was actually shuffled (this has a tiny chance of false failure)
	different := false
	for i := range original {
		if original[i] != shuffled[i] {
			different = true
			break
		}
	}
	as.True(different, "Slice should be shuffled to a different order")

	// Test empty slice
	empty := []int{}
	Shuffle(empty)
	as.Empty(empty)

	// Test single element
	single := []int{1}
	Shuffle(single)
	as.Equal([]int{1}, single)

	// Test type preservation
	type MyInt []int
	myInts := MyInt{1, 2, 3, 4, 5}
	shuffledMyInts := Shuffle(myInts)
	as.IsType(myInts, shuffledMyInts, "type should be preserved")
}

func TestReverse(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test reversing integers
	ints := []int{1, 2, 3, 4, 5}
	reversed := Reverse(append([]int{}, ints...))
	as.Equal([]int{5, 4, 3, 2, 1}, reversed)

	// Test empty slice
	empty := []int{}
	Reverse(empty)
	as.Empty(empty)

	// Test single element
	single := []int{1}
	Reverse(single)
	as.Equal([]int{1}, single)

	// Test even number of elements
	even := []int{1, 2, 3, 4}
	Reverse(even)
	as.Equal([]int{4, 3, 2, 1}, even)

	// Test odd number of elements
	odd := []int{1, 2, 3}
	Reverse(odd)
	as.Equal([]int{3, 2, 1}, odd)

	// Test type preservation
	type MyInt []int
	myInts := MyInt{1, 2, 3, 4, 5}
	reversedMyInts := Reverse(myInts)
	as.IsType(myInts, reversedMyInts, "type should be preserved")
}

type testCloneable struct{ value int }

func (t testCloneable) Clone() testCloneable { return testCloneable{value: t.value} }

func TestFill(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test with empty slice
	empty := []testCloneable{}
	filled := Fill(empty, testCloneable{value: 42})
	as.Empty(filled)

	// Test with non-empty slice
	slice := make([]testCloneable, 3)
	filled = Fill(slice, testCloneable{value: 42})
	as.Len(filled, 3)
	for _, item := range filled {
		as.Equal(42, item.value)
	}

	// Test that clones are independent
	filled[0].value = 100
	as.Equal(42, filled[1].value, "modifying one element should not affect others")
}

func TestRepeat(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test with count <= 0
	as.Empty(Repeat(0, testCloneable{value: 42}))
	as.Empty(Repeat(-1, testCloneable{value: 42}))

	// Test with positive count
	repeated := Repeat(3, testCloneable{value: 42})
	as.Len(repeated, 3)
	for _, item := range repeated {
		as.Equal(42, item.value)
	}

	// Test that clones are independent
	repeated[0].value = 100
	as.Equal(42, repeated[1].value, "modifying one element should not affect others")
}

func TestRepeatBy(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test with count <= 0
	as.Empty(RepeatBy(0, func(i int) int { return i * 2 }))
	as.Empty(RepeatBy(-1, func(i int) int { return i * 2 }))

	// Test with positive count
	repeated := RepeatBy(5, func(i int) int { return i * 2 })
	as.Equal([]int{0, 2, 4, 6, 8}, repeated)

	// Test with string generator
	strRepeated := RepeatBy(3, func(i int) string { return fmt.Sprintf("item-%d", i) })
	as.Equal([]string{"item-0", "item-1", "item-2"}, strRepeated)
}

func TestKeyBy(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test with simple int slice
	type User struct {
		ID   int
		Name string
	}
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	result := KeyBy(users, func(u User) int { return u.ID })
	as.Equal(map[int]User{
		1: {ID: 1, Name: "Alice"},
		2: {ID: 2, Name: "Bob"},
		3: {ID: 3, Name: "Charlie"},
	}, result)

	// Test with string keys
	strResult := KeyBy(users, func(u User) string { return u.Name })
	as.Equal(map[string]User{
		"Alice":   {ID: 1, Name: "Alice"},
		"Bob":     {ID: 2, Name: "Bob"},
		"Charlie": {ID: 3, Name: "Charlie"},
	}, strResult)

	// Test with empty slice
	emptyResult := KeyBy([]User{}, func(u User) int { return u.ID })
	as.Empty(emptyResult)
}

func TestAssociate(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test with basic type transformation
	numbers := []int{1, 2, 3}
	result := Associate(numbers, func(n int) (string, bool) {
		return fmt.Sprintf("num%d", n), n%2 == 0
	})
	as.Equal(map[string]bool{
		"num1": false,
		"num2": true,
		"num3": false,
	}, result)

	// Test with struct transformation
	type User struct {
		ID   int
		Name string
	}
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
	userResult := Associate(users, func(u User) (int, string) {
		return u.ID, strings.ToUpper(u.Name)
	})
	as.Equal(map[int]string{
		1: "ALICE",
		2: "BOB",
	}, userResult)

	// Test with empty slice
	emptyResult := Associate([]int{}, func(n int) (string, int) {
		return fmt.Sprintf("%d", n), n * 2
	})
	as.Empty(emptyResult)
}

func TestSliceToMap(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic transformation
	numbers := []int{1, 2, 3}
	result := SliceToMap(numbers, func(n int) (string, int) {
		return fmt.Sprintf("key%d", n), n * 2
	})
	as.Equal(map[string]int{
		"key1": 2,
		"key2": 4,
		"key3": 6,
	}, result)

	// Test with empty slice
	emptyResult := SliceToMap([]int{}, func(n int) (string, int) {
		return fmt.Sprintf("%d", n), n
	})
	as.Empty(emptyResult)
}

func TestDrop(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic drop
	nums := []int{1, 2, 3, 4, 5}
	result := Drop(nums, 2)
	as.Equal([]int{3, 4, 5}, result)

	// Test drop all
	allDropped := Drop(nums, len(nums))
	as.Empty(allDropped)

	// Test drop more than length
	overDropped := Drop(nums, len(nums)+1)
	as.Empty(overDropped)

	// Test drop zero
	noDrop := Drop(nums, 0)
	as.Equal(nums, noDrop)

	// Test with custom slice type
	type IntSlice []int
	customSlice := IntSlice{1, 2, 3}
	customResult := Drop(customSlice, 1)
	as.Equal(IntSlice{2, 3}, customResult)
}

func TestDropRight(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic drop right
	nums := []int{1, 2, 3, 4, 5}
	result := DropRight(nums, 2)
	as.Equal([]int{1, 2, 3}, result)

	// Test drop all
	allDropped := DropRight(nums, len(nums))
	as.Empty(allDropped)

	// Test drop more than length
	overDropped := DropRight(nums, len(nums)+1)
	as.Empty(overDropped)

	// Test drop zero
	noDrop := DropRight(nums, 0)
	as.Equal(nums, noDrop)

	// Test with custom slice type
	type IntSlice []int
	customSlice := IntSlice{1, 2, 3}
	customResult := DropRight(customSlice, 1)
	as.Equal(IntSlice{1, 2}, customResult)
}

func TestDropWhile(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic drop while
	nums := []int{2, 4, 6, 7, 8, 10}
	result := DropWhile(nums, func(n int) bool {
		return n%2 == 0
	})
	as.Equal([]int{7, 8, 10}, result)

	// Test drop all
	allEven := []int{2, 4, 6, 8}
	allDropped := DropWhile(allEven, func(n int) bool {
		return n%2 == 0
	})
	as.Empty(allDropped)

	// Test drop none
	noDrops := DropWhile(nums, func(n int) bool {
		return n > 100
	})
	as.Equal(nums, noDrops)

	// Test with empty slice
	emptyResult := DropWhile([]int{}, func(n int) bool {
		return true
	})
	as.Empty(emptyResult)
}

func TestDropRightWhile(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic drop right while
	nums := []int{1, 3, 5, 6, 8, 10}
	result := DropRightWhile(nums, func(n int) bool {
		return n%2 == 0
	})
	as.Equal([]int{1, 3, 5}, result)

	// Test drop all
	allEven := []int{2, 4, 6, 8}
	allDropped := DropRightWhile(allEven, func(n int) bool {
		return n%2 == 0
	})
	as.Empty(allDropped)

	// Test drop none
	noDrops := DropRightWhile(nums, func(n int) bool {
		return n > 100
	})
	as.Equal(nums, noDrops)

	// Test with empty slice
	emptyResult := DropRightWhile([]int{}, func(n int) bool {
		return true
	})
	as.Empty(emptyResult)
}

func TestDropByIndex(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic index drops
	nums := []int{1, 2, 3, 4, 5}
	result := DropByIndex(nums, 1, 3)
	as.Equal([]int{1, 3, 5}, result)

	// Test negative indexes
	negResult := DropByIndex(nums, -1, -3)
	as.Equal([]int{1, 2, 4}, negResult)

	// Test out of range indexes
	outRange := DropByIndex(nums, 10, -10)
	as.Equal(nums, outRange)

	// Test duplicate indexes
	dupResult := DropByIndex(nums, 1, 1, 1)
	as.Equal([]int{1, 3, 4, 5}, dupResult)

	// Test empty indexes
	noIndexes := DropByIndex(nums)
	as.Empty(noIndexes)

	// Test empty slice
	emptySlice := DropByIndex([]int{}, 1, 2)
	as.Empty(emptySlice)

	// Test all indexes
	allIndexes := DropByIndex(nums, 0, 1, 2, 3, 4)
	as.Empty(allIndexes)
}

func TestReject(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic rejection
	nums := []int{1, 2, 3, 4, 5}
	result := Reject(nums, func(n int, _ int) bool {
		return n%2 == 0
	})
	as.Equal([]int{1, 3, 5}, result)

	// Test with index
	indexBased := Reject(nums, func(n int, idx int) bool {
		return idx < 2
	})
	as.Equal([]int{3, 4, 5}, indexBased)

	// Test with empty slice
	emptyResult := Reject([]int{}, func(n int, _ int) bool {
		return true
	})
	as.Empty(emptyResult)

	// Test with custom slice type
	type IntSlice []int
	customSlice := IntSlice{1, 2, 3}
	customResult := Reject(customSlice, func(n int, _ int) bool {
		return n > 2
	})
	as.Equal(IntSlice{1, 2}, customResult)
}

func TestRejectMap(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic reject-map
	nums := []int{1, 2, 3, 4, 5}
	result := RejectMap(nums, func(n int, _ int) (string, bool) {
		return fmt.Sprintf("num%d", n), n%2 == 0
	})
	as.Equal([]string{"num1", "num3", "num5"}, result)

	// Test with type transformation
	typeChange := RejectMap(nums, func(n int, idx int) (float64, bool) {
		return float64(n) * 1.5, idx < 2
	})
	as.Equal([]float64{4.5, 6.0, 7.5}, typeChange)

	// Test with empty slice
	emptyResult := RejectMap([]int{}, func(n int, _ int) (string, bool) {
		return "", true
	})
	as.Empty(emptyResult)
}

func TestFilterReject(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test empty slice
	emptyNum, emptyRejected := FilterReject([]int{}, func(n int, _ int) bool { return true })
	as.Empty(emptyNum)
	as.Empty(emptyRejected)

	// Test basic filter-reject
	nums := []int{1, 2, 3, 4, 5}
	kept, rejected := FilterReject(nums, func(n int, _ int) bool {
		return n%2 == 0
	})
	as.Equal([]int{2, 4}, kept)
	as.Equal([]int{1, 3, 5}, rejected)

	// Test with all kept
	allKept, noRejected := FilterReject(nums, func(_ int, _ int) bool {
		return true
	})
	as.Equal(nums, allKept)
	as.Empty(noRejected)

	// Test with all rejected
	noKept, allRejected := FilterReject(nums, func(_ int, _ int) bool {
		return false
	})
	as.Empty(noKept)
	as.Equal(nums, allRejected)

	// Test with custom slice type
	type IntSlice []int
	customSlice := IntSlice{1, 2, 3}
	customKept, customRejected := FilterReject(customSlice, func(n int, _ int) bool {
		return n > 1
	})
	as.Equal(IntSlice{2, 3}, customKept)
	as.Equal(IntSlice{1}, customRejected)
}

func TestCount(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic counting
	nums := []int{1, 2, 2, 3, 2, 4, 5}
	count := Count(nums, 2)
	as.Equal(3, count)

	// Test with no matches
	noMatches := Count(nums, 6)
	as.Equal(0, noMatches)

	// Test with empty slice
	emptyCount := Count([]int{}, 1)
	as.Equal(0, emptyCount)

	// Test with strings
	words := []string{"apple", "banana", "apple", "cherry"}
	wordCount := Count(words, "apple")
	as.Equal(2, wordCount)
}

func TestCountBy(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic counting by predicate
	nums := []int{1, 2, 3, 4, 5, 6}
	evenCount := CountBy(nums, func(n int) bool {
		return n%2 == 0
	})
	as.Equal(3, evenCount)

	// Test with no matches
	noMatches := CountBy(nums, func(n int) bool {
		return n > 10
	})
	as.Equal(0, noMatches)

	// Test with empty slice
	emptyCount := CountBy([]int{}, func(n int) bool {
		return true
	})
	as.Equal(0, emptyCount)

	// Test with complex predicate
	complexCount := CountBy(nums, func(n int) bool {
		return n%2 == 0 && n > 3
	})
	as.Equal(2, complexCount)
}

func TestCountValues(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic value counting
	nums := []int{1, 2, 2, 3, 2, 4, 1}
	counts := CountValues(nums)
	as.Equal(map[int]int{
		1: 2,
		2: 3,
		3: 1,
		4: 1,
	}, counts)

	// Test with empty slice
	emptyCounts := CountValues([]int{})
	as.Empty(emptyCounts)

	// Test with strings
	words := []string{"apple", "banana", "apple", "cherry"}
	wordCounts := CountValues(words)
	as.Equal(map[string]int{
		"apple":  2,
		"banana": 1,
		"cherry": 1,
	}, wordCounts)
}

func TestCountValuesBy(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic value counting with mapper
	nums := []int{1, 2, 3, 4, 5, 6}
	counts := CountValuesBy(nums, func(n int) int {
		return n % 3
	})
	as.Equal(map[int]int{
		0: 2, // 3, 6
		1: 2, // 1, 4
		2: 2, // 2, 5
	}, counts)

	// Test with string mapper
	counts2 := CountValuesBy(nums, func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	})
	as.Equal(map[string]int{
		"even": 3,
		"odd":  3,
	}, counts2)

	// Test with empty slice
	emptyCounts := CountValuesBy([]int{}, func(n int) string {
		return "key"
	})
	as.Empty(emptyCounts)
}

func TestSubSet(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test empty slice
	emptyResult := SubSet([]int{}, 0, 0)
	as.Empty(emptyResult)

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Test basic subset
	result := SubSet(nums, 2, 3)
	as.Equal([]int{2, 3, 4}, result)

	// Test negative offset
	negOffset := SubSet(nums, -3, 2)
	as.Equal([]int{7, 8}, negOffset)

	// Test with length exceeding slice bounds
	overLength := SubSet(nums, 8, 5)
	as.Equal([]int{8, 9}, overLength)

	// Test with offset beyond slice length
	beyondLength := SubSet(nums, 15, 1)
	as.Empty(beyondLength)

	// Test with negative offset beyond slice length
	negBeyond := SubSet(nums, -15, 2)
	as.Equal([]int{0, 1}, negBeyond)

	// Test with zero length
	zeroLength := SubSet(nums, 2, 0)
	as.Equal([]int{}, zeroLength)

	// Test with custom slice type
	type IntSlice []int
	customSlice := IntSlice{1, 2, 3, 4, 5}
	customResult := SubSet(customSlice, 1, 2)
	as.Equal(IntSlice{2, 3}, customResult)
}

func TestReplace(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic replacement
	nums := []int{1, 2, 2, 3, 2, 4, 2}
	result := Replace(nums, 2, 5, 2)
	as.Equal([]int{1, 5, 5, 3, 2, 4, 2}, result)

	// Test replace with n = 0 (no replacements)
	noReplace := Replace(nums, 2, 5, 0)
	as.Equal(nums, noReplace)

	// Test replace with n > actual occurrences
	moreReplace := Replace(nums, 2, 5, 10)
	as.Equal([]int{1, 5, 5, 3, 5, 4, 5}, moreReplace)

	// Test with strings
	words := []string{"apple", "banana", "apple", "cherry", "apple"}
	wordResult := Replace(words, "apple", "orange", 2)
	as.Equal([]string{"orange", "banana", "orange", "cherry", "apple"}, wordResult)

	// Test with custom slice type
	type StringSlice []string
	customSlice := StringSlice{"a", "b", "a", "c", "a"}
	customResult := Replace(customSlice, "a", "x", 2)
	as.Equal(StringSlice{"x", "b", "x", "c", "a"}, customResult)
	as.IsType(customResult, customSlice, "should preserve slice type")

	// Test empty slice
	emptyResult := Replace([]int{}, 1, 2, 1)
	as.Empty(emptyResult)

	// Test no matches
	noMatches := Replace(nums, 9, 5, 1)
	as.Equal(nums, noMatches)
}

func TestReplaceAll(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test basic replacement
	nums := []int{1, 2, 2, 3, 2, 4, 2}
	result := ReplaceAll(nums, 2, 5)
	as.Equal([]int{1, 5, 5, 3, 5, 4, 5}, result)

	// Test with no matches
	noMatches := ReplaceAll(nums, 9, 5)
	as.Equal(nums, noMatches)

	// Test with strings
	words := []string{"apple", "banana", "apple", "cherry", "apple"}
	wordResult := ReplaceAll(words, "apple", "orange")
	as.Equal([]string{"orange", "banana", "orange", "cherry", "orange"}, wordResult)

	// Test with custom slice type
	type StringSlice []string
	customSlice := StringSlice{"a", "b", "a", "c", "a"}
	customResult := ReplaceAll(customSlice, "a", "x")
	as.Equal(StringSlice{"x", "b", "x", "c", "x"}, customResult)
	as.IsType(customResult, customSlice, "should preserve slice type")

	// Test empty slice
	emptyResult := ReplaceAll([]int{}, 1, 2)
	as.Empty(emptyResult)

	// Test single element slice
	single := []int{1}
	singleResult := ReplaceAll(single, 1, 2)
	as.Equal([]int{2}, singleResult)

	// Test all elements same
	allSame := []int{2, 2, 2, 2}
	allResult := ReplaceAll(allSame, 2, 3)
	as.Equal([]int{3, 3, 3, 3}, allResult)
}

func TestSlice(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Test basic slicing
	result := Slice(nums, 2, 5)
	as.Equal([]int{2, 3, 4}, result)

	// Test with start >= end
	invalidRange := Slice(nums, 5, 2)
	as.Empty(invalidRange)

	// Test with start > size
	beyondSize := Slice(nums, 15, 20)
	as.Empty(beyondSize)

	// Test with negative start
	negStart := Slice(nums, -2, 5)
	as.Equal([]int{0, 1, 2, 3, 4}, negStart)

	// Test with negative end
	negEnd := Slice(nums, 2, -1)
	as.Empty(negEnd)

	// Test with negative start and end
	negBoth := Slice(nums, -2, -1)
	as.Equal([]int{}, negBoth)

	// Test with end > size
	largeEnd := Slice(nums, 7, 15)
	as.Equal([]int{7, 8, 9}, largeEnd)

	// Test with custom slice type
	type IntSlice []int
	customSlice := IntSlice{1, 2, 3, 4, 5}
	customResult := Slice(customSlice, 1, 3)
	as.Equal(IntSlice{2, 3}, customResult)
	as.IsType(customResult, customSlice, "should preserve slice type")

	// Test empty slice
	emptyResult := Slice([]int{}, 0, 1)
	as.Empty(emptyResult)

	// Test full range
	fullRange := Slice(nums, 0, len(nums))
	as.Equal(nums, fullRange)
}

func TestCompact(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test with integers
	ints := []int{0, 1, 0, 2, 3, 0, 4, 0, 5}
	compactInts := Compact(ints)
	as.Equal([]int{1, 2, 3, 4, 5}, compactInts)

	// Test with empty slice
	emptyInts := []int{}
	as.Empty(Compact(emptyInts))

	// Test with all zeros
	allZeros := []int{0, 0, 0, 0}
	as.Empty(Compact(allZeros))

	// Test with strings
	strs := []string{"", "hello", "", "world", "", ""}
	as.Equal([]string{"hello", "world"}, Compact(strs))

	// Test with custom type
	type customInt int
	customs := []customInt{0, 1, 0, 2, 0}
	as.Equal([]customInt{1, 2}, Compact(customs))
}

func TestIsSorted(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test with integers
	as.True(IsSorted([]int{1, 2, 3, 4, 5}))
	as.True(IsSorted([]int{1, 2, 2, 3, 4, 5}))
	as.True(IsSorted([]int{1}))
	as.True(IsSorted([]int{}))
	as.False(IsSorted([]int{2, 1, 3, 4, 5}))
	as.False(IsSorted([]int{5, 4, 3, 2, 1}))

	// Test with strings
	as.True(IsSorted([]string{"a", "b", "c"}))
	as.False(IsSorted([]string{"b", "a", "c"}))

	// Test with floats
	as.True(IsSorted([]float64{1.1, 2.2, 3.3}))
	as.False(IsSorted([]float64{2.2, 1.1, 3.3}))
}

func TestIsSortedByKey(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	type person struct {
		name string
		age  int
	}

	people := []person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	// Test sorting by age
	as.True(IsSortedByKey(people, func(p person) int { return p.age }))

	// Test unsorted by age
	unsorted := []person{
		{"Bob", 30},
		{"Alice", 25},
		{"Charlie", 35},
	}
	as.False(IsSortedByKey(unsorted, func(p person) int { return p.age }))

	// Test sorting by name
	as.True(IsSortedByKey(people, func(p person) string { return p.name }))

	// Test empty slice
	as.True(IsSortedByKey([]person{}, func(p person) int { return p.age }))

	// Test single element
	as.True(IsSortedByKey([]person{{"Alice", 25}}, func(p person) int { return p.age }))
}

func TestSplice(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test inserting elements in the middle
	original := []int{1, 2, 3, 4, 5}
	spliced := Splice(original, 2, 10, 20)
	as.Equal([]int{1, 2, 10, 20, 3, 4, 5}, spliced)

	// Test inserting at the beginning
	spliced = Splice(original, 0, 0, 0)
	as.Equal([]int{0, 0, 1, 2, 3, 4, 5}, spliced)

	// Test inserting at the end
	spliced = Splice(original, 5, 6, 7)
	as.Equal([]int{1, 2, 3, 4, 5, 6, 7}, spliced)

	// Test with negative index
	spliced = Splice(original, -2, 10)
	as.Equal([]int{1, 2, 3, 10, 4, 5}, spliced)

	// Test with negative index and multiple elements
	spliced = Splice(original, -10, 1, 2)
	as.Equal([]int{1, 2, 1, 2, 3, 4, 5}, spliced)

	// Test with empty elements
	spliced = Splice(original, 2)
	as.Equal(original, spliced)

	// Test with empty slice
	empty := []int{}
	spliced = Splice(empty, 0, 1, 2, 3)
	as.Equal([]int{1, 2, 3}, spliced)

	// Test with index out of bounds
	spliced = Splice(original, 10, 6, 7)
	as.Equal([]int{1, 2, 3, 4, 5, 6, 7}, spliced)

	// Test with custom type
	type customInt int
	customs := []customInt{1, 2, 3}
	splicedCustoms := Splice(customs, 1, customInt(10))
	as.Equal([]customInt{1, 10, 2, 3}, splicedCustoms)
}

func TestEqual(t *testing.T) {
	t.Parallel()
	as := assert.New(t)

	// Test empty slices
	as.True(Equal([]int{}, []int{}))
	as.True(Equal([]string{}, []string{}))

	// Test nil slices
	var nilSlice1, nilSlice2 []int
	as.True(Equal(nilSlice1, nilSlice2))

	// Test small slices (< 100 elements)
	smallInts1 := []int{1, 2, 3, 4, 5}
	smallInts2 := []int{1, 2, 3, 4, 5}
	smallInts3 := []int{1, 2, 3, 4, 6}
	as.True(Equal(smallInts1, smallInts2))
	as.False(Equal(smallInts1, smallInts3))

	// Test different length slices
	as.False(Equal([]int{1, 2, 3}, []int{1, 2}))
	as.False(Equal([]int{1, 2}, []int{1, 2, 3}))

	// Test strings
	as.True(Equal([]string{"a", "b", "c"}, []string{"a", "b", "c"}))
	as.False(Equal([]string{"a", "b", "c"}, []string{"a", "b", "d"}))

	// Test custom types
	type customInt int
	as.True(Equal([]customInt{1, 2, 3}, []customInt{1, 2, 3}))
	as.False(Equal([]customInt{1, 2, 3}, []customInt{1, 2, 4}))

	// Test large slices (>= 100 elements)
	large1 := make([]int, 200)
	large2 := make([]int, 200)
	large3 := make([]int, 200)
	for i := range large1 {
		large1[i] = i
		large2[i] = i
		large3[i] = i
	}
	large3[100] = 999 // Change middle element
	as.True(Equal(large1, large2))
	as.False(Equal(large1, large3))

	// Test with struct type
	type person struct {
		name string
		age  int
	}
	people1 := []person{{"Alice", 25}, {"Bob", 30}}
	people2 := []person{{"Alice", 25}, {"Bob", 30}}
	people3 := []person{{"Alice", 25}, {"Bob", 31}}
	as.True(Equal(people1, people2))
	as.False(Equal(people1, people3))

	// Test with boundary differences in large slices
	large4 := make([]int, 200)
	copy(large4, large1)
	large4[0] = 999 // Change first element
	large5 := make([]int, 200)
	copy(large5, large1)
	large5[199] = 999 // Change last element
	as.False(Equal(large1, large4))
	as.False(Equal(large1, large5))

	// Test with floating point numbers
	floats1 := []float64{1.1, 2.2, 3.3}
	floats2 := []float64{1.1, 2.2, 3.3}
	floats3 := []float64{1.1, 2.2, 3.4}
	as.True(Equal(floats1, floats2))
	as.False(Equal(floats1, floats3))
}
