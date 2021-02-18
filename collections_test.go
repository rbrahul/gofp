package gofp

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_Map(t *testing.T) {
	mappedItems := Map([]interface{}{1, 2, 3, 4, 5}, func(i int, item interface{}) interface{} {
		return item.(int) * item.(int)
	})
	got := mappedItems[len(mappedItems)-1]
	want := 25
	if mappedItems[len(mappedItems)-1] != want {
		t.Errorf("Map() = %q, want %q", got, want)
	}
}
func Test_Fill(t *testing.T) {
	inputSlice := []interface{}{1, 2, 3, 4, 5, 6, 7}
	got := Fill(inputSlice, "*", 1, 5)
	result := Every(got[1:5], func(index int, item interface{}) bool {
		return fmt.Sprintf("%v", item) == "*"
	})
	if !result {
		t.Errorf("Fill() = %v, want %v", result, true)
	}
}
func Test_Every(t *testing.T) {
	isEveryOneIsAdult := Every([]interface{}{18, 20, 23, 40, 25}, func(i int, age interface{}) bool {
		return age.(int) >= 18
	})
	if !isEveryOneIsAdult {
		t.Errorf("Every() = %v, want %v", isEveryOneIsAdult, true)
	}
}

func Test_Any(t *testing.T) {
	isAnyoneFortieth := Any([]interface{}{18, 20, 23, 40, 25}, func(i int, age interface{}) bool {
		return age.(int) >= 40
	})
	if !isAnyoneFortieth {
		t.Errorf("Any() = %v, want %v", isAnyoneFortieth, true)
	}
}

func Test_Find(t *testing.T) {
	firstAdult := Find([]interface{}{
		map[string]interface{}{"name": "Ron", "sex": "male", "age": 17},
		map[string]interface{}{"name": "Raymond", "sex": "male", "age": 20},
		map[string]interface{}{"name": "Sofia", "sex": "female", "age": 20},
		map[string]interface{}{"name": "Roni", "sex": "male", "age": 30},
	}, func(i int, person interface{}) bool {
		return person.(map[string]interface{})["age"].(int) >= 18
	})
	name := firstAdult.(map[string]interface{})["name"]
	if name != "Raymond" {
		t.Errorf("Find() = %v, want %v", name, "Raymond")
	}
}

func Test_GroupBy(t *testing.T) {
	groupedData := GroupBy([]interface{}{
		map[string]interface{}{"name": "Ron", "sex": "male", "age": 17},
		map[string]interface{}{"name": "Raymond", "sex": "male", "age": 20},
		map[string]interface{}{"name": "Sofia", "sex": "female", "age": 20},
		map[string]interface{}{"name": "Roni", "sex": "male", "age": 30},
	}, func(person interface{}) string {
		return strconv.Itoa(person.(map[string]interface{})["age"].(int))
	})
	data, _ := groupedData["20"]
	numItemsInsideGroup20 := len(data.([]interface{}))
	if numItemsInsideGroup20 != 2 {
		t.Errorf("Find() = %v, want %v", numItemsInsideGroup20, 2)
	}
}

func Test_Head(t *testing.T) {
	firstElement := Head([]interface{}{10, 20, 30, 40, 50})
	if firstElement != 10 {
		t.Errorf("Head() = %v, want %v", firstElement, 10)
	}
}

func Test_Tail(t *testing.T) {
	lastElement := Tail([]interface{}{10, 20, 30, 40, 50})
	if lastElement != 50 {
		t.Errorf("Head() = %v, want %v", lastElement, 50)
	}
}

func Test_Reverse(t *testing.T) {
	reveresed := Reverse([]interface{}{10, 20, 30, 40, 50})
	if reveresed[0].(int) != 50 && reveresed[len(reveresed)-1].(int) != 10 {
		t.Errorf("Head() = %v, want %v", reveresed[0], 50)
	}
}

func Test_Chunk(t *testing.T) {
	// [[1,2][3,4],[5]]
	chunkedItems := Chunk([]interface{}{1, 2, 3, 4, 5}, 2)
	if (len(chunkedItems[0].([]interface{})) != 2) || (chunkedItems[0].([]interface{})[0].(int) != 1) || (chunkedItems[0].([]interface{})[1].(int) != 2) {
		t.Errorf("Chunk() = %v, want %v", len(chunkedItems[0].([]interface{})), 2)
	}
}

func Test_Range(t *testing.T) {
	// [5,6,7,8,9,10]
	rangeItems := Range(5, 10)
	if (len(rangeItems) != 6) || (rangeItems[0] != 5) || (rangeItems[len(rangeItems)-1] != 10) {
		t.Errorf("Range() = %v, want %v", len(rangeItems), 6)
	}
}
