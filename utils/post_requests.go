package utils

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func PostQ(url string, params ...string) ([]byte, int, string) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//http request query parameters
	q := req.URL.Query()
	for i := 0; i < len(params); i += 2 {
		q.Add(params[i], params[i+1])
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Add("ApiKey", ApiKey)

	//log.Println(req.URL)

	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body, response.StatusCode, response.Status
}

// http get with query params and xml
func PostQXml(url string, bodyContent string, params ...string) ([]byte, int, string) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(bodyContent)))
	if err != nil {
		log.Fatalln(err)
	}

	//http request query parameters
	q := req.URL.Query()
	for i := 0; i < len(params); i += 2 {
		q.Add(params[i], params[i+1])
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Add("ApiKey", ApiKey)
	req.Header.Add("Content-Type", "application/xml")

	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body, response.StatusCode, response.Status
}

func PostQJson(url string, bodyContent string, params ...string) ([]byte, int, string) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(bodyContent)))
	if err != nil {
		log.Fatalln(err)
	}

	//http request query parameters
	q := req.URL.Query()
	for i := 0; i < len(params); i += 2 {
		q.Add(params[i], params[i+1])
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Add("ApiKey", ApiKey)
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body, response.StatusCode, response.Status
}
