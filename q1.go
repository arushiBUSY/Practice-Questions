package main

import (
	"fmt"
	"reflect"
)

func help(key string, m map[string]interface{}) (map[string]interface{}, error) {
	//if a key exists in a map
	//this will also take care of default case of switch

	if _, ok := m[key]; ok {
		return m, nil
	}
	for _, val := range m {
		a := reflect.ValueOf(val)
		//If the kind is reflect.Map, it extracts the nested map and recursively calls itself.
		//If the kind is reflect.Slice, it iterates through the slice, and if an element is a map, it recursively calls itself on that nested map.

		switch a.Kind() {
		case reflect.Map:
			//            extracting map value,Interface() is a method used on reflection object "a"

			nested_map := a.Interface().(map[string]interface{})
			if found, err := help(key, nested_map); err == nil {
				return found, nil
			}
		case reflect.Slice:
			nested_slice := a.Interface().([]interface{})

			for _, value := range nested_slice {
				a := reflect.ValueOf(value)
				if a.Kind() == reflect.Map {
					nested_map_slice := a.Interface().(map[string]interface{})
					if found, err := help(key, nested_map_slice); err == nil {
						return found, nil
					}
				}
			}

		}

	}
	//It is used to create and return a new error value formatted according to a format specifier.
	/// Using fmt.Errorf to create an error with a formatted string

	return nil, fmt.Errorf("key not found")
}
func update(key string, m map[string]interface{}, value interface{}) {
	if found_val, err := help(key, m); err != nil {
		fmt.Println(err)
	} else {
		found_val[key] = value
		fmt.Println("VALUE UPDATED !!")

	}
}
func main() {
	var m = map[string]interface{}{
		"Name": "Arushi Sharma",
		"DOB":  24 - 01 - 2002,
		"city": "Delhi",
		"pin":  110075,
		// field named "Address" and assigns it a slice ([]interface{}). The square brackets [] indicate that it's a slice, and interface{} allows elements of any data type to be stored in the slice.

		"Address": []interface{}{
			//nside the slice, there are two elements, each represented by a map.

			map[string]interface{}{
				"street":  "Ashirvad chowk",
				"plot_no": 26,
				"city":    "Dwarka",
				"pin":     110078,
			},
			map[string]interface{}{
				"street":  "Lovely Chowk",
				"plot_no": 26,
				"city":    "Dwarka",
				"pin":     110078,
			},
		},
		"Salary":      100000,
		"Designation": "Developer",
	}
	//
	var val interface{}
	val = "UP"

	update("city", m, val)
	fmt.Println(m)

}
