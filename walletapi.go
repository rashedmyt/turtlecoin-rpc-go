// Copyright (c) 2018-2019 Rashed Mohammed, The TurtleCoin Developers
// Please see the included LICENSE file for more information

package turtlecoinrpc

import (
	"errors"
)

// WalletAPI structure contains the info of wallet
// URL and Port, Daemon URL and Port, and RPCPassword
type WalletAPI struct {
	URL         string
	Port        int
	DaemonURL   string
	DaemonPort  int
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
