package test

// HandleError for error handling in test suite
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
