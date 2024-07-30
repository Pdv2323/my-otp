package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 23
	fmt.Println(a)
	b := Temp()
	fmt.Println(b)
	// c := string(a) == b
	// fmt.Println(c)
	t := reflect.TypeOf(b)
	fmt.Println(t)
}

// func Temp() string {
// 	return string(23)
// }

func Temp() int {
	return 23
}
