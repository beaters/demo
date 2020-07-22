package main

import "fmt"

func main() {
	fmt.Println("")

}

func Hi(name *string) string {
	return fmt.Sprintf("hello, %s", name)
}
