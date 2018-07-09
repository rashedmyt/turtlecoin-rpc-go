package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func MakeGetRequest(method string, hostURL string, hostPort int) *bytes.Buffer {
	req, err := http.NewRequest("GET", "http://"+hostURL+":"+strconv.Itoa(hostPort)+"/"+method, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return decodeResponse(req)
}

func MakePostRequest(hostURL string, hostPort int, method string, params interface{}) *bytes.Buffer {
	payload := make(map[string]interface{})
	payload["jsonrpc"] = "2.0"
	payload["method"] = method
	payload["params"] = params

	jsonpayload, err := json.Marshal(payload)

	body := bytes.NewBuffer(jsonpayload)

	req, err := http.NewRequest("POST", "http://"+hostURL+":"+strconv.Itoa(hostPort)+"/json_rpc", body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return decodeResponse(req)
}

func MakeWalletPostRequest(rpcPassword string, hostURL string, hostPort int, method string, params interface{}) *bytes.Buffer {
	payload := make(map[string]interface{})
	payload["jsonrpc"] = "2.0"
	payload["id"] = 1
	payload["password"] = rpcPassword
	payload["method"] = method
	payload["params"] = params

	jsonpayload, err := json.Marshal(payload)

	body := bytes.NewBuffer(jsonpayload)

	req, err := http.NewRequest("POST", "http://"+hostURL+":"+strconv.Itoa(hostPort)+"/json_rpc", body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return decodeResponse(req)
}

func decodeResponse(req *http.Request) *bytes.Buffer {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	return bytes.NewBuffer(responseBody)
}
