package demo

import "fmt"

func Hi(name *string) string {
	return fmt.Sprintf("hello, %s", name)
}
