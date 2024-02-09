package main

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("LOADGENERATOR_ENV", "TEST")
}

func Test_whenRequestedToGenerateLoad_thenLoadIsGenerated(t *testing.T) {
	// todo: improve the test with counting amount of requests, etc.
	// So far only important that the application runs
	main()
}
