// Copyright (c) 2018-2019 Rashed Mohammed, The TurtleCoin Developers
// Please see the included LICENSE file for more information

package turtlecoinrpc

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
func (daemon *TurtleCoind) Info() (map[string]interface{}, error) {
	daemon.check()
	resp, err := daemon.makeGetRequest("getinfo")
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
Height method returns the height of the blockchain
*/
func (daemon *TurtleCoind) Height() (map[string]interface{}, error) {
	daemon.check()
	resp, err := daemon.makeGetRequest("getheight")
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
func (daemon *TurtleCoind) Transactions() map[string]interface{} {
	daemon.check()
	return daemon.makeGetRequest("gettransactions")
}
*/

/*
Fee method returns the fee set by the node
*/
func (daemon *TurtleCoind) Fee() (map[string]interface{}, error) {
	daemon.check()
	resp, err := daemon.makeGetRequest("feeinfo")
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
Peers method returns array of peers connected to daemon
*/
func (daemon *TurtleCoind) Peers() (map[string]interface{}, error) {
	daemon.check()
	resp, err := daemon.makeGetRequest("getpeers")
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetBlocks method returns information on 30 blocks from specified height (inclusive)
*/
func (daemon *TurtleCoind) GetBlocks(height int) (map[string]interface{}, error) {
	daemon.check()
	params := make(map[string]interface{})
	params["height"] = height
	resp, err := daemon.makePostRequest("f_blocks_list_json", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetBlock method returns the information of block corresponding to given input hash
*/
func (daemon *TurtleCoind) GetBlock(hash string) (map[string]interface{}, error) {
	daemon.check()
	params := make(map[string]interface{})
	params["hash"] = hash
	resp, err := daemon.makePostRequest("f_block_json", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetTransaction method returns information of transaction corresponding to given input hash
*/
func (daemon *TurtleCoind) GetTransaction(hash string) (map[string]interface{}, error) {
	daemon.check()
	params := make(map[string]interface{})
	params["hash"] = hash
	resp, err := daemon.makePostRequest("f_transaction_json", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetTransactionPool method returns the list of unconfirmed transactions present in mem pool
*/
func (daemon *TurtleCoind) GetTransactionPool() (map[string]interface{}, error) {
	daemon.check()
	params := make(map[string]interface{})
	resp, err := daemon.makePostRequest("f_on_transactions_pool_json", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetBlockCount method returns the height of the top block
*/
func (daemon *TurtleCoind) GetBlockCount() (map[string]interface{}, error) {
	daemon.check()
	params := make(map[string]interface{})
	resp, err := daemon.makePostRequest("getblockcount", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetBlockHash method returns the block hash by height
*/
func (daemon *TurtleCoind) GetBlockHash(height int) (map[string]interface{}, error) {
	daemon.check()
	params := []int{height}
	resp, err := daemon.makePostRequest("on_getblockhash", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetBlockTemplate method returns the block template blob of the last block
*/
func (daemon *TurtleCoind) GetBlockTemplate(reserveSize int, walletAddress string) (map[string]interface{}, error) {
	daemon.check()
	params := make(map[string]interface{})
	params["reserve_size"] = reserveSize
	params["wallet_address"] = walletAddress
	resp, err := daemon.makePostRequest("getblocktemplate", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetCurrencyID method returns the currency id of the network
*/
func (daemon *TurtleCoind) GetCurrencyID() (map[string]interface{}, error) {
	daemon.check()
	params := make(map[string]interface{})
	resp, err := daemon.makePostRequest("getcurrencyid", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
SubmitBlock method submits a block to the network corresponding to the input block blob
*/
func (daemon *TurtleCoind) SubmitBlock(blockBlob string) (map[string]interface{}, error) {
	daemon.check()
	params := []string{blockBlob}
	resp, err := daemon.makePostRequest("submitblock", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetLastBlockHeader method returns the block header of the last block
*/
func (daemon *TurtleCoind) GetLastBlockHeader() (map[string]interface{}, error) {
	daemon.check()
	params := make(map[string]interface{})
	resp, err := daemon.makePostRequest("getlastblockheader", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetBlockHeaderByHash method returns the block header corresponding to the input block hash
*/
func (daemon *TurtleCoind) GetBlockHeaderByHash(hash string) (map[string]interface{}, error) {
	daemon.check()
	params := make(map[string]interface{})
	params["hash"] = hash
	resp, err := daemon.makePostRequest("getblockheaderbyhash", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

/*
GetBlockHeaderByHeight method returns the block header corresponding to the input block height
*/
func (daemon *TurtleCoind) GetBlockHeaderByHeight(height int) (map[string]interface{}, error) {
	daemon.check()
	params := make(map[string]interface{})
	params["height"] = height
	resp, err := daemon.makePostRequest("getblockheaderbyheight", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}
