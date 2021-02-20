package gofp

// Map returns a new slice with transformed elements
func Map(items []interface{}, fn func(index int, item interface{}) interface{}) []interface{} {
	mappedItems := []interface{}{}
	for index, value := range items {
		mappedItems = append(mappedItems, fn(index, value))
	}
	return mappedItems
}

//Fill substitues the elements of slice with given string from the start to end position
func Fill(args ...interface{}) []interface{} {
	var (
		filler string
		items  []interface{}
		start  int
		end    int
	)

	if len(args) < 2 {
		panic("Invalid number of arguments has been passed, at least 2 arguments are required")
	}
	if len(args) >= 1 {
		items = args[0].([]interface{})
		start = 0
		end = len(items)
	}

	if len(args) >= 2 {
		filler = args[1].(string)
	}

	if len(args) >= 3 {
		start = args[2].(int)
	}

	if len(args) >= 4 {
		end = args[3].(int)
	}

	newItems := []interface{}{}
	for index, value := range items {
		if index >= start && index < end {
			newItems = append(newItems, filler)
			continue
		}
		newItems = append(newItems, value)
	}
	return newItems
}

// Filter returns a new slice of items which satisfies the condition
func Filter(items []interface{}, fn func(index int, item interface{}) bool) []interface{} {
	filteredItems := []interface{}{}
	for index, value := range items {
		if fn(index, value) {
			filteredItems = append(filteredItems, value)
		}
	}
	return filteredItems
}

// Reduce iterate overs all the items in the slice and returns accumulated result
func Reduce(items []interface{}, fn func(index int, current interface{}, accumulator interface{}, source []interface{}) interface{}, initialValue interface{}) interface{} {
	accumulator := initialValue
	for index, value := range items {
		accumulator = fn(index, value, accumulator, items)
	}
	return accumulator
}

// Every returns true if all the items satisfies the given condition with the function
func Every(items []interface{}, fn func(index int, item interface{}) bool) bool {
	for index, value := range items {
		if !fn(index, value) {
			return false
		}
	}
	return true
}

// Any returns true if any of the item satisfies the given condition with the function
func Any(items []interface{}, fn func(index int, item interface{}) bool) bool {
	for index, value := range items {
		if fn(index, value) {
			return true
		}
	}
	return false
}

// Find returns a item from the slice if that element satisfies the given condition with the function
func Find(items []interface{}, fn func(index int, item interface{}) bool) interface{} {
	for index, value := range items {
		if fn(index, value) {
			return value
		}
	}
	return nil
}

// GroupBy returns a item from the slice if that element satisfies the given condition with the function
func GroupBy(items []interface{}, fn func(item interface{}) string) map[string]interface{} {
	group := map[string]interface{}{}
	for _, value := range items {
		key := fn(value)
		if Has(group, key) {
			group[key] = append(group[key].([]interface{}), value)
		} else {
			items := []interface{}{}
			group[key] = append(items, value)
		}
	}
	return group
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

// Chunk Returns a new slice(chunks) of slices. Every slice has fixed number of elements which was given as a limit in the 2nd parameter
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

// Uniq returns a new slice of unique items
func Uniq(items []interface{}) []interface{} {
	uniqueItems := []interface{}{}
	for _, item := range items {
		if !Contains(uniqueItems, item) {
			uniqueItems = append(uniqueItems, item)
		}
	}
	return uniqueItems
}

// IndexOf returns the poisition of the item in a slice, if item doesn't exist returns -1 otherwise
func IndexOf(items []interface{}, item interface{}) int {
	for index, value := range items {
		if value == item {
			return index
		}
	}
	return -1
}

// Contains returns true if item exists in the slice and false otherwise
func Contains(items []interface{}, item interface{}) bool {
	return IndexOf(items, item) > -1
}

//Shuffle returns a new slice with shuffled elements
func Shuffle(items []interface{}) []interface{} {
	copiedSlice := make([]interface{}, len(items))
	for i := range items {
		copiedSlice[i] = items[i]
	}
	for i := range items {
		index := Randomer().Intn(len(items) - 1)
		temp := copiedSlice[index]
		copiedSlice[index] = copiedSlice[i]
		copiedSlice[i] = temp
	}
	return copiedSlice
}

var prevChosenItem interface{}

//ChooseRandom returns a random element from the slice
func ChooseRandom(items []interface{}) interface{} {
	index := Randomer().Intn(len(items) - 1)
	for items[index] == prevChosenItem {
		index = Randomer().Intn(len(items) - 1)
	}
	prevChosenItem = items[index]
	return items[index]
}
