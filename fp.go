package gofp

/* TODO: FP Utils
- Curry()
*/

// Pipe allows to process the chainable functional operations
func Pipe(fns ...(func(arg interface{}) interface{})) func(interface{}) interface{} {
	return func(data interface{}) interface{} {
		var result interface{} = data
		for _, fn := range fns {
			result = fn(result)
		}
		return result
	}
}

// Compose allows to process the chainable functional operations in reverse order
func Compose(fns ...(func(arg interface{}) interface{})) func(interface{}) interface{} {
	return func(data interface{}) interface{} {
		var result interface{} = data

		for i := len(fns) - 1; i >= 0; i-- {
			result = fns[i](result)
		}
		return result
	}
}
