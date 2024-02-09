package main

import "sync"

// todo: improve error handling

func loadGenerator(inputRetriever UserInputRetriever, requestMaker RequestMaker) {
	var scenario = scenarioOfUserInteraction()
	var input = inputRetriever.retrieve()

	var runningScenarios sync.WaitGroup

	for websiteUser := 0; websiteUser < input.numberOfUsers; websiteUser++ {

		runningScenarios.Add(1) // marking the scenario as running

		go func() { // starting parallel execution

			for i := range scenario.pages {
				requestMaker.visitPage(input.frontendBaseUrl, scenario.pages[i])
			}

			waitRandomTimeBetweenWebsiteUsers()

			runningScenarios.Done() // marking the scenario as finished

		}()

		runningScenarios.Wait() // waiting for all the scenarios to complete 👌
	}
}
