/*

Copyright (c) 2018 Rashed Mohammed, The TurtleCoin Developers

Please see the included LICENSE file for more information

*/

package turtlecoind

import (
	"bytes"
)

/*
GetHeight method returns the height of the blockchain
*/
func GetHeight(hostURL string, hostPort int) *bytes.Buffer {
	return makeGetRequest("getheight", hostURL, hostPort)
}

/*
GetInfo method returns information related to network and connection
*/
func GetInfo(hostURL string, hostPort int) *bytes.Buffer {
	return makeGetRequest("getinfo", hostURL, hostPort)
}

/*
GetTransactions returns array of missed transactions
*/
func GetTransactions(hostURL string, hostPort int) *bytes.Buffer {
	return makeGetRequest("gettransactions", hostURL, hostPort)
}

/*
GetPeers method returns array of peers connected to daemon
*/
func GetPeers(hostURL string, hostPort int) *bytes.Buffer {
	return makeGetRequest("getpeers", hostURL, hostPort)
}

/*
GetFeeInfo method returns the fee set by the remote node
*/
func GetFeeInfo(hostURL string, hostPort int) *bytes.Buffer {
	return makeGetRequest("feeinfo", hostURL, hostPort)
}

/*
GetBlockCount method returns the height of the top block
*/
func GetBlockCount(hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(hostURL, hostPort, "getblockcount", params)
}

/*
GetBlockHash method returns the block hash by height
*/
func GetBlockHash(hostURL string, hostPort int, height int) *bytes.Buffer {
	params := []int{height}
	return makePostRequest(hostURL, hostPort, "on_getblockhash", params)
}

/*
GetBlockTemplate method returns the block template blob of the last block
*/
func GetBlockTemplate(hostURL string, hostPort int, reserveSize int, walletAddress string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["reserve_size"] = reserveSize
	params["wallet_address"] = walletAddress
	return makePostRequest(hostURL, hostPort, "getblocktemplate", params)
}

/*
SubmitBlock method submits a block to the network corresponding to the input block blob
*/
func SubmitBlock(hostURL string, hostPort int, blockBlob string) *bytes.Buffer {
	params := []string{blockBlob}
	return makePostRequest(hostURL, hostPort, "submitblock", params)
}

/*
GetLastBlockHeader method returns the block header of the last block
*/
func GetLastBlockHeader(hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(hostURL, hostPort, "getlastblockheader", params)
}

/*
GetBlockHeaderByHash method returns the block header corresponding to the input block hash
*/
func GetBlockHeaderByHash(hostURL string, hostPort int, hash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["hash"] = hash
	return makePostRequest(hostURL, hostPort, "getblockheaderbyhash", params)
}

/*
GetBlockHeaderByHeight method returns the block header corresponding to the input block height
*/
func GetBlockHeaderByHeight(hostURL string, hostPort int, height int) *bytes.Buffer {
	params := make(map[string]interface{})
	params["height"] = height
	return makePostRequest(hostURL, hostPort, "getblockheaderbyheight", params)
}

/*
GetCurrencyID method returns the currency id of the network
*/
func GetCurrencyID(hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(hostURL, hostPort, "getcurrencyid", params)
}

/*
GetBlocks method returns information on 30 blocks from specified height (inclusive)
*/
func GetBlocks(hostURL string, hostPort int, height int) *bytes.Buffer {
	params := make(map[string]interface{})
	params["height"] = height
	return makePostRequest(hostURL, hostPort, "f_blocks_list_json", params)
}

/*
GetBlock method returns the information of block corresponding to given input hash
*/
func GetBlock(hostURL string, hostPort int, hash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["hash"] = hash
	return makePostRequest(hostURL, hostPort, "f_block_json", params)
}

/*
GetTransaction method returns informaiton of transaction corresponding to given input hash
*/
func GetTransaction(hostURL string, hostPort int, hash string) *bytes.Buffer {
	params := make(map[string]interface{})
	params["hash"] = hash
	return makePostRequest(hostURL, hostPort, "f_transaction_json", params)
}

/*
GetTransactionPool method returns the list of unconfirmed transactions present in mem pool
*/
func GetTransactionPool(hostURL string, hostPort int) *bytes.Buffer {
	params := make(map[string]interface{})
	return makePostRequest(hostURL, hostPort, "f_on_transactions_pool_json", params)
}
