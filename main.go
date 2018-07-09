package main

import (
	"fmt"
)

func init() {
	println("init")
}
func main() {
	fmt.Println("Hello")
	var arr = []string{
		"1",
		"2",
	}
	fmt.Print(arr)

	slice := []string{}
	slice = append(slice, "AA")
	println(slice)
	// key type string and value type string
	mapVal := map[string]string{
		"key": "value",
	}
	println(mapVal)
	// key type string and value type interface = anytype
	mapVal2 := map[string]interface{}{
		"sss": 1,
		"nestedMap": map[string]interface{}{
			"key": 2,
		},
	}
	println(mapVal2)

	// for i := 1; i < 11; i++ {
	// 	println(i)
	// }
	Hello("ssadasdasdkaslkdas")
	Hellos("ssadasdasdkaslkdas", "surmane")
	MyPrintLn("a", "b", "c")
	fmt.Print(Return())
}

func Hello(name string) {
	println("My name is" + name)
}

func Hellos(name, surmane string) {
	println("My name is"+name, surmane)
}

func MyPrintLn(s ...interface{}) {
	for _, e := range s {
		fmt.Print(e)
	}
}

func Return() int {
	return 1
}
