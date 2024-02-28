package main

import (
	"fmt"
	"reflect"
)

func PoplateStruct(data map[string]interface{}, res interface{}) {
	//.Elem(): This method call dereferences the pointer represented by the reflect.Value
	//.Elem() returns the Value that the pointer points to

	resVal := reflect.ValueOf(res).Elem()
	for key, value := range data {
		//This method of the reflect.Value type is used to retrieve a field of the struct by its name. It takes a string parameter (key) representing the name of the field to be retrieved.

		field := resVal.FieldByName(key)

		if field.IsValid() {
			if field.Kind() == reflect.Struct {
				//, assuming it is a map with string keys and interface{} values. If successful, it assigns the value of value to nestedMap, and ok will be true

				if nested_map, ok := value.(map[string]interface{}); ok {
					//This line creates a new instance of the struct type corresponding to the field using reflection. field.Type() retrieves the type of the field, and reflect.New() creates a new instance of that type. .Interface() converts the reflect.Value representing the new struct instance to an interface{}, allowing it to be passed as an argument to PopulateStruct.

					nested_struct := reflect.New(field.Type()).Interface()
					PoplateStruct(nested_map, nested_struct)
					field.Set(reflect.ValueOf(nested_struct).Elem())
				} else {
					field.Set(reflect.ValueOf(value))
				}

			}
		}
	}

}

type Person struct {
	Name    string
	Age     int
	Address Address
}
type Address struct {
	City  string
	State string
}

func main() {
	data := map[string]interface{}{
		"Name": "Arushi",
		"Age":  22,
		"Pin":  110075,
		"Address": map[string]interface{}{
			//n the Address struct, I've capitalized the field names City and State, making them exported and accessible from outside the package.

			"City":  "Dwarka",
			"State": "Delhi",
		},
	}
	//This line declares a variable named personPtr and initializes it with the memory address of a newly created Person struct

	var person_ptr *Person = &Person{}
	PoplateStruct(data, person_ptr)
	// This is a format string specifying how the data should be printed. Here, %+v is a formatting verb that tells Printf to print the struct fields and their values with field names. The + flag tells Printf to also include field names in the output.

	fmt.Printf("%+v\n", *person_ptr)

}
