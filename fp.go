package gofp

/* TODO: FP Utils
- Curry()
*/

// FunctionAsArg represents the type of fuction passed as operation
type FunctionAsArg func(arg interface{}) interface{}

// OperationHandler represents the type of fuction which accpets data nd returns processed result
type OperationHandler func(interface{}) interface{}

// Pipe allows to process the chainable functional operations
func Pipe(fns ...(FunctionAsArg)) OperationHandler {
	return func(data interface{}) interface{} {
		var result interface{} = data
		for _, fn := range fns {
			result = fn(result)
		}
		return result
	}
}

// Compose allows to process the chainable functional operations in reverse order
func Compose(fns ...(FunctionAsArg)) OperationHandler {
	return func(data interface{}) interface{} {
		var result interface{} = data

		for i := len(fns) - 1; i >= 0; i-- {
			result = fns[i](result)
		}
		return result
	}
}
