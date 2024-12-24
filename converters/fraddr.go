package converters

import "github.com/xssnick/tonutils-go/address"

func FrAddr(addr *address.Address, testnet, bounceable bool) string {
	return addr.Testnet(testnet).Bounce(bounceable).String()
}
