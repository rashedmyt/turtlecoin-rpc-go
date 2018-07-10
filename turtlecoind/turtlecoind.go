package turtlecoind

import (
	"bytes"
)

func GetHeight(hostURL string, hostPort int) *bytes.Buffer {
	return makeGetRequest("getheight", hostURL, hostPort)
}

func GetInfo(hostURL string, hostPort int) *bytes.Buffer {
	return makeGetRequest("getinfo", hostURL, hostPort)
}

func GetTransactions(hostURL string, hostPort int) *bytes.Buffer {
	return makeGetRequest("gettransactions", hostURL, hostPort)
}

func GetPeers(hostURL string, hostPort int) *bytes.Buffer {
	return makeGetRequest("getpeers", hostURL, hostPort)
}

func GetBlockCount(hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(hostURL, hostPort, "getblockcount", params)
}

func GetBlockHash(hostURL string, hostPort int, height int) *bytes.Buffer {
	params := []int{height}
	return makePostRequest(hostURL, hostPort, "on_getblockhash", params)
}

func GetBlockTemplate(hostURL string, hostPort int, reserveSize int, walletAddress string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["reserve_size"] = reserveSize
	params["wallet_address"] = walletAddress
	return makePostRequest(hostURL, hostPort, "getblocktemplate", params)
}

func SubmitBlock(hostURL string, hostPort int, blockBlob string) *bytes.Buffer {
	params := []string{blockBlob}
	return makePostRequest(hostURL, hostPort, "submitblock", params)
}

func GetLastBlockHeader(hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(hostURL, hostPort, "getlastblockheader", params)
}

func GetBlockHeaderByHash(hostURL string, hostPort int, hash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["hash"] = hash
	return makePostRequest(hostURL, hostPort, "getblockheaderbyhash", params)
}

func GetBlockHeaderByHeight(hostURL string, hostPort int, height int) *bytes.Buffer {
	params := make(map[string]interface{})
	params["height"] = height
	return makePostRequest(hostURL, hostPort, "getblockheaderbyheight", params)
}

func GetCurrencyID(hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(hostURL, hostPort, "getcurrencyid", params)
}

func GetBlocks(hostURL string, hostPort int, height int) *bytes.Buffer {
	params := make(map[string]interface{})
	params["height"] = height
	return makePostRequest(hostURL, hostPort, "f_blocks_list_json", params)
}

func GetBlock(hostURL string, hostPort int, hash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["hash"] = hash
	return makePostRequest(hostURL, hostPort, "f_block_json", params)
}

func GetTransaction(hostURL string, hostPort int, hash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["hash"] = hash
	return makePostRequest(hostURL, hostPort, "f_transaction_json", params)
}

func GetTransactionPool(hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(hostURL, hostPort, "f_on_transactions_pool_json", params)
}
