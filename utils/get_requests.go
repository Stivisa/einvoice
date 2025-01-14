package utils

import (
	"io"
	"log"
	"net/http"
)

// http get with query params
func GetQ(url string, params ...string) ([]byte, int, string) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

// http get with path params
func GetP(url string, params ...string) ([]byte, int, string) {

	//http request path parameters
	for i := 0; i < len(params); i += 1 {
		url = url + "/" + params[i]
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("ApiKey", ApiKey)

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

func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
