package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

func init() {
	log.SetPrefix("Reflect package:")
}

type person struct {
	Name string
	Age  int
}

func main() {
	log.Println("Started ............")
	p := person{Name: "tony", Age: 27}
	res, err := convertStructToMap(p)
	if err != nil {
		log.Println("failed to convert struct to map")
	}
	fmt.Printf("%s\n", strings.Repeat("/", 50))
	fmt.Printf("The resulting map is %v\n", res)
	fmt.Printf("%s\n", strings.Repeat("/", 50))
	person1 := convertMapToStruct(res)
	fmt.Printf("The resulting map is %+v\n", person1)
	fmt.Printf("%s\n", strings.Repeat("/", 50))

}
func convertStructToMap(v any) (map[string]any, error) {
	refObjValue := reflect.ValueOf(v)
	refObjType := reflect.TypeOf(v)
	if refObjValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("kind is %v not of supported", refObjValue.Kind())
	}
	arr := make(map[string]any)
	for i := 0; i < refObjValue.NumField(); i++ {
		structFieldObjtype := refObjType.Field(i).Name
		structValueObj := refObjValue.Field(i).Interface()
		arr[structFieldObjtype] = structValueObj
	}
	return arr, nil
}
func convertMapToStruct(ar map[string]any) person {
	var p person
	for k, v := range ar {
		field := reflect.ValueOf(&p).Elem().FieldByName(k)
		if field.IsValid() && field.CanSet() {
			val := reflect.ValueOf(v)
			if val.Type().AssignableTo(field.Type()) {
				field.Set(val)
			}
		}
	}
	return p
}
func mapToStruct(objMap map[string]interface{}) person {
	person := person{
		Name: objMap["Name"].(string),
		Age:  objMap["Age"].(int),
	}

	return person
}
