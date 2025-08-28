package main

import (
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan any)
	typeIdentifier(12)
	typeIdentifier("something")
	typeIdentifier(true)
	typeIdentifier(ch)
}

func typeIdentifier(something interface{}) {
	if isChannel(something) {
		fmt.Println("Channel")
		return
	}

	switch something.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	}
}

func isChannel(something interface{}) bool {
	return reflect.TypeOf(something).Kind() == reflect.Chan
}
