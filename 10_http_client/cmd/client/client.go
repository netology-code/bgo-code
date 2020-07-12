package main

import (
	"lectionhttpclient/pkg/vk"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	vkToken, ok := os.LookupEnv("VK_TOKEN")
	if !ok {
		log.Println("no vk token specified (VK_TOKEN env variable)")
		os.Exit(1)
	}

	vkBaseURL, ok := os.LookupEnv("VK_URL")
	if !ok {
		log.Println("no base vk url specified (VK_URL env variable)")
		os.Exit(1)
	}

	vkVersion, ok := os.LookupEnv("VK_VERSION")
	if !ok {
		log.Println("no base vk version specified (VK_VERSION env variable)")
		os.Exit(1)
	}

	vkTimeout, ok := os.LookupEnv("VK_TIMEOUT")
	if !ok {
		log.Println("no vk timeout specified (VK_TIMEOUT env variable)")
		os.Exit(1)
	}

	timeout, err := strconv.Atoi(vkTimeout)
	if err != nil {
		log.Println("bad vk timeout value")
		os.Exit(1)
	}

	if err := execute(vkToken, vkBaseURL, vkVersion, timeout); err != nil {
		os.Exit(1)
	}
}

func execute(token, baseURL, version string, timeout int) (err error) {
	svc := vk.NewService(
		token,
		baseURL,
		version,
		time.Duration(timeout),
		&http.Client{},
	)
	conversations, err := svc.Conversations()
	if err != nil {
		return err
	}
	// TODO: work with conversation
	log.Println(conversations.Response.Count)
	return nil
}


