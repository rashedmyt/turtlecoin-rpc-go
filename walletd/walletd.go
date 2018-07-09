package walletd

import (
	"bytes"

	"../utils"
)

func Save(rpcPassword string, hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "save", params)
}

func Reset(rpcPassword string, hostURL string, hostPort int, viewSecretKey string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["viewSecretKey"] = viewSecretKey
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "reset", params)
}

func GetViewKey(rpcPassword string, hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "getViewKey", params)
}

func GetSpendKeys(rpcPassword string, hostURL string, hostPort int, address string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["address"] = address
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "getSpendKeys", params)
}

func GetMnemonicSeed(rpcPassword string, hostURL string, hostPort int, address string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["address"] = address
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "getMnemonicSeed", params)
}

func GetStatus(rpcPassword string, hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "getStatus", params)
}

func GetAddresses(rpcPassword string, hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "getAddresses", params)
}

func CreateAddress(rpcPassword string, hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "createAddress", params)
}

func DeleteAddress(rpcPassword string, hostURL string, hostPort int, address string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["address"] = address
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "deleteAddress", params)
}

func GetBalance(rpcPassword string, hostURL string, hostPort int, address string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["address"] = address
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "getBalance", params)
}

func GetBlockHashes(rpcPassword string, hostURL string, hostPort int, firstBlockIndex int, blockCount int) *bytes.Buffer {
	params := make(map[string]interface{})
	params["firstBlockIndex"] = firstBlockIndex
	params["blockCount"] = blockCount
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "getBlockHashes", params)
}

func GetTransactionHashes(rpcPassword string, hostURL string, hostPort int, firstBlockIndex int, blockCount int) *bytes.Buffer {
	params := make(map[string]interface{})
	params["firstBlockIndex"] = firstBlockIndex
	params["blockCount"] = blockCount
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "getTransactionHashes", params)
}

func GetTransactions(rpcPassword string, hostURL string, hostPort int, firstBlockIndex int, blockCount int) *bytes.Buffer {
	params := make(map[string]interface{})
	params["firstBlockIndex"] = firstBlockIndex
	params["blockCount"] = blockCount
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "getTransactions", params)
}

func GetUnconfirmedTransactionHashes(rpcPassword string, hostURL string, hostPort int, address string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["address"] = address
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "getUnconfirmedTransactionHashes", params)
}

func GetTransaction(rpcPassword string, hostURL string, hostPort int, transactionHash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["transactionHash"] = transactionHash
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "getTransaction", params)
}

func SendTransaction(rpcPassword string, hostURL string, hostPort int,
	addresses []string,
	transfers []map[string]interface{},
	fee int, unlockTime int, anonymity int, extra string, paymentID string, changeAddress string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["addresses"] = addresses
	params["transfers"] = transfers
	params["fee"] = fee
	params["unlockTime"] = unlockTime
	params["anonymity"] = anonymity
	params["extra"] = extra
	params["paymentId"] = paymentID
	params["changeAddress"] = changeAddress

	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "sendTransaction", params)
}

func CreateDelayedTransaction(rpcPassword string, hostURL string, hostPort int,
	addresses []string,
	transfers []map[string]interface{},
	fee int, unlockTime int, anonymity int, extra string, paymentID string, changeAddress string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["addresses"] = addresses
	params["transfers"] = transfers
	params["fee"] = fee
	params["unlockTime"] = unlockTime
	params["anonymity"] = anonymity
	params["extra"] = extra
	params["paymentId"] = paymentID
	params["changeAddress"] = changeAddress

	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "createDelayedTransaction", params)
}

func GetDelayedTransactionHashes(rpcPassword string, hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "getDelayedTransactionHashes", params)
}

func DeleteDelayedTransaction(rpcPassword string, hostURL string, hostPort int, transactionHash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["transactionHash"] = transactionHash
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "deleteDelayedTransaction", params)
}

func SendDelayedTransaction(rpcPassword string, hostURL string, hostPort int, transactionHash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["transactionHash"] = transactionHash
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "sendDelayedTransaction", params)
}

func SendFusionTransaction(rpcPassword string, hostURL string, hostPort int,
	threshold int, anonymity int, addresses []string, destinationAddress string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["threshold"] = threshold
	params["anonymity"] = anonymity
	params["addresses"] = addresses
	params["destinationAddress"] = destinationAddress
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "sendFusionTransaction", params)
}

func EstimateFusion(rpcPassword string, hostURL string, hostPort int, threshold int, addresses []string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["threshold"] = threshold
	params["addresses"] = addresses
	return utils.MakeWalletPostRequest(rpcPassword, hostURL, hostPort, "estimateFusion", params)
}
