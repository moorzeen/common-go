package blockchain

import "github.com/xssnick/tonutils-go/ton"

type TonClients struct {
	Mainnet ton.APIClientWrapped
	Testnet ton.APIClientWrapped
}
