package main

import (
	"fmt"
	"reflect"
)

func help(key string, m map[string]interface{}) (map[string]interface{}, error) {
	if _, ok := (m[key]); ok {
		return m, nil
	}
	for _, val := range m {
		a := reflect.ValueOf(val)
		switch a.Kind() {
		case reflect.Map:
			nested_map := a.Interface().(map[string]interface{})
			//A nil error denotes success; a non-nil error denotes failure.

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
	return nil, fmt.Errorf("Key not found ;)")

}
func remove(key string, m map[string]interface{}) {
	if found, err := help(key, m); err != nil {
		fmt.Println(err)
	} else {
		delete(found, key)
		fmt.Println("The key has been deleted!!")
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
				"plot_no": 27,
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
		"work_ex": map[string]interface{}{
			"Name": []interface{}{"Phonepe", "uber", "Cred"},
		},
	}

	remove("street", m)
	fmt.Println(m)

}
