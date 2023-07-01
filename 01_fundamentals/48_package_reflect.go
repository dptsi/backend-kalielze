package main

import (
	"fmt"
	"reflect"
)

type Sport struct {
	Name string	`required:"true" min:"5"`
	Rule string 
}

func main() {

	sport := Sport{
		Name: "Football",
	}
	sportType := reflect.TypeOf(sport)
	fmt.Println(sportType)
	fmt.Println(sportType.NumField())
	fmt.Println(sportType.Field(0))

	structField := sportType.Field(0)
	fmt.Println("FieldName:", structField.Name)
	fmt.Println("FiledType:", structField.Type.Name())

	required := sportType.Field(1).Tag.Get("required")
	fmt.Println(required)

	fmt.Println(IsValid(sport))
	/*
		masih banyak documentation mengenai package reflect

		https://golang.org/pkg/reflect/
	*/
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			val := reflect.ValueOf(data).Field(i).Interface()
			if val == "" {
				return false
			}
		}
	}
	return true
}
