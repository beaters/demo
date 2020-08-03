package mypackage

import (
	"fmt"
	anotherpackage "github.com/beaters/utilModule/anotherpackage"
)

// Util4Print 打印输入内容
func Util4Print(input *string) {
	fmt.Printf("input content is %s\n", *input)
	person := anotherpackage.Person{
		Name: "lxg",
		Age: 20,
	}
	fmt.Println(person)
}