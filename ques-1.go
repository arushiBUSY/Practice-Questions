package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func print_type(m map[string]interface{}) {
	for key, value := range m {
		keyType := reflect.TypeOf(key)
		valueType := reflect.ValueOf(value)
		switch valueType.Kind() {
		case reflect.Map:
			nested_map := valueType.Interface().(map[string]interface{})

			fmt.Printf("Key is-> %v and its Type: %v ,Value is %v and its Type: %v\n", key, keyType, value, reflect.TypeOf(value))
			fmt.Printf("Now checking type in nested map also\n")
			fmt.Printf("\n")
			print_type(nested_map)

		case reflect.Slice:
			nested_slice := valueType.Interface().([]interface{})
			fmt.Printf("Key is-> %v and its Type: %v ,Value is %v and its Type: %v\n", key, keyType, value, reflect.TypeOf(value))
			fmt.Printf("Now checking type whether slice is nested with some data structure also\n")
			fmt.Printf("\n")
			for _, val := range nested_slice {
				a := reflect.ValueOf(val)
				if a.Kind() == reflect.Map {
					nested_map_slice := a.Interface().(map[string]interface{})
					print_type(nested_map_slice)
				}

			}
		default:
			fmt.Printf("Key is-> %v and its Type: %v ,Value is %v and its Type: %v\n", key, keyType, value, reflect.TypeOf(value))

		}

	}
}
func main() {
	var input = `{
        "name": "Tolexo Online Pvt. Ltd",
        "age_in_years": 8.5,
        "origin": "Noida",
        "head_office": "Noida, Uttar Pradesh",
        "address": [
            {
                "street": "91 Springboard",
                "landmark": "Axis Bank",
                "city": "Noida",
                "pincode": 201301,
                "state": "Uttar Pradesh"
            },
            {
                "street": "91 Springboard",
                "landmark": "Axis Bank",
                "city": "Noida",
                "pincode": 201301,
                "state": "Uttar Pradesh"
            }
        ],
        "sponsors": {
            "name": "One"
        },
        "revenue": "19.8 million$",
        "no_of_employee": 630,
        "str_text": ["one", "two"],
        "int_text": [1, 3, 4]
    }`
	m := make(map[string]interface{})
	//converting "input" variable  as byte slice
	//'&m' passes pointer to a map variable 'm'
	//the unmarshal func will populate this map with decoded JSON data
	err := json.Unmarshal([]byte(input), &m)
	if err != nil {
		return
	}

	print_type(m)

}
