## A simple Utility library for Go

Go doesn't not provide lot of useful methods while working with data structure like Slice and Map. This library provides most frequently needed utility functions which is inspired from lodash(a Javascript Utility library).

## Why do I need this

- Implementing Functional programming is way easier using `Pipe(), Compose(), Reduce(), Map(), Filter(), Extend(), Find() etc.`

- This library offers many utility function for dealing with collections or slice related operation

- Access any property by path or index from the map, slice and even struct by simply using the most useful function `Get`

- Utility functions are implmented based on `interface{}`. The main focus is not to use the `reflect` package whenever possible.


## Installation
Please run the following command in terminal to install

```
go get github.com/rbrahul/gofp
```

## How to use?
This the example how you can implement pipe operations using `gofp`
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

### Map
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

### Filter
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


### Find
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
### Reduce
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
### Every
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