package gofp

import (
	"strings"
	"testing"
)

func Test_Keys(t *testing.T) {
	keys := Keys(map[string]interface{}{"name": "Rahul", "age": 32})
	if len(keys) != 2 {
		t.Errorf("Keys() got= %q, want %q", len(keys), 2)
	}
}

func Test_Values(t *testing.T) {
	values := Values(map[string]interface{}{"name": "Rahul", "age": 32})
	if len(values) != 2 {
		t.Errorf("Values() got= %q, want %q", len(values), 2)
	}
}

func Test_Omit(t *testing.T) {
	omitedData := Omit(map[string]interface{}{"name": "Rahul", "age": 32}, []string{"name"})
	if Has(omitedData, "name") {
		t.Errorf("Omit() got= %v, want %v", Has(omitedData, "name"), false)
	}
}

func Test_Pick(t *testing.T) {
	pickedData := Pick(map[string]interface{}{"name": "Rahul", "age": 32}, []string{"age"})
	if Has(pickedData, "name") {
		t.Errorf("Pick() got= %v, want %v", Has(pickedData, "name"), false)
	}
}

func Test_MapValues(t *testing.T) {
	mappedValues := MapValues(map[string]interface{}{"firstName": "john", "lastName": "doe", "gender": "unknown"}, func(value interface{}) interface{} {
		return strings.Title(value.(string))
	})
	if (mappedValues["firstName"].(string) != "John") || (mappedValues["lastName"].(string) != "Doe") {
		t.Errorf("MapValues() got= %v, want %v", mappedValues["firstName"].(string), "John")
		t.Errorf("MapValues() got= %v, want %v", mappedValues["lastName"].(string), "Doe")
	}
}

func Test_MapKeys(t *testing.T) {
	mappedKeys := MapKeys(map[string]interface{}{"firstName": "john", "lastName": "doe", "gender": "unknown"}, func(value interface{}) interface{} {
		return strings.Title(value.(string))
	})
	hasFirstNameAsKey := Contains(StringToInterfaceSlice(Keys(mappedKeys)), "FirstName")
	if !hasFirstNameAsKey {
		t.Errorf("MapKeys() got= %v, want %v", hasFirstNameAsKey, true)
	}
}

func Test_Extend(t *testing.T) {
	extendedData := Extend(
		map[string]interface{}{
			"name": "John",
			"age":  32,
			"contacts": map[string]interface{}{
				"home":  12345,
				"email": "johndoe@gmail.com",
				"address": map[string]interface{}{
					"post_code":    "SW1A",
					"geo_location": []string{"51.529011463529635", "-0.1098365614770662"},
				},
			}},
		map[string]interface{}{
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
		})
	geoLocation := extendedData["contacts"].(map[string]interface{})["address"].(map[string]interface{})["geo_location"]
	postCode := extendedData["contacts"].(map[string]interface{})["address"].(map[string]interface{})["post_code"]
	if geoLocation.([]string)[0] != "51.529011463529636" {
		t.Errorf("Extend() got= %v, want %v", postCode, "SW3")
	}
}

func Test_Get(t *testing.T) {
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
	geoLocation := data["contacts"].(map[string]interface{})["address"].(map[string]interface{})["geo_location"].([]string)[0]
	if geoLocation != geoLocationFromGet {
		t.Errorf("Get() got= %v, want %v", geoLocationFromGet, geoLocation)
	}
}
