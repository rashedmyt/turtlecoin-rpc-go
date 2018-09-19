/*

Copyright (c) 2018 Rashed Mohammed, The TurtleCoin Developers

Please see the included LICENSE file for more information

*/

package turtlecoinrpc

import "bytes"

// Walletd structure contains the URL and Port
// info of the node and RPC Password for RPC calls
type Walletd struct {
	URL         string
	Port        int
	RPCPassword string
}

func (wallet *Walletd) check() {
	if wallet.URL == "" {
		wallet.URL = "127.0.0.1"
	}
	if wallet.Port == 0 {
		wallet.Port = 8070
	}
}

/*
Save method saves the wallet without closing it.
*/
func (wallet *Walletd) Save() *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	return wallet.makePostRequest("save", params)
}

/*
Reset method resyncs the wallet if no viewSecretKey is given.
If viewSecretKey is given then it replaces the existing wallet with a new one
corresponding to the viewSecretKey
*/
func (wallet *Walletd) Reset(viewSecretKey string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["viewSecretKey"] = viewSecretKey
	return wallet.makePostRequest("reset", params)
}

/*
CreateAddress method creates a new address inside the container along with old addresses
*/
func (wallet *Walletd) CreateAddress() *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	return wallet.makePostRequest("createAddress", params)
}

/*
DeleteAddress method deletes the specified address from the container
*/
func (wallet *Walletd) DeleteAddress(address string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["address"] = address
	return wallet.makePostRequest("deleteAddress", params)
}

/*
GetSpendKeys method returns the spendPublicKey and spendSecretKey corresponding
the given input wallet address
*/
func (wallet *Walletd) GetSpendKeys(address string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["address"] = address
	return wallet.makePostRequest("getSpendKeys", params)
}

/*
GetBalance method returns the balance present in the specified address
If the address is empty then returns the balance present in the container
*/
func (wallet *Walletd) GetBalance(address string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["address"] = address
	return wallet.makePostRequest("getBalance", params)
}

/*
GetBlockHashes method returns array of hashes starting from specified blockIndex upto blockCount
*/
func (wallet *Walletd) GetBlockHashes(firstBlockIndex int, blockCount int) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["firstBlockIndex"] = firstBlockIndex
	params["blockCount"] = blockCount
	return wallet.makePostRequest("getBlockHashes", params)
}

/*
GetTransactionHashes method returns array of objects containing block and transaction hashes
of the specified address
*/
func (wallet *Walletd) GetTransactionHashes(firstBlockIndex int, blockCount int) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["firstBlockIndex"] = firstBlockIndex
	params["blockCount"] = blockCount
	return wallet.makePostRequest("getTransactionHashes", params)
}

/*
GetTransactions method returns array of objects containing block and transaction details
of the specified address
*/
func (wallet *Walletd) GetTransactions(firstBlockIndex int, blockCount int) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["firstBlockIndex"] = firstBlockIndex
	params["blockCount"] = blockCount
	return wallet.makePostRequest("getTransactions", params)
}

/*
GetUnconfirmedTransactionHashes method returns array of hashes of unconfirmed transactions of the specified address
*/
func (wallet *Walletd) GetUnconfirmedTransactionHashes(address string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["address"] = address
	return wallet.makePostRequest("getUnconfirmedTransactionHashes", params)
}

/*
GetTransaction method returns the transaction details of a particular specified transaction hash
*/
func (wallet *Walletd) GetTransaction(transactionHash string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["transactionHash"] = transactionHash
	return wallet.makePostRequest("getTransaction", params)
}

/*
SendTransaction method sends specified transactions
*/
func (wallet *Walletd) SendTransaction(
	rpcPassword string,
	hostURL string,
	hostPort int,
	addresses []string,
	transfers []map[string]interface{},
	fee int,
	unlockTime int,
	extra string,
	paymentID string,
	changeAddress string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["addresses"] = addresses
	params["transfers"] = transfers
	params["fee"] = fee
	params["unlockTime"] = unlockTime
	params["changeAddress"] = changeAddress
	if extra != "" && paymentID != "" {
		panic("Can't set paymentId and extra together")
	} else if extra != "" {
		params["extra"] = extra
	} else {
		params["paymentId"] = paymentID
	}

	return wallet.makePostRequest("sendTransaction", params)
}

/*
CreateDelayedTransaction method allows you to create a delayed transaction
Such transactions are not sent into the network automatically and should be pushed
using SendDelayedTransaction method
*/
func (wallet *Walletd) CreateDelayedTransaction(
	rpcPassword string,
	hostURL string,
	hostPort int,
	addresses []string,
	transfers []map[string]interface{},
	fee int,
	unlockTime int,
	extra string,
	paymentID string,
	changeAddress string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["addresses"] = addresses
	params["transfers"] = transfers
	params["fee"] = fee
	params["unlockTime"] = unlockTime
	params["changeAddress"] = changeAddress
	if extra != "" && paymentID != "" {
		panic("Can't set paymentId and extra together")
	} else if extra != "" {
		params["extra"] = extra
	} else {
		params["paymentId"] = paymentID
	}

	return wallet.makePostRequest("createDelayedTransaction", params)
}

/*
GetDelayedTransactionHashes method returns array of delayedTransactionHashes
*/
func (wallet *Walletd) GetDelayedTransactionHashes() *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	return wallet.makePostRequest("getDelayedTransactionHashes", params)
}

/*
DeleteDelayedTransaction method deletes the specified delayedTransactionHash
*/
func (wallet *Walletd) DeleteDelayedTransaction(transactionHash string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["transactionHash"] = transactionHash
	return wallet.makePostRequest("deleteDelayedTransaction", params)
}

/*
SendDelayedTransaction method sends the delayedTransaction created using CreateDelayedTransaction
method into the network
*/
func (wallet *Walletd) SendDelayedTransaction(transactionHash string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["transactionHash"] = transactionHash
	return wallet.makePostRequest("sendDelayedTransaction", params)
}

/*
GetViewKey method returns the viewSecretKey of the wallet
*/
func (wallet *Walletd) GetViewKey() *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	return wallet.makePostRequest("getViewKey", params)
}

/*
GetMnemonicSeed method returns the 25 word random seed corresponding to
the given input wallet address
*/
func (wallet *Walletd) GetMnemonicSeed(address string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["address"] = address
	return wallet.makePostRequest("getMnemonicSeed", params)
}

/*
GetStatus method returns the sync state of the wallet and known top block height
*/
func (wallet *Walletd) GetStatus() *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	return wallet.makePostRequest("getStatus", params)
}

/*
GetAddresses method returns an array of addresses present in the container
*/
func (wallet *Walletd) GetAddresses() *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	return wallet.makePostRequest("getAddresses", params)
}

/*
SendFusionTransaction method allows you to send a fusion transaction from selected address to destination
address. If there aren't any outputs that can be optimized it returns an error.
*/
func (wallet *Walletd) SendFusionTransaction(
	rpcPassword string,
	hostURL string,
	hostPort int,
	threshold int,
	addresses []string,
	destinationAddress string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["threshold"] = threshold
	params["addresses"] = addresses
	params["destinationAddress"] = destinationAddress
	return wallet.makePostRequest("sendFusionTransaction", params)
}

/*
EstimateFusion method returns the number of outputs that can be optimized
This is helpful for sending fusion transactions
*/
func (wallet *Walletd) EstimateFusion(threshold int, addresses []string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["threshold"] = threshold
	params["addresses"] = addresses
	return wallet.makePostRequest("estimateFusion", params)
}

/*
CreateIntegratedAddress method creates a unique 236 char long address which corresponds to
the specified address with paymentID
*/
func (wallet *Walletd) CreateIntegratedAddress(address string, paymentID string) *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	params["address"] = address
	params["paymentId"] = paymentID
	return wallet.makePostRequest("createIntegratedAddress", params)
}

/*
GetFeeInfo method returns the fee information that the service picks up from the
connected daemon
*/
func (wallet *Walletd) GetFeeInfo() *bytes.Buffer {
	wallet.check()
	params := make(map[string]interface{})
	return wallet.makePostRequest("getFeeInfo", params)
}
