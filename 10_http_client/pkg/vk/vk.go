package vk

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lectionhttpclient/pkg/vk/dto/conversations"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Service struct {
	token   string
	baseURL string
	version string
	timeout time.Duration
	client  *http.Client
}

func NewService(
	token string,
	baseURL string,
	version string,
	timeout time.Duration,
	client *http.Client,
) *Service {
	return &Service{
		// только для презентации
		token: token, baseURL: baseURL, version: version,
		timeout: timeout * time.Second,
		client:  client,
	}
}

const (
	formURLEncoded = "application/x-www-form-urlencoded"
	multipart      = "multipart/form-data"
)

const (
	headerContentType = "Content-Type"
)

type apiMethod string

const (
	conversationsMethod apiMethod = "messages.getConversations"
	sendMessageMethod   apiMethod = "messages.send"
)

func (s *Service) Conversations() (*conversations.ResponseDTO, error) {
	values := make(url.Values)
	// data - []byte
	data, err := s.execute(conversationsMethod, values)
	if err != nil {
		return nil, err
	}

	var response *conversations.ResponseDTO

	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *Service) execute(method apiMethod, values url.Values) (data []byte, err error) {
	values.Set("access_token", s.token)
	values.Set("v", s.version)

	reqURL := fmt.Sprintf("%s/%s", s.baseURL, method)

	ctx, _ := context.WithTimeout(context.Background(), s.timeout)
	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, // перенесли только для презентации
		reqURL,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set(headerContentType, formURLEncoded)

	// данные будут отправлены не в URL, но VK API обработает
	resp, err := s.client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			if err == nil {
				log.Println(err)
				err = cerr
			}
		}
	}()

	return respBody, nil
}
