package main

import "os"

func main() {

	// @see: src/loadgenerator/load-generator_test.go:8
	if os.Getenv("LOADGENERATOR_ENV") == "TEST" {

		requestMaker := RequestMakerMemory{}
		inputRetriever := UserInputRetrieverFromMemory{}
		loadGenerator(inputRetriever, requestMaker)

	} else {

		requestMaker := RequestMakerVegeta{}
		inputRetriever := UserInputRetrieverFromCli{}
		loadGenerator(inputRetriever, requestMaker)

	}
}
