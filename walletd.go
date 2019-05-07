// Copyright (c) 2018-2019 Rashed Mohammed, The TurtleCoin Developers
// Please see the included LICENSE file for more information

package turtlecoinrpc

import (
	"errors"
)

// Walletd structure contains the URL and Port info of
// the wallet service and RPC Password for RPC calls
type Walletd struct {
	URL         string
	Port        int
	RPCPassword string
}

func (wallet *Walletd) check() error {
	if wallet.URL == "" {
		wallet.URL = "127.0.0.1"
	}
	if wallet.Port == 0 {
		wallet.Port = 8070
	}
	if wallet.RPCPassword == "" {
		return errors.New("RPCPassword not specified")
	}
	return nil
}

/*
Save method saves the wallet without closing it.
*/
func (wallet *Walletd) Save() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	resp, err := wallet.makePostRequest("save", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
Reset method resyncs the wallet if no viewSecretKey is given.
If viewSecretKey is given then it replaces the existing wallet with a new one
corresponding to the viewSecretKey
*/
func (wallet *Walletd) Reset(scanHeight int) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["scanHeight"] = scanHeight
	resp, err := wallet.makePostRequest("reset", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
CreateAddress method creates a new address inside the container along with old addresses
*/
func (wallet *Walletd) CreateAddress(
	spendSecretKey string,
	spendPublicKey string,
	scanHeight int,
	newAddress bool) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	params := make(map[string]interface{})

	if spendSecretKey != "" && spendPublicKey != "" {
		return nil, errors.New("Cannot specify both spend keys.. either or both should be empty")
	} else if spendPublicKey == "" {
		params["spendSecretKey"] = spendSecretKey
	} else {
		params["spendPublicKey"] = spendPublicKey
	}

	if newAddress {
		params["newAddress"] = newAddress
	} else {
		params["scanHeight"] = scanHeight
	}

	resp, err := wallet.makePostRequest("createAddress", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
DeleteAddress method deletes the specified address from the container
*/
func (wallet *Walletd) DeleteAddress(address string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["address"] = address
	resp, err := wallet.makePostRequest("deleteAddress", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetSpendKeys method returns the spendPublicKey and spendSecretKey corresponding
the given input wallet address
*/
func (wallet *Walletd) GetSpendKeys(address string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["address"] = address
	resp, err := wallet.makePostRequest("getSpendKeys", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetBalance method returns the balance present in the specified address
If the address is empty then returns the balance present in the container
*/
func (wallet *Walletd) GetBalance(address string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["address"] = address
	resp, err := wallet.makePostRequest("getBalance", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetBlockHashes method returns array of hashes starting from specified blockIndex upto blockCount
*/
func (wallet *Walletd) GetBlockHashes(firstBlockIndex int, blockCount int) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["firstBlockIndex"] = firstBlockIndex
	params["blockCount"] = blockCount
	resp, err := wallet.makePostRequest("getBlockHashes", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetTransactionHashes method returns array of objects containing block and transaction hashes
of the specified address
*/
func (wallet *Walletd) GetTransactionHashes(
	addresses []string,
	blockHash string,
	firstBlockIndex int,
	blockCount int,
	paymentID string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})

	if blockHash != "" {
		params["blockHash"] = blockHash
	} else {
		params["firstBlockIndex"] = firstBlockIndex
	}

	params["addresses"] = addresses
	params["blockCount"] = blockCount
	params["paymentId"] = paymentID
	resp, err := wallet.makePostRequest("getTransactionHashes", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetTransactions method returns array of objects containing block and transaction details
of the specified address
*/
func (wallet *Walletd) GetTransactions(
	addresses []string,
	blockHash string,
	firstBlockIndex int,
	blockCount int,
	paymentID string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})

	if blockHash != "" {
		params["blockHash"] = blockHash
	} else {
		params["firstBlockIndex"] = firstBlockIndex
	}

	params["addresses"] = addresses
	params["blockCount"] = blockCount
	params["paymentId"] = paymentID
	resp, err := wallet.makePostRequest("getTransactions", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetUnconfirmedTransactionHashes method returns array of hashes of unconfirmed transactions of the specified address
*/
func (wallet *Walletd) GetUnconfirmedTransactionHashes(addresses []string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["addresses"] = addresses
	resp, err := wallet.makePostRequest("getUnconfirmedTransactionHashes", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetTransaction method returns the transaction details of a particular specified transaction hash
*/
func (wallet *Walletd) GetTransaction(transactionHash string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})

	if transactionHash == "" {
		return nil, errors.New("transactionHash cannot be empty.. please specify a valid hash")
	}

	params["transactionHash"] = transactionHash
	resp, err := wallet.makePostRequest("getTransaction", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
SendTransaction method sends specified transactions
*/
func (wallet *Walletd) SendTransaction(
	addresses []string,
	transfers []map[string]interface{},
	fee int,
	unlockTime int,
	extra string,
	paymentID string,
	changeAddress string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["addresses"] = addresses
	params["transfers"] = transfers
	params["fee"] = fee
	params["unlockTime"] = unlockTime
	params["changeAddress"] = changeAddress

	if extra != "" && paymentID != "" {
		return nil, errors.New("Can't set paymentID and extra together.. either or both should be empty")
	} else if extra != "" {
		params["extra"] = extra
	} else {
		params["paymentId"] = paymentID
	}

	resp, err := wallet.makePostRequest("sendTransaction", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
CreateDelayedTransaction method allows you to create a delayed transaction
Such transactions are not sent into the network automatically and should be pushed
using SendDelayedTransaction method
*/
func (wallet *Walletd) CreateDelayedTransaction(
	addresses []string,
	transfers []map[string]interface{},
	fee int,
	unlockTime int,
	extra string,
	paymentID string,
	changeAddress string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["addresses"] = addresses
	params["transfers"] = transfers
	params["fee"] = fee
	params["unlockTime"] = unlockTime
	params["changeAddress"] = changeAddress

	if extra != "" && paymentID != "" {
		return nil, errors.New("Can't set paymentID and extra together.. either or both should be empty")
	} else if extra != "" {
		params["extra"] = extra
	} else {
		params["paymentId"] = paymentID
	}

	resp, err := wallet.makePostRequest("createDelayedTransaction", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetDelayedTransactionHashes method returns array of delayedTransactionHashes
*/
func (wallet *Walletd) GetDelayedTransactionHashes() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	resp, err := wallet.makePostRequest("getDelayedTransactionHashes", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
DeleteDelayedTransaction method deletes the specified delayedTransactionHash
*/
func (wallet *Walletd) DeleteDelayedTransaction(transactionHash string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["transactionHash"] = transactionHash
	resp, err := wallet.makePostRequest("deleteDelayedTransaction", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
SendDelayedTransaction method sends the delayedTransaction created using CreateDelayedTransaction
method into the network
*/
func (wallet *Walletd) SendDelayedTransaction(transactionHash string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["transactionHash"] = transactionHash
	resp, err := wallet.makePostRequest("sendDelayedTransaction", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetViewKey method returns the viewSecretKey of the wallet
*/
func (wallet *Walletd) GetViewKey() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	resp, err := wallet.makePostRequest("getViewKey", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetMnemonicSeed method returns the 25 word random seed corresponding to
the given input wallet address
*/
func (wallet *Walletd) GetMnemonicSeed(address string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["address"] = address
	resp, err := wallet.makePostRequest("getMnemonicSeed", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetStatus method returns the sync state of the wallet and known top block height
*/
func (wallet *Walletd) GetStatus() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	resp, err := wallet.makePostRequest("getStatus", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetAddresses method returns an array of addresses present in the container
*/
func (wallet *Walletd) GetAddresses() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	resp, err := wallet.makePostRequest("getAddresses", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
SendFusionTransaction method allows you to send a fusion transaction from selected address to destination
address. If there aren't any outputs that can be optimized it returns an error.
*/
func (wallet *Walletd) SendFusionTransaction(
	threshold int,
	addresses []string,
	destinationAddress string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["threshold"] = threshold
	params["addresses"] = addresses
	params["destinationAddress"] = destinationAddress
	resp, err := wallet.makePostRequest("sendFusionTransaction", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
EstimateFusion method returns the number of outputs that can be optimized
This is helpful for sending fusion transactions
*/
func (wallet *Walletd) EstimateFusion(threshold int, addresses []string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["threshold"] = threshold
	params["addresses"] = addresses
	resp, err := wallet.makePostRequest("estimateFusion", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
CreateIntegratedAddress method creates a unique 236 char long address which corresponds to
the specified address with paymentID
*/
func (wallet *Walletd) CreateIntegratedAddress(address string, paymentID string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["address"] = address
	params["paymentId"] = paymentID
	resp, err := wallet.makePostRequest("createIntegratedAddress", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetFeeInfo method returns the fee information that the service picks up from the
connected daemon
*/
func (wallet *Walletd) GetFeeInfo() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	resp, err := wallet.makePostRequest("getFeeInfo", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}
