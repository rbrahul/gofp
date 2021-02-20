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