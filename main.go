package main

import (
	"fmt"

	"github.com/deveshdama/mydummypackage"
)

func main() {
	data, err := mydummypackage.LoadData()
	if err != nil {
		fmt.Printf("failed to load data: %v\n", err)
		return
	}
	fmt.Printf("Foo: %s\nBar: %s\nBaz: %s\n", data.Foo, data.Bar, data.Baz)
}
