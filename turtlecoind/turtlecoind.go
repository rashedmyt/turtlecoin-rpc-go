package turtlecoind

import (
	"bytes"

	"../utils"
)

func GetHeight(hostURL string, hostPort int) *bytes.Buffer {
	return utils.MakeGetRequest("getheight", hostURL, hostPort)
}

func GetInfo(hostURL string, hostPort int) *bytes.Buffer {
	return utils.MakeGetRequest("getinfo", hostURL, hostPort)
}

func GetTransactions(hostURL string, hostPort int) *bytes.Buffer {
	return utils.MakeGetRequest("gettransactions", hostURL, hostPort)
}

func GetPeers(hostURL string, hostPort int) *bytes.Buffer {
	return utils.MakeGetRequest("getpeers", hostURL, hostPort)
}

func GetBlockCount(hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return utils.MakePostRequest(hostURL, hostPort, "getblockcount", params)
}

func GetBlockHash(hostURL string, hostPort int, height int) *bytes.Buffer {
	params := []int{height}
	return utils.MakePostRequest(hostURL, hostPort, "on_getblockhash", params)
}

func GetBlockTemplate(hostURL string, hostPort int, reserveSize int, walletAddress string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["reserve_size"] = reserveSize
	params["wallet_address"] = walletAddress
	return utils.MakePostRequest(hostURL, hostPort, "getblocktemplate", params)
}

func SubmitBlock(hostURL string, hostPort int, blockBlob string) *bytes.Buffer {
	params := []string{blockBlob}
	return utils.MakePostRequest(hostURL, hostPort, "submitblock", params)
}

func GetLastBlockHeader(hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return utils.MakePostRequest(hostURL, hostPort, "getlastblockheader", params)
}

func GetBlockHeaderByHash(hostURL string, hostPort int, hash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["hash"] = hash
	return utils.MakePostRequest(hostURL, hostPort, "getblockheaderbyhash", params)
}

func GetBlockHeaderByHeight(hostURL string, hostPort int, height int) *bytes.Buffer {
	params := make(map[string]interface{})
	params["height"] = height
	return utils.MakePostRequest(hostURL, hostPort, "getblockheaderbyheight", params)
}

func GetCurrencyID(hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return utils.MakePostRequest(hostURL, hostPort, "getcurrencyid", params)
}

func GetBlocks(hostURL string, hostPort int, height int) *bytes.Buffer {
	params := make(map[string]interface{})
	params["height"] = height
	return utils.MakePostRequest(hostURL, hostPort, "f_blocks_list_json", params)
}

func GetBlock(hostURL string, hostPort int, hash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["hash"] = hash
	return utils.MakePostRequest(hostURL, hostPort, "f_block_json", params)
}

func GetTransaction(hostURL string, hostPort int, hash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["hash"] = hash
	return utils.MakePostRequest(hostURL, hostPort, "f_transaction_json", params)
}

func GetTransactionPool(hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return utils.MakePostRequest(hostURL, hostPort, "f_on_transactions_pool_json", params)
}
