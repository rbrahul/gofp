package gofp

import (
	"reflect"
	"strconv"
	"strings"
)

// Keys returns all the keys of any map
func Keys(mapData map[string]interface{}) []string {
	keys := []string{}
	for key := range mapData {
		keys = append(keys, key)
	}
	return keys
}

// Values returns all the keys of any map
func Values(mapData map[string]interface{}) []interface{} {
	values := []interface{}{}
	for _, value := range mapData {
		values = append(values, value)
	}
	return values
}

// Omit returns a new map containing keys that doesn't exists in the provided omittable Keys
func Omit(mapData map[string]interface{}, omittableKeys []string) map[string]interface{} {
	newMap := map[string]interface{}{}
	for key, value := range mapData {
		if !Contains(StringToInterfaceSlice(omittableKeys), key) {
			newMap[key] = value
		}
	}
	return newMap
}

// MapValues returns a map transforming the values applying provided function
func MapValues(mapData map[string]interface{}, fn func(interface{}) interface{}) map[string]interface{} {
	newMap := map[string]interface{}{}
	for key, value := range mapData {
		newMap[key] = fn(value)
	}
	return newMap
}

// MapKeys returns a map transforming the keys applying provided function
func MapKeys(mapData map[string]interface{}, fn func(interface{}) interface{}) map[string]interface{} {
	newMap := map[string]interface{}{}
	for key, value := range mapData {
		newKey := fn(key).(string)
		newMap[newKey] = value
	}
	return newMap
}

// Pick returns a new map with matched keys
func Pick(mapData map[string]interface{}, keys []string) map[string]interface{} {
	newMap := map[string]interface{}{}
	for _, key := range keys {
		value, ok := mapData[key]
		if ok {
			newMap[key] = value
		}
	}
	return newMap
}

// Has returns all the keys of any map
func Has(mapData map[string]interface{}, key string) (exists bool) {
	_, exists = mapData[key]
	return
}

func isMap(data interface{}) (isMap bool) {
	if data == nil {
		return false
	}
	_, isMap = data.(map[string]interface{})
	return
}

// Extend returns a map extending all the property with given map
func Extend(initialMap map[string]interface{}, extendingMap map[string]interface{}) map[string]interface{} {
	newMap := map[string]interface{}{}
	for key, value := range initialMap {
		newMap[key] = value
	}
	for key, value := range extendingMap {
		if isMap(value) && Has(initialMap, key) && isMap(initialMap[key]) {
			newMap[key] = Extend(initialMap[key].(map[string]interface{}), value.(map[string]interface{}))
		} else {
			newMap[key] = value
		}
	}
	return newMap
}

//Get returns the value by path, if path is invalid returns nil
func Get(args ...interface{}) interface{} {
	if len(args) < 2 {
		panic("Invalid number of argument. Atleast 2 arguments are required")
	}
	var (
		path     string
		mapData  interface{}
		fallback interface{}
	)
	if len(args) >= 1 {
		mapData = args[0]
	}
	if len(args) >= 2 {
		path = args[1].(string)
	}
	if len(args) >= 3 {
		fallback = args[2]
	}
	if mapData == nil {
		return fallback
	}
	defer func() interface{} {
		if r := recover(); r != nil {
			return fallback
		}
		return fallback
	}()
	data := mapData
	paths := strings.Split(path, ".")
	for _, key := range paths {
		value := reflect.ValueOf(data)
		dataType := value.Type().Kind()
		if dataType == reflect.Map {
			data = data.(map[string]interface{})[key]
			continue
		}

		if dataType == reflect.Slice {
			indx, err := strconv.Atoi(key)
			if err != nil {
				return fallback
			}
			data = value.Index(indx).Interface().(interface{})
			continue
		}

		if dataType == reflect.Struct {
			data = value.FieldByName(key).Interface().(interface{})
			continue
		}

		if dataType == reflect.String {
			indx, err := strconv.Atoi(key)
			if err != nil {
				return fallback
			}
			data = string(data.(string)[indx])
			continue
		}
		return fallback
	}
	return data
}
