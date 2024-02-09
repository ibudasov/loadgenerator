package main

type UserInputRetrieverFromMemory struct {
}

func (uic UserInputRetrieverFromMemory) retrieve() LoadGeneratorInput {
	return LoadGeneratorInput{
		numberOfUsers:   1,
		frontendBaseUrl: "http://memohost",
	}
}
