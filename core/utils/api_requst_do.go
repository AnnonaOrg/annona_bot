package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	API_TIME_OUT_Second = 8
)

func DoPostJsonToOpenAPI(apiDomain, apiPath, apiToken string, req interface{}) (retByte []byte, err error) {
	dataBody := JsonEncode(req)
	apiURL := fmt.Sprintf("%s%s", apiDomain, apiPath)
	// fmt.Println("apiURL", apiURL)
	// fmt.Println("dataBody", dataBody)
	rep, err := http.NewRequest("POST", apiURL, strings.NewReader(dataBody))
	if err != nil {
		return
	}
	rep.Header.Add("Content-Type", "application/json")
	if len(apiToken) > 0 {
		rep.Header.Add("Apiclient", apiToken)
	}

	client := &http.Client{Timeout: time.Second * API_TIME_OUT_Second}

	response, err := client.Do(rep)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	return body, nil
}

func DoPostFormToOpenAPI(apiDomain, apiPath, apiToken string, dataBody string) (retByte []byte, err error) {
	// payload := strings.NewReader("filter=%E5%8F%B2&page=1&size=10")
	// dataBody := util.JsonEncode(req)
	apiURL := fmt.Sprintf("%s%s", apiDomain, apiPath)
	// fmt.Println("apiURL", apiURL)

	rep, err := http.NewRequest("POST", apiURL, strings.NewReader(dataBody))
	if err != nil {
		return
	}
	rep.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if len(apiToken) > 0 {
		rep.Header.Add("Apiclient", apiToken)
	}

	client := &http.Client{Timeout: time.Second * API_TIME_OUT_Second}

	response, err := client.Do(rep)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	return body, nil
}

func DoDeleteToOpenAPI(apiDomain, apiPath, apiToken string) (retByte []byte, err error) {
	// payload := strings.NewReader("filter=%E5%8F%B2&page=1&size=10")
	// dataBody := util.JsonEncode(req)
	apiURL := fmt.Sprintf("%s%s", apiDomain, apiPath)
	// fmt.Println("apiURL", apiURL)

	rep, err := http.NewRequest("DELETE", apiURL, nil)
	if err != nil {
		return
	}
	// rep.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if len(apiToken) > 0 {
		rep.Header.Add("Apiclient", apiToken)
	}

	client := &http.Client{Timeout: time.Second * API_TIME_OUT_Second}

	response, err := client.Do(rep)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	return body, nil
}

func DoPubToOpenAPI(apiDomain, apiPath, apiToken string) (retByte []byte, err error) {
	// payload := strings.NewReader("filter=%E5%8F%B2&page=1&size=10")
	// dataBody := util.JsonEncode(req)
	apiURL := fmt.Sprintf("%s%s", apiDomain, apiPath)
	// fmt.Println("apiURL", apiURL)

	rep, err := http.NewRequest("PUT", apiURL, nil)
	if err != nil {
		return
	}
	// rep.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if len(apiToken) > 0 {
		rep.Header.Add("Apiclient", apiToken)
	}

	client := &http.Client{Timeout: time.Second * API_TIME_OUT_Second}

	response, err := client.Do(rep)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	return body, nil
}
