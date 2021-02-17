package gofp

import (
	"math/rand"
	"time"
)

// get
// set
// equal
// deepEual

func randomer() *rand.Rand {
	seed := rand.NewSource(time.Now().UnixNano())
	return rand.New(seed)
}

//StringToInterfaceSlice converts a slice of string to slice of interface
func StringToInterfaceSlice(stringSlice []string) []interface{} {
	interfaceSlice := make([]interface{}, len(stringSlice), len(stringSlice))
	for i := range stringSlice {
		interfaceSlice[i] = stringSlice[i]
	}
	return interfaceSlice
}

//StringSlice converts a slice of interface to slice of string
func StringSlice(interfaceSlice []interface{}) []string {
	stringSlice := make([]string, len(interfaceSlice), len(interfaceSlice))
	for i := range interfaceSlice {
		stringSlice[i] = interfaceSlice[i].(string)
	}
	return stringSlice
}

//IntSlice converts a slice of interface to slice of int
func IntSlice(interfaceSlice []interface{}) []int {
	intSlice := make([]int, len(interfaceSlice), len(interfaceSlice))
	for i := range interfaceSlice {
		intSlice[i] = interfaceSlice[i].(int)
	}
	return intSlice
}

//Float64Slice converts a slice of interface to slice of int
func Float64Slice(interfaceSlice []interface{}) []float64 {
	float64Slice := make([]float64, len(interfaceSlice), len(interfaceSlice))
	for i := range interfaceSlice {
		float64Slice[i] = interfaceSlice[i].(float64)
	}
	return float64Slice
}
