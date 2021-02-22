package gofp

import (
	"reflect"
	"testing"
)

func Test_StringToInterfaceSlice(t *testing.T) {
	slice := StringToInterfaceSlice([]string{"name", "age"})
	if reflect.TypeOf(slice).Elem().Kind() != reflect.Interface {
		t.Errorf("StringToInterfaceSlice() got= %q, want %q", reflect.TypeOf(slice).Elem().Kind(), reflect.Interface)
	}
}

func Test_StringSlice(t *testing.T) {
	slice := StringSlice([]interface{}{"name", "age"})
	if reflect.TypeOf(slice).Elem().Kind() != reflect.String {
		t.Errorf("StringSlice() got= %q, want %q", reflect.TypeOf(slice).Elem().Kind(), reflect.String)
	}
}

func Test_IntSlice(t *testing.T) {
	slice := IntSlice([]interface{}{1, 2, 3})
	if reflect.TypeOf(slice).Elem().Kind() != reflect.Int {
		t.Errorf("IntSlice() got= %q, want %q", reflect.TypeOf(slice).Elem().Kind(), reflect.Int)
	}
}

func Test_Float64Slice(t *testing.T) {
	slice := Float64Slice([]interface{}{1.4, 2.5, 3.1})
	if reflect.TypeOf(slice).Elem().Kind() != reflect.Float64 {
		t.Errorf("Float64Slice() got= %q, want %q", reflect.TypeOf(slice).Elem().Kind(), reflect.Float64)
	}
}

func Test_Randomer(t *testing.T) {
	randomer := Randomer()
	if reflect.TypeOf(randomer.Intn(5)).Kind() != reflect.Int {
		t.Errorf("Randomer() got= %q, want %q", reflect.TypeOf(randomer.Intn(5)).Kind(), reflect.Int)
	}
}
