// Copyright (c) 2018-2019 Rashed Mohammed, The TurtleCoin Developers
// Please see the included LICENSE file for more information

package turtlecoinrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (daemon *TurtleCoind) makeGetRequest(method string) (interface{}, error) {
	req, err := http.NewRequest("GET", "http://"+daemon.URL+":"+strconv.Itoa(daemon.Port)+"/"+method, nil)
	if err != nil {
		return nil, err
	}

	resp, err := performRequest(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(strconv.Itoa(resp.StatusCode) + " " + resp.Status)
	}

	return decodeResponse(resp.Body)
}

func (daemon *TurtleCoind) makePostRequest(method string, params interface{}) (interface{}, error) {
	payload := make(map[string]interface{})
	payload["jsonrpc"] = "2.0"
	payload["method"] = method
	payload["params"] = params

	jsonpayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(jsonpayload)

	req, err := http.NewRequest("POST", "http://"+daemon.URL+":"+strconv.Itoa(daemon.Port)+"/json_rpc", body)
	if err != nil {
		return nil, err
	}

	resp, err := performRequest(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(strconv.Itoa(resp.StatusCode) + " " + resp.Status)
	}

	return decodeResponse(resp.Body)
}

func (wallet *Walletd) makePostRequest(method string, params interface{}) (interface{}, error) {
	payload := make(map[string]interface{})
	payload["jsonrpc"] = "2.0"
	payload["id"] = 1
	payload["password"] = wallet.RPCPassword
	payload["method"] = method
	payload["params"] = params

	jsonpayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(jsonpayload)

	req, err := http.NewRequest("POST", "http://"+wallet.URL+":"+strconv.Itoa(wallet.Port)+"/json_rpc", body)
	if err != nil {
		return nil, err
	}

	resp, err := performRequest(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(strconv.Itoa(resp.StatusCode) + " " + resp.Status)
	}

	return decodeResponse(resp.Body)
}

func (wallet *WalletAPI) makeGetRequest(method string) (interface{}, error) {
	req, err := http.NewRequest("GET", wallet.checkSSL()+"://"+wallet.URL+":"+strconv.Itoa(wallet.Port)+"/"+method, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", wallet.RPCPassword)
	resp, err := performRequest(req)
	if err != nil {
		return nil, err
	}

	err = handleResponseStatusCode(resp)
	if err != nil {
		return nil, err
	}

	return decodeResponse(resp.Body)
}

func (wallet *WalletAPI) makeDeleteRequest(method string) (interface{}, error) {
	req, err := http.NewRequest("DELETE", wallet.checkSSL()+"://"+wallet.URL+":"+strconv.Itoa(wallet.Port)+"/"+method, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", wallet.RPCPassword)
	resp, err := performRequest(req)
	if err != nil {
		return nil, err
	}

	err = handleResponseStatusCode(resp)
	if err != nil {
		return nil, err
	}

	return decodeResponse(resp.Body)
}

func (wallet *WalletAPI) makePutRequest(method string, params map[string]interface{}) (interface{}, error) {
	jsonBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(jsonBody)

	req, err := http.NewRequest("PUT", wallet.checkSSL()+"://"+wallet.URL+":"+strconv.Itoa(wallet.Port)+"/"+method, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", wallet.RPCPassword)
	resp, err := performRequest(req)
	if err != nil {
		return nil, err
	}

	err = handleResponseStatusCode(resp)
	if err != nil {
		return nil, err
	}

	return decodeResponse(resp.Body)
}

func (wallet *WalletAPI) makePostRequest(method string, params map[string]interface{}) (interface{}, error) {
	jsonBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(jsonBody)

	req, err := http.NewRequest("POST", wallet.checkSSL()+"://"+wallet.URL+":"+strconv.Itoa(wallet.Port)+"/"+method, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", wallet.RPCPassword)
	resp, err := performRequest(req)
	if err != nil {
		return nil, err
	}

	err = handleResponseStatusCode(resp)
	if err != nil {
		return nil, err
	}

	return decodeResponse(resp.Body)
}

func (wallet *WalletAPI) checkSSL() string {
	if wallet.DaemonSSL {
		return "https"
	}

	return "http"
}

func performRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func handleResponseStatusCode(resp *http.Response) error {
	if resp.StatusCode == http.StatusBadRequest {
		responseError, err := decodeResponse(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(responseError.(map[string]interface{})["errorMessage"].(string))
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted:
		return nil
	case http.StatusUnauthorized:
		return errors.New("API key is missing or invalid")
	case http.StatusForbidden:
		return errors.New("A wallet is already open. Call DELETE on /wallet first, to close it")
	case http.StatusNotFound:
		return errors.New("The transaction hash was not found")
	case http.StatusInternalServerError:
		return errors.New("An exception was thrown whilst processing the request. See the WalletAPI console for logs")
	}

	return errors.New("Unhandled Status " + strconv.Itoa(resp.StatusCode))
}

func decodeResponse(body io.ReadCloser) (interface{}, error) {
	defer body.Close()

	respBody, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	if len(respBody) == 0 {
		return nil, nil
	}

	var respBodyInterface interface{}
	if err = json.Unmarshal(respBody, &respBodyInterface); err != nil {
		return nil, err
	}

	return respBodyInterface, nil
}

// PrettyPrint prints the given map
// as a JSON object
func PrettyPrint(response interface{}) {
	resp, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		println(err)
		return
	}

	println(string(resp))
}
