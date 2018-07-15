/*

Copyright (c) 2018 Rashed Mohammed, The TurtleCoin Developers

Please see the included LICENSE file for more information

*/

package walletd

import (
	"bytes"
)

/*
Save method saves the wallet without closing it.
*/
func Save(rpcPassword string, hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(rpcPassword, hostURL, hostPort, "save", params)
}

/*
Reset method resyncs the wallet if no viewSecretKey is given.

If viewSecretKey is given then it replaces the existing wallet with a new one
corresponding to the viewSecretKey
*/
func Reset(rpcPassword string, hostURL string, hostPort int, viewSecretKey string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["viewSecretKey"] = viewSecretKey
	return makePostRequest(rpcPassword, hostURL, hostPort, "reset", params)
}

/*
GetViewKey method returns the viewSecretKey of the wallet
*/
func GetViewKey(rpcPassword string, hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(rpcPassword, hostURL, hostPort, "getViewKey", params)
}

/*
GetSpendKeys method returns the spendPublicKey and spendSecretKey corresponding
the given input wallet address
*/
func GetSpendKeys(rpcPassword string, hostURL string, hostPort int, address string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["address"] = address
	return makePostRequest(rpcPassword, hostURL, hostPort, "getSpendKeys", params)
}

/*
GetMnemonicSeed method returns the 25 word random seed corresponding to
the given input wallet address
*/
func GetMnemonicSeed(rpcPassword string, hostURL string, hostPort int, address string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["address"] = address
	return makePostRequest(rpcPassword, hostURL, hostPort, "getMnemonicSeed", params)
}

/*
GetStatus method returns the sync state of the wallet and known top block height
*/
func GetStatus(rpcPassword string, hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(rpcPassword, hostURL, hostPort, "getStatus", params)
}

/*
GetAddresses method returns an array of addresses present in the container
*/
func GetAddresses(rpcPassword string, hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(rpcPassword, hostURL, hostPort, "getAddresses", params)
}

/*
CreateAddress method creates a new address inside the container along with old addresses
*/
func CreateAddress(rpcPassword string, hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(rpcPassword, hostURL, hostPort, "createAddress", params)
}

/*
DeleteAddress method deletes the specified address from the container
*/
func DeleteAddress(rpcPassword string, hostURL string, hostPort int, address string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["address"] = address
	return makePostRequest(rpcPassword, hostURL, hostPort, "deleteAddress", params)
}

/*
GetBalance method returns the balance present in the specified address

If the address is empty then returns the balance present in the container
*/
func GetBalance(rpcPassword string, hostURL string, hostPort int, address string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["address"] = address
	return makePostRequest(rpcPassword, hostURL, hostPort, "getBalance", params)
}

/*
GetBlockHashes method returns array of hashes starting from specified blockIndex upto blockCount
*/
func GetBlockHashes(rpcPassword string, hostURL string, hostPort int, firstBlockIndex int, blockCount int) *bytes.Buffer {
	params := make(map[string]interface{})
	params["firstBlockIndex"] = firstBlockIndex
	params["blockCount"] = blockCount
	return makePostRequest(rpcPassword, hostURL, hostPort, "getBlockHashes", params)
}

/*
GetTransactionHashes method returns array of objects containing block and transaction hashes
of the specified address
*/
func GetTransactionHashes(rpcPassword string, hostURL string, hostPort int, firstBlockIndex int, blockCount int) *bytes.Buffer {
	params := make(map[string]interface{})
	params["firstBlockIndex"] = firstBlockIndex
	params["blockCount"] = blockCount
	return makePostRequest(rpcPassword, hostURL, hostPort, "getTransactionHashes", params)
}

/*
GetTransactions method returns array of objects containing block and transaction details
of the specified address
*/
func GetTransactions(rpcPassword string, hostURL string, hostPort int, firstBlockIndex int, blockCount int) *bytes.Buffer {
	params := make(map[string]interface{})
	params["firstBlockIndex"] = firstBlockIndex
	params["blockCount"] = blockCount
	return makePostRequest(rpcPassword, hostURL, hostPort, "getTransactions", params)
}

/*
GetUnconfirmedTransactionHashes method returns array of hashes of unconfirmed transactions of the specified address
*/
func GetUnconfirmedTransactionHashes(rpcPassword string, hostURL string, hostPort int, address string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["address"] = address
	return makePostRequest(rpcPassword, hostURL, hostPort, "getUnconfirmedTransactionHashes", params)
}

/*
GetTransaction method returns the transaction details of a particular specified transaction hash
*/
func GetTransaction(rpcPassword string, hostURL string, hostPort int, transactionHash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["transactionHash"] = transactionHash
	return makePostRequest(rpcPassword, hostURL, hostPort, "getTransaction", params)
}

/*
SendTransaction method sends specified transactions
*/
func SendTransaction(
	rpcPassword string,
	hostURL string,
	hostPort int,
	addresses []string,
	transfers []map[string]interface{},
	fee int,
	unlockTime int,
	anonymity int,
	extra string,
	paymentID string,
	changeAddress string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["addresses"] = addresses
	params["transfers"] = transfers
	params["fee"] = fee
	params["unlockTime"] = unlockTime
	params["anonymity"] = anonymity
	params["changeAddress"] = changeAddress
	if extra != "" && paymentID != "" {
		panic("Can't set paymentId and extra together")
	} else if extra != "" {
		params["extra"] = extra
	} else {
		params["paymentId"] = paymentID
	}

	return makePostRequest(rpcPassword, hostURL, hostPort, "sendTransaction", params)
}

/*
CreateDelayedTransaction method allows you to create a delayed transaction

Such transactions are not sent into the network automatically and should be pushed
using SendDelayedTransaction method
*/
func CreateDelayedTransaction(
	rpcPassword string,
	hostURL string,
	hostPort int,
	addresses []string,
	transfers []map[string]interface{},
	fee int,
	unlockTime int,
	anonymity int,
	extra string,
	paymentID string,
	changeAddress string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["addresses"] = addresses
	params["transfers"] = transfers
	params["fee"] = fee
	params["unlockTime"] = unlockTime
	params["anonymity"] = anonymity
	params["changeAddress"] = changeAddress
	if extra != "" && paymentID != "" {
		panic("Can't set paymentId and extra together")
	} else if extra != "" {
		params["extra"] = extra
	} else {
		params["paymentId"] = paymentID
	}

	return makePostRequest(rpcPassword, hostURL, hostPort, "createDelayedTransaction", params)
}

/*
GetDelayedTransactionHashes method returns array of delayedTransactionHashes
*/
func GetDelayedTransactionHashes(rpcPassword string, hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(rpcPassword, hostURL, hostPort, "getDelayedTransactionHashes", params)
}

/*
DeleteDelayedTransaction method deletes the specified delayedTransactionHash
*/
func DeleteDelayedTransaction(rpcPassword string, hostURL string, hostPort int, transactionHash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["transactionHash"] = transactionHash
	return makePostRequest(rpcPassword, hostURL, hostPort, "deleteDelayedTransaction", params)
}

/*
SendDelayedTransaction method sends the delayedTransaction created using CreateDelayedTransaction
method into the network
*/
func SendDelayedTransaction(rpcPassword string, hostURL string, hostPort int, transactionHash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["transactionHash"] = transactionHash
	return makePostRequest(rpcPassword, hostURL, hostPort, "sendDelayedTransaction", params)
}

/*
SendFusionTransaction method allows you to send a fusion transaction from selected address to destination
address. If there aren't any outputs that can be optimized it returns an error.
*/
func SendFusionTransaction(
	rpcPassword string,
	hostURL string,
	hostPort int,
	threshold int,
	anonymity int,
	addresses []string,
	destinationAddress string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["threshold"] = threshold
	params["anonymity"] = anonymity
	params["addresses"] = addresses
	params["destinationAddress"] = destinationAddress
	return makePostRequest(rpcPassword, hostURL, hostPort, "sendFusionTransaction", params)
}

/*
EstimateFusion method returns the number of outputs that can be optimized

This is helpful for sending fusion transactions
*/
func EstimateFusion(rpcPassword string, hostURL string, hostPort int, threshold int, addresses []string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["threshold"] = threshold
	params["addresses"] = addresses
	return makePostRequest(rpcPassword, hostURL, hostPort, "estimateFusion", params)
}

/*
CreateIntegratedAddress method creates a unique 236 char long address which corresponds to
the specified address with paymentID
*/
func CreateIntegratedAddress(rpcPassword string, hostURL string, hostPort int, address string, paymentID string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["address"] = address
	params["paymentId"] = paymentID
	return makePostRequest(rpcPassword, hostURL, hostPort, "createIntegratedAddress", params)
}
