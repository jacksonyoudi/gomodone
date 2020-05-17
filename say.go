package gomodone

import "fmt"

// Hi returns a friendly greeting
func SayHi(name string) string {
	return fmt.Sprintf("Hi, %s!", name)
}
