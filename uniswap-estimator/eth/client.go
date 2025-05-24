package eth

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	uniswap "uniswap-estimator/abi"
)

type EthClient struct {
	rpc *ethclient.Client
}

func NewEthClient(rpcURL string) (*EthClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return &EthClient{rpc: client}, nil
}

func (ec *EthClient) GetReserves(pool, src, dst string) (reserveIn, reserveOut *big.Int, err error) {
	contract, err := uniswap.NewUniswap(common.HexToAddress(pool), ec.rpc)
	if err != nil {
		return nil, nil, err
	}

	reserves, err := contract.GetReserves(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return nil, nil, err
	}

	token0, err := contract.Token0(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return nil, nil, err
	}
	token1, err := contract.Token1(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return nil, nil, err
	}

	srcAddr := strings.ToLower(src)
	dstAddr := strings.ToLower(dst)

	token0Addr := strings.ToLower(token0.Hex())
	token1Addr := strings.ToLower(token1.Hex())

	if token0Addr == srcAddr && token1Addr == dstAddr {
		return reserves.Reserve0, reserves.Reserve1, nil
	}
	if token1Addr == srcAddr && token0Addr == dstAddr {
		return reserves.Reserve1, reserves.Reserve0, nil
	}

	return nil, nil, errors.New("token pair mismatch in pool")
}
