package main

import (
	"fmt"
	"reflect"
)

const (
	IntType    = "int"
	StringType = "string"
	BoolType   = "bool"
	ChanType   = "chan int"
)

func main() {
	var v interface{}

	v = 42
	checkType(v)
	checkTypeSprintf(v)
	checkTypeSwitch(v)

	v = "Hello"
	checkType(v)
	checkTypeSprintf(v)
	checkTypeSwitch(v)

	v = true
	checkType(v)
	checkTypeSprintf(v)
	checkTypeSwitch(v)

	v = make(chan int)
	checkType(v)
	checkTypeSprintf(v)
	checkTypeSwitch(v)
}

func checkType(v interface{}) {
	// Получаем тип переменной
	typ := reflect.TypeOf(v)

	// Выводим тип переменной
	fmt.Println("Type reflect:", typ)
}

func checkTypeSprintf(v interface{}) {
	typ := fmt.Sprintf("%T", v)

	fmt.Println("Type Sprintf:", typ)
}

func checkTypeSwitch(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Type switch:", IntType)
	case string:
		fmt.Println("Type switch:", StringType)
	case bool:
		fmt.Println("Type switch:", BoolType)
	case chan int:
		fmt.Println("Type switch:", ChanType)
	}
}