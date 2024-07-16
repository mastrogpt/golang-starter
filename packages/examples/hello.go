//--kind go:default
//--web true

package main

import "hello" // TODO: go init mod e importare le dipendenze nella sottocartella

func Main(obj map[string]any) map[string]any {
	return hello.HelloWorld(obj)
}
