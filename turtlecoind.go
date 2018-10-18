/*

Copyright (c) 2018 Rashed Mohammed, The TurtleCoin Developers

Please see the included LICENSE file for more information

*/

package turtlecoinrpc

import (
	"bytes"
)

// TurtleCoind structure contains the
// URL and Port info of node for RPC calls
type TurtleCoind struct {
	URL  string
	Port int
}

func (daemon *TurtleCoind) check() {
	if daemon.URL == "" {
		daemon.URL = "127.0.0.1"
	}
	if daemon.Port == 0 {
		daemon.Port = 11898
	}
}

/*
Info method returns information related to network and connection
*/
func (daemon *TurtleCoind) Info() *bytes.Buffer {
	daemon.check()
	return daemon.makeGetRequest("getinfo")
}

/*
Height method returns the height of the blockchain
*/
func (daemon *TurtleCoind) Height() *bytes.Buffer {
	daemon.check()
	return daemon.makeGetRequest("getheight")
}

/*
func (daemon *TurtleCoind) Transactions() *bytes.Buffer {
	daemon.check()
	return daemon.makeGetRequest("gettransactions")
}
*/

/*
Fee method returns the fee set by the node
*/
func (daemon *TurtleCoind) Fee() *bytes.Buffer {
	daemon.check()
	return daemon.makeGetRequest("feeinfo")
}

/*
Peers method returns array of peers connected to daemon
*/
func (daemon *TurtleCoind) Peers() *bytes.Buffer {
	daemon.check()
	return daemon.makeGetRequest("getpeers")
}

/*
GetBlocks method returns information on 30 blocks from specified height (inclusive)
*/
func (daemon *TurtleCoind) GetBlocks(height int) *bytes.Buffer {
	daemon.check()
	params := make(map[string]interface{})
	params["height"] = height
	return daemon.makePostRequest("f_blocks_list_json", params)
}

/*
GetBlock method returns the information of block corresponding to given input hash
*/
func (daemon *TurtleCoind) GetBlock(hash string) *bytes.Buffer {
	daemon.check()
	params := make(map[string]interface{})
	params["hash"] = hash
	return daemon.makePostRequest("f_block_json", params)
}

/*
GetTransaction method returns information of transaction corresponding to given input hash
*/
func (daemon *TurtleCoind) GetTransaction(hash string) *bytes.Buffer {
	daemon.check()
	params := make(map[string]interface{})
	params["hash"] = hash
	return daemon.makePostRequest("f_transaction_json", params)
}

/*
GetTransactionPool method returns the list of unconfirmed transactions present in mem pool
*/
func (daemon *TurtleCoind) GetTransactionPool() *bytes.Buffer {
	daemon.check()
	params := make(map[string]interface{})
	return daemon.makePostRequest("f_on_transactions_pool_json", params)
}

/*
GetBlockCount method returns the height of the top block
*/
func (daemon *TurtleCoind) GetBlockCount() *bytes.Buffer {
	daemon.check()
	params := make(map[string]interface{})
	return daemon.makePostRequest("getblockcount", params)
}

/*
GetBlockHash method returns the block hash by height
*/
func (daemon *TurtleCoind) GetBlockHash(height int) *bytes.Buffer {
	daemon.check()
	params := []int{height}
	return daemon.makePostRequest("on_getblockhash", params)
}

/*
GetBlockTemplate method returns the block template blob of the last block
*/
func (daemon *TurtleCoind) GetBlockTemplate(reserveSize int, walletAddress string) *bytes.Buffer {
	daemon.check()
	params := make(map[string]interface{})
	params["reserve_size"] = reserveSize
	params["wallet_address"] = walletAddress
	return daemon.makePostRequest("getblocktemplate", params)
}

/*
GetCurrencyID method returns the currency id of the network
*/
func (daemon *TurtleCoind) GetCurrencyID() *bytes.Buffer {
	daemon.check()
	params := make(map[string]interface{})
	return daemon.makePostRequest("getcurrencyid", params)
}

/*
SubmitBlock method submits a block to the network corresponding to the input block blob
*/
func (daemon *TurtleCoind) SubmitBlock(blockBlob string) *bytes.Buffer {
	daemon.check()
	params := []string{blockBlob}
	return daemon.makePostRequest("submitblock", params)
}

/*
GetLastBlockHeader method returns the block header of the last block
*/
func (daemon *TurtleCoind) GetLastBlockHeader() *bytes.Buffer {
	daemon.check()
	params := make(map[string]interface{})
	return daemon.makePostRequest("getlastblockheader", params)
}

/*
GetBlockHeaderByHash method returns the block header corresponding to the input block hash
*/
func (daemon *TurtleCoind) GetBlockHeaderByHash(hash string) *bytes.Buffer {
	daemon.check()
	params := make(map[string]interface{})
	params["hash"] = hash
	return daemon.makePostRequest("getblockheaderbyhash", params)
}

/*
GetBlockHeaderByHeight method returns the block header corresponding to the input block height
*/
func (daemon *TurtleCoind) GetBlockHeaderByHeight(height int) *bytes.Buffer {
	daemon.check()
	params := make(map[string]interface{})
	params["height"] = height
	return daemon.makePostRequest("getblockheaderbyheight", params)
}
