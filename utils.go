/*

Copyright (c) 2018-2019 Rashed Mohammed, The TurtleCoin Developers

Please see the included LICENSE file for more information

*/

package turtlecoinrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	return decodeResponse(req)
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

	req, err1 := http.NewRequest("POST", "http://"+daemon.URL+":"+strconv.Itoa(daemon.Port)+"/json_rpc", body)
	if err1 != nil {
		fmt.Println(err1)
		return nil
	}

	return decodeResponse(req)
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

	req, err1 := http.NewRequest("POST", "http://"+wallet.URL+":"+strconv.Itoa(wallet.Port)+"/json_rpc", body)
	if err1 != nil {
		fmt.Println(err1)
		return nil
	}

	return decodeResponse(req)
}

func decodeResponse(req *http.Request) map[string]interface{} {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
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

	fmt.Println(resp.StatusCode, resp.Status)
	return nil
}
