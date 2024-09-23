package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) call() {
	fmt.Println("user method is called\n")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "Aceld", 18}
	DoFieldAndMethod(user)
}

func DoFieldAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("inputtype is :", inputType)

	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	//获取里面的字段和值
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
