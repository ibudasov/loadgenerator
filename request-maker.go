package main

type RequestMaker interface {
	visitPage(frontendBaseUrl string, target UrlConfiguration)
}

type UrlConfiguration struct {
	method                string
	targetPath            string
	numberOfHitsPerTarget int
	body                  string
}

type CollectionOfUrls struct {
	pages []UrlConfiguration
}
