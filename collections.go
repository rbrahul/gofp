package gofp

/* Collection Utils
TODO:
- groupBy()
- shuffle()
- random()
- fill()
- unique()
*/

// Map returns a new slice with transformed elements
func Map(fn func(index int, item interface{}) interface{}, items []interface{}) []interface{} {
	mappedItems := []interface{}{}
	for index, value := range items {
		mappedItems = append(mappedItems, fn(index, value))
	}
	return mappedItems
}

// Filter returns a new slice of items which satisfies the condition
func Filter(fn func(index int, item interface{}) bool, items []interface{}) []interface{} {
	filteredItems := []interface{}{}
	for index, value := range items {
		if fn(index, value) {
			filteredItems = append(filteredItems, value)
		}
	}
	return filteredItems
}

type reduceFnArgType func(index int, current interface{}, accumulator interface{}, source []interface{}) interface{}

// Reduce iterate overs all the items in the slice and returns accumulated result
func Reduce(fn reduceFnArgType, items []interface{}, initialValue interface{}) interface{} {
	accumulator := initialValue
	for index, value := range items {
		accumulator = fn(index, value, accumulator, items)
	}
	return accumulator
}

// Every returns true if all the items satisfies the given condition with the function
func Every(fn func(index int, item interface{}) bool, items []interface{}) bool {
	for index, value := range items {
		if !fn(index, value) {
			return false
		}
	}
	return true
}

// Any returns true if any of the item satisfies the given condition with the function
func Any(fn func(index int, item interface{}) bool, items []interface{}) bool {
	for index, value := range items {
		if fn(index, value) {
			return true
		}
	}
	return false
}

// Head returns the first item of slice if exist otherwise nil
func Head(items []interface{}) interface{} {
	if len(items) >= 1 {
		return items[0]
	}
	return nil
}

// Tail returns the last item of slice if exist otherwise nil
func Tail(items []interface{}) interface{} {
	if len(items) >= 1 {
		return items[len(items)-1]
	}
	return nil
}

// Reverse returns a new slice of reversed items
func Reverse(items []interface{}) []interface{} {
	reversed := []interface{}{}
	for i := len(items) - 1; i >= 0; i-- {
		reversed = append(reversed, items[i])
	}
	return reversed
}

// Chunk returns a new slice of reversed items
func Chunk(items []interface{}, size int) []interface{} {
	chunks := []interface{}{}
	startAt := 0
	lengthOfItems := len(items)
	if lengthOfItems == 0 {
		return chunks
	}
	if size >= lengthOfItems {
		return append(chunks, items)
	}
	for startAt <= lengthOfItems {
		upperLimit := startAt + size
		if upperLimit < lengthOfItems {
			chunks = append(chunks, items[startAt:upperLimit])
		} else {
			chunks = append(chunks, items[startAt:])
		}
		startAt = upperLimit
	}
	return chunks
}

// Range returns a new array with elements starting from min to max
func Range(args ...int) []int {
	var (
		min  = 1
		max  int
		step = 1
	)
	if len(args) == 1 {
		max = args[0]
	}
	if len(args) >= 2 {
		min = args[0]
		max = args[1]
		if len(args) >= 3 {
			step = args[2]
		}
	}

	items := []int{}
	for i := min; i <= max; i += step {
		items = append(items, i)
	}
	return items
}
