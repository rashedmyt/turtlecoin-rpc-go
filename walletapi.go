// Copyright (c) 2018-2019 Rashed Mohammed, The TurtleCoin Developers
// Please see the included LICENSE file for more information

package turtlecoinrpc

import (
	"errors"
	"strconv"
)

// WalletAPI structure contains the info of wallet
// URL and Port, Daemon URL and Port, and RPCPassword
type WalletAPI struct {
	URL         string
	Port        int
	DaemonURL   string
	DaemonPort  int
	DaemonSSL   bool
	RPCPassword string
}

func (wallet *WalletAPI) check() error {
	if wallet.URL == "" {
		wallet.URL = "127.0.0.1"
	}
	if wallet.Port == 0 {
		wallet.Port = 8070
	}
	if wallet.DaemonURL == "" {
		wallet.DaemonURL = "127.0.0.1"
	}
	if wallet.DaemonPort == 0 {
		wallet.DaemonPort = 11898
	}
	if wallet.RPCPassword == "" {
		return errors.New("RPCPassword not specified")
	}
	return nil
}

// <--------- Wallet Operations --------->

// CreateWallet creates a wallet with the
// specified filename and password.
func (wallet *WalletAPI) CreateWallet(
	filename string,
	password string) error {
	err := wallet.check()
	if err != nil {
		return err
	}
	if filename == "" {
		return errors.New("Filename of the wallet is required")
	}
	if password == "" {
		return errors.New("Password of the wallet is required")
	}
	params := make(map[string]interface{})
	params["daemonHost"] = wallet.DaemonURL
	params["daemonPort"] = wallet.DaemonPort
	params["daemonSSL"] = wallet.DaemonSSL
	params["filename"] = filename
	params["password"] = password

	_, err = wallet.makePostRequest("wallet/create", params)
	return err
}

// ImportKey imports a wallet with a
// private spend and view key
func (wallet *WalletAPI) ImportKey(
	filename string,
	password string,
	scanHeight int,
	spendKey string,
	viewKey string) error {
	err := wallet.check()
	if err != nil {
		return err
	}
	if filename == "" {
		return errors.New("Filename of the wallet is required")
	}
	if password == "" {
		return errors.New("Password of the wallet is required")
	}
	if spendKey == "" || len(spendKey) != 64 {
		return errors.New("Private Spend Key is invalid")
	}
	if viewKey == "" || len(viewKey) != 64 {
		return errors.New("Private View Key is invalid")
	}
	params := make(map[string]interface{})
	params["daemonHost"] = wallet.DaemonURL
	params["daemonPort"] = wallet.DaemonPort
	params["daemonSSL"] = wallet.DaemonSSL
	params["filename"] = filename
	params["password"] = password
	params["scanHeight"] = scanHeight
	params["privateSpendKey"] = spendKey
	params["privateViewKey"] = viewKey

	_, err = wallet.makePostRequest("wallet/import/key", params)
	return err
}

// ImportSeed imports a wallet using
// a mnemonic seed
func (wallet *WalletAPI) ImportSeed(
	filename string,
	password string,
	scanHeight int,
	mnemonicSeed string) error {
	err := wallet.check()
	if err != nil {
		return err
	}
	if filename == "" {
		return errors.New("Filename of the wallet is required")
	}
	if password == "" {
		return errors.New("Password of the wallet is required")
	}
	if mnemonicSeed == "" {
		return errors.New("Mnemonic Seed is invalid")
	}
	params := make(map[string]interface{})
	params["daemonHost"] = wallet.DaemonURL
	params["daemonPort"] = wallet.DaemonPort
	params["daemonSSL"] = wallet.DaemonSSL
	params["filename"] = filename
	params["password"] = password
	params["scanHeight"] = scanHeight
	params["mnemonicSeed"] = mnemonicSeed

	_, err = wallet.makePostRequest("wallet/import/seed", params)
	return err
}

// ImportViewOnly imports a wallet using
// a mnemonic seed
func (wallet *WalletAPI) ImportViewOnly(
	filename string,
	password string,
	scanHeight int,
	viewkey string,
	address string) error {
	err := wallet.check()
	if err != nil {
		return err
	}
	if filename == "" {
		return errors.New("Filename of the wallet is required")
	}
	if password == "" {
		return errors.New("Password of the wallet is required")
	}
	if viewkey == "" {
		return errors.New("Mnemonic Seed is invalid")
	}
	params := make(map[string]interface{})
	params["daemonHost"] = wallet.DaemonURL
	params["daemonPort"] = wallet.DaemonPort
	params["daemonSSL"] = wallet.DaemonSSL
	params["filename"] = filename
	params["password"] = password
	params["scanHeight"] = scanHeight
	params["privateViewKey"] = viewkey
	params["address"] = address

	_, err = wallet.makePostRequest("wallet/import/view", params)
	return err
}

// OpenWallet opens an already
// existing wallet
func (wallet *WalletAPI) OpenWallet(
	filename string,
	password string) error {
	err := wallet.check()
	if err != nil {
		return err
	}
	if filename == "" {
		return errors.New("Filename of the wallet is required")
	}
	if password == "" {
		return errors.New("Password of the wallet is required")
	}
	params := make(map[string]interface{})
	params["daemonHost"] = wallet.DaemonURL
	params["daemonPort"] = wallet.DaemonPort
	params["daemonSSL"] = wallet.DaemonSSL
	params["filename"] = filename
	params["password"] = password

	_, err = wallet.makePostRequest("wallet/open", params)
	return err
}

// CloseWallet saves and closes the
// opened wallet
func (wallet *WalletAPI) CloseWallet() error {
	err := wallet.check()
	if err != nil {
		return err
	}

	_, err = wallet.makeDeleteRequest("wallet")
	return err
}

// <--------- Address Operations --------->

// Addresses gets a list of all addresses
// in the wallet container
func (wallet *WalletAPI) Addresses() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	resp, err := wallet.makeGetRequest("addresses")
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// DeleteAddress deletes the subwallet address from
// the container. Note that you cannot delete the
// primary address (first address in the container)
func (wallet *WalletAPI) DeleteAddress(address string) error {
	err := wallet.check()
	if err != nil {
		return err
	}

	_, err = wallet.makeDeleteRequest("addresses/" + address)
	return err
}

// Primary returns the primary address. It is the first
// address created and is used as change address if not
// specified.
func (wallet *WalletAPI) Primary() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	resp, err := wallet.makeGetRequest("addresses/primary")
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// CreateAddress creates a new random address
// in the wallet container.
func (wallet *WalletAPI) CreateAddress() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	resp, err := wallet.makePostRequest("addresses/create", nil)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// ImportAddress imports a subwallet with given
// private spend key
func (wallet *WalletAPI) ImportAddress(
	scanHeight int,
	spendKey string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	if spendKey == "" {
		return nil, errors.New("Private Spend Key is required")
	}
	params := make(map[string]interface{})
	params["scanHeight"] = scanHeight
	params["privateSpendKey"] = spendKey

	resp, err := wallet.makePostRequest("addresses/import", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// ImportViewAddress imports a view only subwallet
// with the given public spend key
func (wallet *WalletAPI) ImportViewAddress(
	scanHeight int,
	spendKey string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	if spendKey == "" {
		return nil, errors.New("Public Spend Key is required")
	}
	params := make(map[string]interface{})
	params["scanHeight"] = scanHeight
	params["publicSpendKey"] = spendKey

	resp, err := wallet.makePostRequest("addresses/import/view", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// CreateIntegratedAddress creates an integrated address
// from the specified address and payment id
func (wallet *WalletAPI) CreateIntegratedAddress(
	address string,
	paymentID string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	if address == "" {
		return nil, errors.New("Address is required")
	}
	if paymentID == "" {
		return nil, errors.New("Payment ID is required")
	}

	resp, err := wallet.makeGetRequest("addresses/" + address + "/" + paymentID)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// <--------- Node Operations --------->

// GetNodeDetails gets the node address, port
// fee and fee address.
func (wallet *WalletAPI) GetNodeDetails() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	resp, err := wallet.makeGetRequest("node")
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// SetNode sets the node address and port
func (wallet *WalletAPI) SetNode(
	daemonHost string,
	daemonPort int,
	daemonSSL bool) error {
	err := wallet.check()
	if err != nil {
		return err
	}
	if daemonHost == "" {
		return errors.New("Host address is required")
	}
	if daemonPort == 0 {
		return errors.New("Host port is required")
	}
	wallet.DaemonURL = daemonHost
	wallet.DaemonPort = daemonPort
	wallet.DaemonSSL = daemonSSL
	params := make(map[string]interface{})
	params["daemonHost"] = daemonHost
	params["daemonPort"] = daemonPort
	params["daemonSSL"] = daemonSSL

	_, err = wallet.makePutRequest("node", params)
	return err
}

// <---------- Key Operations --------->

// PrivateViewKey returns the shared private view
// key of the wallet container
func (wallet *WalletAPI) PrivateViewKey() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	resp, err := wallet.makeGetRequest("keys")
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// Keys returns the public and private
// key of the given address
func (wallet *WalletAPI) Keys(address string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	if address == "" {
		return nil, errors.New("Address is required")
	}

	resp, err := wallet.makeGetRequest("keys/" + address)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// MnemonicSeed return the mnemonic seed
// for the given address if possible
func (wallet *WalletAPI) MnemonicSeed(address string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	if address == "" {
		return nil, errors.New("Address is required")
	}

	resp, err := wallet.makeGetRequest("keys/mnemonic/" + address)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// <--------- Balance Operations --------->

// TotalBalance returns the total balance of
// the wallet container
func (wallet *WalletAPI) TotalBalance() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	resp, err := wallet.makeGetRequest("balance")
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// Balance returns the balance of specific
// address in the wallet container
func (wallet *WalletAPI) Balance(address string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	if address == "" {
		return nil, errors.New("Address is required")
	}

	resp, err := wallet.makeGetRequest("balance/" + address)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// Balances returns the balance of all
// addresses in the wallet container
func (wallet *WalletAPI) Balances() ([]map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	resp, err := wallet.makeGetRequest("balances")
	if resp == nil {
		return nil, err
	}

	temp := resp.([]interface{})
	response := make([]map[string]interface{}, len(temp))
	for i, obj := range temp {
		response[i] = obj.(map[string]interface{})
	}
	return response, err
}

// <--------- Miscellaneous Operations --------->

// Save saves the current wallet state
func (wallet *WalletAPI) Save() error {
	err := wallet.check()
	if err != nil {
		return err
	}

	_, err = wallet.makePutRequest("save", nil)
	return err
}

// Reset method resets and saves the wallet,
// beginning scanning from the height given.
func (wallet *WalletAPI) Reset(scanHeight int) error {
	err := wallet.check()
	if err != nil {
		return nil
	}

	params := make(map[string]interface{})
	params["scanHeight"] = scanHeight

	_, err = wallet.makePutRequest("reset", params)
	return err
}

// ValidateAddress method validates an address for
// TRTL compatibility and returns the address break-down
func (wallet *WalletAPI) ValidateAddress(address string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}
	if address == "" {
		return nil, errors.New("Address is required")
	}

	params := make(map[string]interface{})
	params["address"] = address

	resp, err := wallet.makePostRequest("addresses/validate", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// Status method returns the current sync
// status of the wallet container
func (wallet *WalletAPI) Status() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	resp, err := wallet.makeGetRequest("status")
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// <--------- Transaction Operations --------->

// Transactions return a list of transactions
// within the range of startHeight and endHeight.
// If endHeight is less than startHeight then it
// returns all transactions from the startHeight
// for 1000 blocks.
func (wallet *WalletAPI) Transactions(startHeight int, endHeight int) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	method := "transactions"

	if startHeight != 0 {
		method += "/" + strconv.Itoa(startHeight)
		if endHeight != 0 && endHeight > startHeight {
			method += "/" + strconv.Itoa(endHeight)
		}
	}

	resp, err := wallet.makeGetRequest(method)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// GetTransactionDetails returns the details of
// the given transaction hash if it is present
// in the wallet, else error occurs
func (wallet *WalletAPI) GetTransactionDetails(hash string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	resp, err := wallet.makeGetRequest("transactions/hash/" + hash)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// UnconfirmedTransactions returns the list of
// unconfirmed transactions for the specified
// address. If the address is empty then it
// returns all the unconfirmed transactions in
// the wallet container.
func (wallet *WalletAPI) UnconfirmedTransactions(address string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	method := "transactions/unconfirmed"
	if address != "" {
		method += "/" + address
	}

	resp, err := wallet.makeGetRequest(method)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// TransactionsByAddress returns list of transactions
// corresponding to that address. If endHeight is less
// than startHeight then it returns all transactions
// from startHeight for 1000 blocks.
func (wallet *WalletAPI) TransactionsByAddress(
	address string,
	startHeight int,
	endHeight int) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	if address == "" {
		return nil, errors.New("Address is required")
	}

	method := "transactions/address" + "/" + address + "/" + strconv.Itoa(startHeight)

	if endHeight != 0 && endHeight > startHeight {
		method += "/" + strconv.Itoa(endHeight)
	}

	resp, err := wallet.makeGetRequest(method)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// TransactionPrivateKey returns the private key
// of the transaction for auditing purposes.
func (wallet *WalletAPI) TransactionPrivateKey(hash string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	if hash == "" {
		return nil, errors.New("Transaction hash is required")
	}

	resp, err := wallet.makeGetRequest("transactions/privatekey/" + hash)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// SendBasicTransaction sends the specified amount
// to the specified address with the specified paymentID.
func (wallet *WalletAPI) SendBasicTransaction(
	destinationAddress string,
	amount int,
	paymentID string) (map[string]interface{}, error) {
	err := wallet.check()

	if destinationAddress == "" {
		return nil, errors.New("Destination Address is required")
	}

	if amount == 0 {
		return nil, errors.New("Amount must be greater than 0")
	}

	params := make(map[string]interface{})
	params["destination"] = destinationAddress
	params["amount"] = amount

	if paymentID != "" {
		params["paymentID"] = paymentID
	}

	resp, err := wallet.makePostRequest("transactions/send/basic", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// SendAdvancedTransaction sends the specified amount
// to the specified address with the specified paymentID.
func (wallet *WalletAPI) SendAdvancedTransaction(
	destinations []map[string]interface{},
	mixin int,
	fee int,
	sourceAddresses []string,
	paymentID string,
	changeAddress string,
	unlockTime int) (map[string]interface{}, error) {
	err := wallet.check()

	for _, obj := range destinations {
		if obj["address"] == "" {
			return nil, errors.New("Address is required in every destination")
		}

		if obj["amount"] == 0 {
			return nil, errors.New("Amount must be greater than 0 in every destination")
		}
	}

	params := make(map[string]interface{})
	params["destinations"] = destinations
	params["fee"] = fee
	params["unlockTime"] = unlockTime
	params["mixin"] = mixin

	if paymentID != "" {
		params["paymentID"] = paymentID
	}
	if len(sourceAddresses) != 0 {
		params["sourceAddresses"] = sourceAddresses
	}
	if changeAddress != "" {
		params["changeAddress"] = changeAddress
	}

	resp, err := wallet.makePostRequest("transactions/send/basic", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// SendBasicFusion sends a single fusion transaction
// if it can and returns the transaction hash
func (wallet *WalletAPI) SendBasicFusion() (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	resp, err := wallet.makePostRequest("transactions/send/fusion/basic", nil)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}

// SendAdvancedFusion sends a single fusion transaction
// if it can and returns the transaction hash
func (wallet *WalletAPI) SendAdvancedFusion(
	mixin int,
	sourceAddress []string,
	destinationAddress string) (map[string]interface{}, error) {
	err := wallet.check()
	if err != nil {
		return nil, err
	}

	if destinationAddress == "" {
		return nil, errors.New("Destination Address is required")
	}

	params := make(map[string]interface{})
	params["destination"] = destinationAddress
	params["mixin"] = mixin

	if len(sourceAddress) != 0 {
		params["sourceAddresses"] = sourceAddress
	}

	resp, err := wallet.makePostRequest("transactions/send/fusion/basic", params)
	if resp != nil {
		return resp.(map[string]interface{}), err
	}

	return nil, err
}
