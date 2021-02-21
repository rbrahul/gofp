![Go test workflow](https://github.com/rbrahul/gofp/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/rbrahul/gofp)](https://goreportcard.com/report/github.com/rbrahul/gofp)
[![Go Reference](https://pkg.go.dev/badge/github.com/rbrahul/gofp.svg)](https://pkg.go.dev/github.com/rbrahul/gofp)

## A simple Utility library for Go

Go does not provide a many essential functions while working with the data structure like Slice and Map. This library provides most frequently needed utility functions which is inspired from lodash(a Javascript Utility library).

## Why do I need gofp ?

- Implementing Functional programming is way easier using `Pipe(), Compose(), Reduce(), Map(), Filter(), Extend(), Find() and many others`.

- This library offers many utility function for dealing with collections or slice related operation.

- Access any property by path or index from the map, slice and even struct by simply using the most useful function `Get`.

- Utility functions are implmented based on `interface{}`. The main focus is not to use the `reflect` package whenever possible.


## Installation
Please run the following command in terminal to install

```
go get github.com/rbrahul/gofp
```

## How to use?

This the example which describes how you can implement `pipe` operations using `gofp`

```go

package main

import (
"fmt"
"strings"
"github.com/rbrahul/gofp"
)

func main() {
	user := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
		"contacts": map[string]interface{}{
			"email":  "johndoe@gmail.com",
			"office": "Google Inc.",
			"fax": map[string]interface{}{
				"uk": "+44-208-1234567",
			},
		},
	}
	getContacts := func(data interface{}) interface{} {
		return data.(map[string]interface{})["contacts"]
	}

	getEmail := func(data interface{}) interface{} {
		return data.(map[string]interface{})["email"]
	}
	getUpperCaseEmail := func(data interface{}) interface{} {
		return strings.ToUpper(data.(string))
	}

	email := gofp.Pipe(
		getContacts,
		getEmail,
		getUpperCaseEmail,
	)(user)

    fmt.Println("Email is: ", email) // Output: Email is: JOHNDOE@GMAIL.COM

}
```


## Documentation:

### Most commonly used utility functions for Collection or slice

### Map()
Returns a new slice executing the iterator function on each element. Map has 2 parameters, 1st one is slice and 2nd one is the iterator function. The iterator function must have 2 parameters, index and current value of that iteration.

```go
    ...
    mappedItems := Map([]interface{}{1, 2, 3, 4, 5},
            func(i int, item interface{}) interface{} {
                return item.(int) * item.(int)
            })
    
    fmt.Println(mappedItems) //Output: 1, 4, 9, 16, 25
    ...
```

### Filter()
Returns a new slice containing the filtered elements. The new slice contains those elements who satisfy the condition of the iterator function. Filter has 2 parameters, 1st one is the slice and 2nd one is the iterator function. The iterator function must have 2 parameters, index and current value of that iteration.

```go
    ...
        filteredItems := Filter([]interface{}{12, 16, 18, 20, 23, 40, 25},
            func(i int, age interface{}) bool {
                return age.(int) >= 20
            })
    
    fmt.Println(filteredItems) //Output:  20, 23, 40, 25
    ...
```


### Find()
Returns the first matched element of the slice who satisfy the condition of iterator function. If there is no such element that satisfy the condition of the function then nil is returned. Find has 2 parameters, 1st one is the slice and 2nd one is the iterator function. The iterator function must have 2 parameters, index and current value of that iteration.

```go
    ...
      user := Find([]interface{}{
            map[string]interface{}{"name": "Ron", "sex": "male", "age": 17},
            map[string]interface{}{"name": "Raymond", "sex": "male", "age": 20},
            map[string]interface{}{"name": "Sofia", "sex": "female", "age": 20},
            map[string]interface{}{"name": "Roni", "sex": "male", "age": 30},
        }, func(i int, person interface{}) bool {
            return person.(map[string]interface{})["age"].(int) >= 18
        })
    fmt.Println(user) //Output:  {"name": "Raymond", "sex": "male", "age": 20}
    ...
```
### Reduce()
Executes a iterator function on each element of the slice, resulting in single output accumulated value. Reducer has 3 parameters, 1st one is the slice and 2nd one is the iterator function and 3rd one is the initial value. The iterator function must have 3 parameters which are index, current value of that iteration and accumulated value or result of previous iterations.

```go
    ...
     reducedItems := Reduce([]interface{}{10, 20, 30, 40},
        func(index int, current interface{}, accumulator interface{}, source []interface{}) interface{} {
            return accumulator.(int) + current.(int)
	    }, 0)
    fmt.Println(reducedItems) //Output:  100
    ...
```
### Every()
Returns `true` if each element matches the condition of the given iterator function. If there is any element that doesn't satisfy the condition of the function then it returns `false`. Every has 2 parameters, 1st one is the slice and 2nd one is the iterator function. And the iterator function must have 2 parameters, index and current value of that iteration.

```go
    ...
     isEveryOneIsAdult := Every([]interface{}{18, 20, 23, 40, 25},
        func(i int, age interface{}) bool {
		    return age.(int) >= 18
	    })
    fmt.Println(isEveryOneIsAdult) //Output:  true
    ...
```

### Any()
Returns `true` if any of the element matches the condition of the given iterator function. If there is no such element that satisfy the condition of the function then it returns `false`. Any has 2 parameters, 1st one is the slice and 2nd one is the iterator function. And the iterator function must have 2 parameters, index and current value of that iteration.

```go
    ...
     hasAnyAdult := Any([]interface{}{18, 20, 23, 40, 25},
        func(i int, age interface{}) bool {
		    return age.(int) >= 18
	    })
    fmt.Println(hasAnyAdult) //Output:  true
    ...
```

### GroupBy()

Returns a new map composed of keys generated from the results of running each element of slice thru iterator function. The order of grouped values is determined by the order they occur in slice. GroupBy has 2 parameters, 1st one is the slice and 2nd one is the iterator function. The output of the iterator function will be used as the key of the newly created group or map. And the iterator function must have 2 parameters, index and current value of that iteration.

```go
    ...
     groupedData := GroupBy([]interface{}{
            map[string]interface{}{"name": "Ron", "sex": "male", "age": 17},
            map[string]interface{}{"name": "Raymond", "sex": "male", "age": 20},
            map[string]interface{}{"name": "Sofia", "sex": "female", "age": 20},
            map[string]interface{}{"name": "Roni", "sex": "male", "age": 30},
	    }, func(person interface{}) string {
		    return strconv.Itoa(person.(map[string]interface{})["age"].(int))
	})
    fmt.Println(groupedData) 
    /*
    Output:
      { 
          "17": [{"name": "Ron", "sex": "male", "age": 17}],
          "20": [
                {"name": "Raymond", "sex": "male", "age": 20},
                {"name": "Sofia", "sex": "female", "age": 20}
               ],
          "30": [{"name": "Roni", "sex": "male", "age": 30}]
     }
    */
    ...
```

### Chunk()

Returns a new slice(chunks) of slices. Every slice has fixed number of elements which was given as a limit in the 2nd parameter. Chunk accepts 2 parameters, 1st one is the slice and 2nd one is the limit which will define the maxium number of elements in each slice.

```go
    ...
	chunkedItems := Chunk([]interface{}{1, 2, 3, 4, 5}, 2)
    fmt.Println(chunkedItems) //Output:  {{1,2},{3,4},{5}}
    ...
```
### Reverse()

Returns a new slice with all the elements in reveresed order. Reverse accepsts 1 parameter which a slice.

```go
    ...
    reveresed := Reverse([]interface{}{10, 20, 30, 40, 50})
    fmt.Println(reveresed) //Output:  {50,40,30,20,10}
    ...
```

### Range()

Returns a new slice of range where the value starts from 1st parameter to the 2nd parameter. Reverse accepsts 2 parameters, 1st one is the starting value 2nd one is the maximum value in the range.

```go
    ...
    rangeItems := Range(5, 10)
    fmt.Println(rangeItems) //Output:  {5,6,7,8,9,10}
    ...
```

### Uniq()

Returns a new slice where each elements are unique removing all the duplicate elements. `Uniq` accepsts 1 parameter which is a slice.

```go
    ...
    // [1,2,3,10,4,5,100]
	uniqueItems := Uniq([]interface{}{1, 2, 2, 3, 10, 4, 5, 10, 100})
    fmt.Println(uniqueItems) //Output:  {1,2,3,10,4,5,100}
    ...
```

### Head()

Returns the first matched element of the slice. Head accepsts 1 parameter which a slice.

```go
    ...
      firstItem := Head([]interface{}{
            map[string]interface{}{"name": "Ron", "sex": "male", "age": 17},
            map[string]interface{}{"name": "Raymond", "sex": "male", "age": 20},
            map[string]interface{}{"name": "Sofia", "sex": "female", "age": 20},
            map[string]interface{}{"name": "Roni", "sex": "male", "age": 30},
        })
    fmt.Println(firstItem) //Output:  {"name": "Ron", "sex": "male", "age": 17}
    ...
```

### Tail()

Returns the last matched element of the slice. Head accepsts 1 parameter which a slice.

```go
    ...
      lastItem := Tail([]interface{}{
            map[string]interface{}{"name": "Ron", "sex": "male", "age": 17},
            map[string]interface{}{"name": "Raymond", "sex": "male", "age": 20},
            map[string]interface{}{"name": "Sofia", "sex": "female", "age": 20},
            map[string]interface{}{"name": "Roni", "sex": "male", "age": 30},
        })
    fmt.Println(lastItem) //Output:  {"name": "Roni", "sex": "male", "age": 30}
    ...
```

### Fill()

Returns a new slice where every elements is replaced from the start to end index with the given string. `Fill` has 4 arguments first 2 are required and last two are optional. First one is slice, 2nd one is the string which will be used as substitute while filling/replacing and 3rd one is the starting index and 4th one is the end index. If start and end index is not given then it fills all the elements with given string.

```go
    ...
	filledItems := Fill([]interface{}{1, 2, 3, 4, 5, 6, 7}, "*", 1, 5)
    fmt.Println(filledItems) //Output:  {1, *, *, *, *, 6, 7}
    ...
```

### IndexOf()

Returns the index of the first occurance of any element in the slice which is equal to the given item.

```go
    ...
	index := IndexOf([]interface{}{1, 2, 2, 3, 10, 4, 5, 10, 100}, 10)
    fmt.Println(index) //Output: 4
    ...
```

### Contains()

Returns `true` if the given item exists in the slice or false otherwise.

```go
    ...
    exists := Contains([]interface{}{1, 2, 2, 3, 10, 4, 5, 10, 100}, 10)
    fmt.Println(exists) //Output: true
    ...
```

### ChooseRandom()

Returns a randomly selected element of the slice. It has one parameter which is a slice.

```go
    ...
    randomElement := ChooseRandom([]interface{}{1, 2, 3, 4, 5, 10, 100})
    fmt.Println("Could be any:",randomElement) //Output Could be any: 4 
    ...
```

### Shuffle()

Returns a new slice where elements are randomly ordered(shuffled). It accepts one parameter which is slice.

```go
    ...
    shuffledItems := ChooseRandom([]interface{}{1, 2, 3, 4, 5, 10, 100})
    fmt.Println(shuffledItems) //Output: {100, 2, 1, 4, 5, 3, 10} 
    ...
```
## Map related utitlity function:

### Keys():

Returns a slice of keys of the map.  

```go
    ...
    keys := Keys(map[string]interface{}{
        "firstName": "John", 
        "lastName": "Doe",
        "age": 32
        })
    fmt.Println(keys) //Output: {firstName, lastName, age}
    ...
```

### Values():

Returns a slice of values of the map.

```go
    ...
    values := Values(map[string]interface{}{
        "firstName": "John", 
        "lastName": "Doe",
        "age": 32
        })
    fmt.Println(values) //Output: {John, Doe, 32}
    ...
```

### Has():

Returns `true` if key exists in the map or false otherwise. It has two parametes 1st one is a map and 2nd one is key.

```go
    ...
    exists := Has(map[string]interface{}{
        "firstName": "John", 
        "lastName": "Doe",
        "age": 32
        }, "age")
    fmt.Println(exists) //Output: true
    ...
```

### Pick():

Returns a new map containing only properties which are specified as 2nd argument. It accepts 2 parameters 1st one is the map and second one is the keys which is a slice of string.

```go
    ...
	pickedData := Pick(map[string]interface{}{
        "firstName": "John", 
        "lastName": "Doe",
        "age": 32
        }, []string{"lastName"})
    fmt.Println(pickedData) //Output: {"lastName": "Doe"}
    ...
```

### Omit():

Returns a new map omitting the given keys of that map. It accepts 2 parameters 1st one is the map and second one is the keys which is a slice of string.

```go
    ...
	omittedData := Omit(map[string]interface{}{
        "firstName": "John", 
        "lastName": "Doe",
        "age": 32
        }, []string{"lastName"})
    fmt.Println(omittedData) //Output: {"firstName": "John", "age": 32}
    ...
```

### MapValues():

`MapValues` works similarly to the `Map()` unlikely it deals with only map. It returns a new map applying an iterator function on each `value` of the map. The iterator function transforms the each value.  

```go
    ...
	mappedValues := MapValues(map[string]interface{}{
        "firstName": "john",
        "lastName": "doe",
        "gender": "unknown"}, func(value interface{}) interface{} {
		return strings.Title(value.(string))
	})
    fmt.Println(mappedValues) //Output: {"firstName": "JOHN",  "lastName":"DOE": 32, "gender":"UNKNOWN"}
    ...
```

### MapKeys():

`MapKeys` works similarly to the `Map()` unlikely it deals with only map. It returns a new map applying an iterator function on each `key` of the map. The iterator function transforms the each `key`.  

```go
    ...
	mappedKeys := MapKeys(map[string]interface{}{
        "firstName": "john",
        "lastName": "doe",
        "gender": "unknown"}, func(value interface{}) interface{} {
		return strings.Title(value.(string))
	})
    fmt.Println(mappedKeys) //Output: {"FIRSTNAME": "john","LASTNAME": "doe","GENDER": "unknown"}
    ...
```

### Get():

`Get()` returns the value of a given path. If no data is available in the given path then `nil` is returned. It deals with `map`, `slice`, and `struct`. It accepts 3 parameters, `data`, `path` and `fallback value`. 3rd parameter is optional. 

```go
    ...
	data := map[string]interface{}{
		"age":  30,
		"male": true,
		"contacts": map[string]interface{}{
			"office": 12345,
			"fax": map[string]interface{}{
				"uk": "+44-208-1234567",
			},
			"address": map[string]interface{}{
				"post_code":    "SW3",
				"street":       "10 Downing Street",
				"geo_location": []string{"51.529011463529636", "-0.1098365614770662"},
			},
		},
	}
	geoLocationFromGet := Get(data, "contacts.address.geo_location.0")
    fmt.Println(geoLocationFromGet) //Output: 51.529011463529636
    ...
```

### Extend():

`Extend()` returns a new map extending the values with a given map. Where extend or override operation happens deeply(recursively). It accepts two parameters both are map. 1st map gets extended with the 2nd map.

```go
    ...
	extendedMap := Extends(
        map[string]interface{}{
        "firstName": "john",
        "lastName": "doe",
        "gender": "unknown"
        },
         map[string]interface{}{
        "gender": "male"
        })
    fmt.Println(extendedMap) //Output: {"firstName": "john","lastName": "doe","gender": "male"}
    ...
```
