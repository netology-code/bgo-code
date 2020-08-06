package main

import (
	"log"
	"regexp"
)

func main() {
	pattern := `^/resources/(?P<rId>\d+)/subresources/(?P<sId>\d+)$`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
		return
	}

	matches := regex.FindStringSubmatch("/resources/10/subresources/20")
	log.Printf("%#v", matches)
	log.Printf("%#v", regex.SubexpNames())
}

