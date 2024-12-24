package blockchain

import (
	"context"
	"errors"
	"fmt"

	"github.com/moorzeen/common-go/converters"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

const mainnetConfigURL = "https://ton-blockchain.github.io/global.config.json"
const testnetConfigURL = "https://ton.org/testnet-global.config.json"

type LiteserverEndpoint struct {
	IP   int    `yaml:"ip"`
	Port int    `yaml:"port"`
	Key  string `yaml:"key"`
}

func NewAPI(testnet bool, endpoint ...*LiteserverEndpoint) (ton.APIClientWrapped, error) {
	if testnet {
		return NewTestnetAPI()
	}

	return nil, errors.New("only testnet return implemented")
}

func NewTestnetAPI() (ton.APIClientWrapped, error) {
	ctx := context.Background()

	testnetCfg, err := liteclient.GetConfigFromUrl(ctx, testnetConfigURL)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch testnet config: %w", err)
	}

	client := liteclient.NewConnectionPool()

	err = client.AddConnectionsFromConfig(ctx, testnetCfg)
	if err != nil {
		return nil, fmt.Errorf("unable to add testnet connections: %w", err)
	}

	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(testnetCfg)

	return api, nil
}

func NewMainnetAPI(endpoint *LiteserverEndpoint) (ton.APIClientWrapped, error) {
	ctx := context.Background()

	client := liteclient.NewConnectionPool()

	addr := fmt.Sprintf("%s:%d", converters.StrIP(endpoint.IP), endpoint.Port)

	err := client.AddConnection(ctx, addr, endpoint.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to add mainnet connection to pool: %w", err)
	}

	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()

	globalConfig, err := liteclient.GetConfigFromUrl(ctx, mainnetConfigURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mainnet config: %w", err)
	}

	api.SetTrustedBlockFromConfig(globalConfig)

	return api, nil
}
