package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string "namestr"
	age int
}

func ShowTag(i interface{}) {
	switch tt := i.(type) {
	case *Person:
		println(tt)
		t := reflect.TypeOf(i)
		v := reflect.ValueOf(i)
		tag := t.Elem().Field(0).Tag
		name := v.Elem().Field(0).String()
		fmt.Printf("%v\n", tag)
		fmt.Printf("%v\n", name)
	}
}

func main() {
	p := Person{"123", 0}
	ShowTag(&p)
}
