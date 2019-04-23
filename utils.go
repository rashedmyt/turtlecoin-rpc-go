// Copyright (c) 2018-2019 Rashed Mohammed, The TurtleCoin Developers
// Please see the included LICENSE file for more information

package turtlecoinrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (daemon *TurtleCoind) makeGetRequest(method string) map[string]interface{} {
	req, err := http.NewRequest("GET", "http://"+daemon.URL+":"+strconv.Itoa(daemon.Port)+"/"+method, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	resp := performRequest(req)

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode, resp.Status)
		return nil
	}

	return decodeResponse(resp.Body)
}

func (daemon *TurtleCoind) makePostRequest(method string, params interface{}) map[string]interface{} {
	payload := make(map[string]interface{})
	payload["jsonrpc"] = "2.0"
	payload["method"] = method
	payload["params"] = params

	jsonpayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	body := bytes.NewBuffer(jsonpayload)

	req, err := http.NewRequest("POST", "http://"+daemon.URL+":"+strconv.Itoa(daemon.Port)+"/json_rpc", body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	resp := performRequest(req)

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode, resp.Status)
		return nil
	}

	return decodeResponse(resp.Body)
}

func (wallet *Walletd) makePostRequest(method string, params interface{}) map[string]interface{} {
	payload := make(map[string]interface{})
	payload["jsonrpc"] = "2.0"
	payload["id"] = 1
	payload["password"] = wallet.RPCPassword
	payload["method"] = method
	payload["params"] = params

	jsonpayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	body := bytes.NewBuffer(jsonpayload)

	req, err := http.NewRequest("POST", "http://"+wallet.URL+":"+strconv.Itoa(wallet.Port)+"/json_rpc", body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	resp := performRequest(req)

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode, resp.Status)
		return nil
	}

	return decodeResponse(resp.Body)
}

func (wallet *WalletAPI) makeGetRequest(method string) map[string]interface{} {
	req, err := http.NewRequest("GET", "http://"+wallet.URL+":"+strconv.Itoa(wallet.Port)+"/"+method, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Set("X-API-KEY", wallet.RPCPassword)
	resp := performRequest(req)

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode, resp.Status)
		return nil
	}

	return decodeResponse(resp.Body)
}

func (wallet *WalletAPI) makeDeleteRequest(method string) map[string]interface{} {
	req, err := http.NewRequest("DELETE", "http://"+wallet.URL+":"+strconv.Itoa(wallet.Port)+"/"+method, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Set("X-API-KEY", wallet.RPCPassword)
	resp := performRequest(req)

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode, resp.Status)
		return nil
	}

	return decodeResponse(resp.Body)
}

func performRequest(req *http.Request) *http.Response {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return resp
}

func decodeResponse(body io.ReadCloser) map[string]interface{} {
	defer body.Close()

	respBody, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var respBodyInterface interface{}
	if err = json.Unmarshal(respBody, &respBodyInterface); err != nil {
		fmt.Println(err)
		return nil
	}

	return respBodyInterface.(map[string]interface{})
}
