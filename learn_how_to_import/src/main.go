package main

import (
	"fmt"
	"learn_how_to_import/simple_import_test"
)

func main() {
	fmt.Println("This is the main package")
	fmt.Println(simple_import_test.SimpleImportTest())
	fmt.Println(simple_import_test.ThisIsAfunctionThatStartsWithLowerCase())
}
