package blockchain

import (
	"strings"

	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

func GetWallet(api ton.APIClientWrapped, seed string, ver wallet.Version) (*wallet.Wallet, error) {
	words := strings.Split(seed, " ")

	w, err := wallet.FromSeed(api, words, ver)
	if err != nil {
		return nil, err
	}

	return w, nil
}
