package simple_import_test

func SimpleImportTest() string {
	return "Simple import test"
}

// It's not possible to import a function that starts with a lower case letter
// from another package. This is because it's not exported.
func ThisIsAfunctionThatStartsWithLowerCase() string {
	return "Another simple test"
}
