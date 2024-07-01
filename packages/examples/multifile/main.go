//--web true
//--kind go:default

package main

import (
	"fmt"
)

// Main forwarding to Hello
func Main(args map[string]interface{}) map[string]interface{} {
	fmt.Println("Main")
	return Hello(args)
}
