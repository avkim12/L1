package main

import (
	"fmt"
	"reflect"
)

func printfTFlag(v interface{}) {
	fmt.Printf("%T\n", v)
}

func typeSwitch(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Println(t)
	case string:
		fmt.Println(t)
	case bool:
		fmt.Println(t)
	case chan int:
		fmt.Println(t)
	default:
		fmt.Println("Unknown type")
	}
}

func reflectTypeOf(v interface{}) {
	fmt.Println(reflect.TypeOf(v))
}

func reflectValueOfKind(v interface{}) {
	fmt.Println(reflect.ValueOf(v).Kind())
}

func main() {
	printfTFlag(0)
	typeSwitch("string")
	reflectTypeOf(true)
	reflectValueOfKind(make(chan int))
}
