package main

import (
	"fmt"
	"math/rand"
	"time"
)

func scenarioOfUserInteraction() CollectionOfUrls {

	var collectionOfUrls CollectionOfUrls

	// products := []string{
	// 	"0PUK6V6EV0",
	// 	"1YMWWN1N4O",
	// 	"2ZYFJ3GM2N",
	// 	"66VCHSJNUP",
	// 	"6E92ZMYYFZ",
	// 	"9SIQT8TOJO",
	// 	"L9ECAV7KIM",
	// 	"LS4PSXUNUM",
	// 	"OLJCESPC7Z",
	// }

	// index
	collectionOfUrls.pages = append(collectionOfUrls.pages, UrlConfiguration{
		method:                "GET",
		targetPath:            "/",
		numberOfHitsPerTarget: 1,
		body:                  "",
	})

	// setCurrency
	collectionOfUrls.pages = append(collectionOfUrls.pages, UrlConfiguration{
		method:                "POST",
		targetPath:            "/setCurrency",
		numberOfHitsPerTarget: 1,
		body:                  "{'currency_code': random.choice(currencies)}",
	})

	// browse products
	// collectionOfUrls.pages = append(collectionOfUrls.pages, UrlConfiguration{
	// 	method:                "GET",
	// 	targetPath:            "/product/" + chooseRandom(products),
	// 	numberOfHitsPerTarget: 1,
	// 	body:                  "",
	// })

	// view cart
	collectionOfUrls.pages = append(collectionOfUrls.pages, UrlConfiguration{
		method:                "GET",
		targetPath:            "/cart",
		numberOfHitsPerTarget: 1,
		body:                  "",
	})

	// todo: add other pages

	return collectionOfUrls
}

func waitRandomTimeBetweenWebsiteUsers() {
	var randomWaitingTimeBetweenUsers = rand.Intn(10) + 1
	time.Sleep(time.Duration(randomWaitingTimeBetweenUsers) * time.Second)
	fmt.Println(randomWaitingTimeBetweenUsers, "seconds waited")
}

func chooseRandom(choices []string) string {
	randomIndex := rand.Intn(len(choices))

	return choices[randomIndex]
}
