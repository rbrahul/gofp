package gofp

/**MAP UTILS
- get()
- merge()
*/

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
