package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	aInt := 5
	aType := reflect.TypeOf(aInt)
	fmt.Printf("type %s | kind % s \n", aType.Name(), aType.Kind())
}
