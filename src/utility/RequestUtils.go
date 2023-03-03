package utility

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/RaRa-Delivery/rara-ms-models/src/utility/lg"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Post() {
	fmt.Println("2. Performing Http Post...")
	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, _ := json.Marshal(todo)
	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("%+v\n", todoStruct)
}

func ApiResponseUtility(requestMethod string, url string, headers map[string]string, payload string, reqId string, desc string) ([]byte, error) {
	log.Println(reqId+": ", desc, " ", "URL:", url)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	var str io.Reader
	if payload == "no" {
		str = nil
	} else {
		str = strings.NewReader(payload)
	}

	req, err := http.NewRequest(requestMethod, url, str)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	log.Println(reqId+": ", desc)
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	log.Println(reqId+": ", desc, " ", bodyString)
	log.Println("response.StatusCode: ", response.StatusCode)
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return bodyBytes, nil
	} else {

		return nil, errors.New(bodyString)

	}
}

func GetApiResponse(url string, headers map[string]string) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	head, _ := json.Marshal(&headers)

	log.Println("headers: ", string(head))
	log.Println("URL: ", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	for k, v := range headers {
		log.Println(k, " : ", v)
		req.Header.Add(k, v)
	}

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	log.Println("response.StatusCode: ", response.StatusCode)
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return bodyString, nil
	} else {

		return bodyString, errors.New(bodyString)

	}
}

func PostApiResponse(url string, headers http.Header, payload string) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	str := strings.NewReader(payload)
	req, err := http.NewRequest("POST", url, str)
	if err != nil {
		return "", err
	}

	req.Header = headers

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	log.Println("response.StatusCode: ", response.StatusCode)
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return bodyString, nil
	} else {

		return bodyString, errors.New(bodyString)

	}
}

func ApiResponse(requestMethod string, url string, headers map[string]string, payload string) (string, error) {

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	var str io.Reader
	if payload == "no" {
		str = nil
	} else {
		log.Println("payload:", payload)
		str = strings.NewReader(payload)
	}
	log.Println("requestMethod:", requestMethod)
	log.Println("URL:", url)
	log.Println("headers:", headers)
	req, err := http.NewRequest(requestMethod, url, str)
	if err != nil {
		return "", err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	//log.Println(response.Body)
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	log.Println("response.StatusCode: ", response.StatusCode)
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return bodyString, nil
	} else {

		return bodyString, errors.New(bodyString)

	}
	//	log.Println(bodyString)

}

func HandlePanic() {

	if e := recover(); e != nil {
		err := fmt.Errorf("recovered from panic: %v", e)
		log.Println(lg.Error("RECOVER_FROM_PANIC", err))
	}
}
