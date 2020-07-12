package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := execute("TOKEN_MUST_BE_HERE"); err != nil {
		os.Exit(1)
	}
}

func execute(token string) (err error) {
	reqURL := "https://api.vk.com/method/METHOD_NAME?PARAMETERS&access_token=ACCESS_TOKEN&v=V"
	resp, err := http.Get(reqURL)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(resp.StatusCode)

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

//func execute(token string) (err error) {
//	apiVersion := "5.110"
//
//	reqURL := fmt.Sprintf(
//		"https://api.vk.com/method/messages.send?peer_id=%d&message=%stoken=%s&v=%s",
//		peerId,
//		"Доброго времени суток, специалист скоро подключиться к вашему вопросу!",
//		token,
//		apiVersion,
//	)
//	resp, err := http.Get(reqURL)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//
//	log.Println(resp.StatusCode)
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	defer func() {
//		if cerr := resp.Body.Close(); cerr != nil {
//			if err == nil {
//				log.Println(err)
//				err = cerr
//			}
//		}
//	}()
//
//	log.Println(string(body))
//
//	return nil
//}

//func execute(token string) (err error) {
//	values := make(url.Values)
//	values.Set("access_token", token)
//	values.Set("v", "5.110")
//
//	reqURL := fmt.Sprintf(
//		"https://api.vk.com/method/messages.getConversations?%s",
//		values.Encode(),
//	)
//
//	resp, err := http.Get(reqURL)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	defer func() {
//		if cerr := resp.Body.Close(); cerr != nil {
//			if err == nil {
//				log.Println(err)
//				err = cerr
//			}
//		}
//	}()
//
//	log.Println(string(body))
//
//	return nil
//}

//func execute(token string) (err error) {
//	values := make(url.Values)
//	values.Set("access_token", token)
//	values.Set("v", "5.110")
//
//	reqURL := "https://api.vk.com/method/messages.getConversations"
//
//	// данные будут отправлены не в URL, но VK API обработает
//	resp, err := http.PostForm(reqURL, values)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	defer func() {
//		if cerr := resp.Body.Close(); cerr != nil {
//			if err == nil {
//				log.Println(err)
//				err = cerr
//			}
//		}
//	}()
//
//	log.Println(string(body))
//
//	return nil
//}

//func execute(token string) (err error) {
//	values := make(url.Values)
//	values.Set("access_token", token)
//	values.Set("v", "5.110")
//
//	reqURL := "https://api.vk.com/method/messages.getConversations"
//
//	scratchclient := &http.Client{Timeout: 5 * time.Millisecond}
//
//	// данные будут отправлены не в URL, но VK API обработает
//	resp, err := scratchclient.PostForm(reqURL, values)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	defer func() {
//		if cerr := resp.Body.Close(); cerr != nil {
//			if err == nil {
//				log.Println(err)
//				err = cerr
//			}
//		}
//	}()
//
//	log.Println(string(body))
//
//	return nil
//}

//func execute(token string) (err error) {
//	values := make(url.Values)
//	values.Set("access_token", token)
//	values.Set("v", "5.110")
//
//	reqURL := "https://api.vk.com/method/messages.getConversations"
//
//	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
//	client := &http.Client{}
//
//	req, err := http.NewRequestWithContext(
//		ctx,
//		http.MethodPost,
//		reqURL,
//		strings.NewReader(values.Encode()),
//	)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//
//	// данные будут отправлены не в URL, но VK API обработает
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	defer func() {
//		if cerr := resp.Body.Close(); cerr != nil {
//			if err == nil {
//				log.Println(err)
//				err = cerr
//			}
//		}
//	}()
//
//	log.Println(string(body))
//
//	return nil
//}
