package main

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"time"
)

type RequestMakerVegeta struct{}

// todo: error handling is missing here

func (rmv RequestMakerVegeta) visitPage(frontendBaseUrl string, target UrlConfiguration) {
	request := vegeta.NewStaticTargeter(vegeta.Target{
		Method: target.method,
		URL:    frontendBaseUrl,
	})

	rate := vegeta.Rate{Freq: target.numberOfHitsPerTarget, Per: time.Second}

	duration := 1 * time.Second

	vegeta.NewAttacker().Attack(request, rate, duration, "Big Bang!")

	fmt.Println("Requested", frontendBaseUrl, target.targetPath)
}
