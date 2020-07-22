package demo

import "fmt"

/*func Hi(name *string) string {
	return fmt.Sprintf("hello, %s", *name)
}*/

func Hi(name *string) (string, uint32) {
	return fmt.Sprintf("hello, %s", *name), nil
}
