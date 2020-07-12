package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
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
		log.Println("no vk version specified (VK_VERSION env variable)")
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
	values := make(url.Values)
	values.Set("access_token", token)
	values.Set("v", version)

	reqURL := fmt.Sprintf("%s/%s", baseURL, "messages.getConversations")

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout) * time.Second)
	client := &http.Client{}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		reqURL,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// данные будут отправлены не в URL, но VK API обработает
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			if err == nil {
				log.Println(err)
				err = cerr
			}
		}
	}()

	log.Println(string(body))

	return nil
}


