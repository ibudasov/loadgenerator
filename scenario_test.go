package main

import (
	"os"
	"reflect"
	"testing"
)

func init() {
	os.Setenv("LOADGENERATOR_ENV", "TEST")
}

func Test_givenFewPagesConfigured_whenRequested_thenTheyAreReturned(t *testing.T) {
	tests := []struct {
		name string
		want CollectionOfUrls
	}{
		{
			"",
			expectedCollectionOfUrls(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scenarioOfUserInteraction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scenarioOfUserInteraction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func expectedCollectionOfUrls() CollectionOfUrls {
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

	var collectionOfUrls CollectionOfUrls

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

	// view cart
	collectionOfUrls.pages = append(collectionOfUrls.pages, UrlConfiguration{
		method:                "GET",
		targetPath:            "/cart",
		numberOfHitsPerTarget: 1,
		body:                  "",
	})

	return collectionOfUrls
}
