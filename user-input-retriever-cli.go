package main

import "flag"

type UserInputRetrieverFromCli struct {
}

func (uic UserInputRetrieverFromCli) retrieve() LoadGeneratorInput {
	var numberOfUsers int
	var frontendBaseUrl string

	flag.IntVar(&numberOfUsers, "u", 0, "number of users")
	flag.StringVar(&frontendBaseUrl, "h", "", "frontend base URL")

	flag.Parse()

	// todo: error validation

	return LoadGeneratorInput{
		numberOfUsers:   numberOfUsers,
		frontendBaseUrl: frontendBaseUrl,
	}
}
