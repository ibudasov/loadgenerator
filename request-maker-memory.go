package main

import (
	"fmt"
)

type RequestMakerMemory struct{}

func (rmv RequestMakerMemory) visitPage(frontendBaseUrl string, target UrlConfiguration) {
	fmt.Println("Requested in memory", frontendBaseUrl, target.targetPath)
}
