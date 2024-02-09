package main

type LoadGeneratorInput struct {
	numberOfUsers   int
	frontendBaseUrl string
}

type UserInputRetriever interface {
	retrieve() LoadGeneratorInput
}
