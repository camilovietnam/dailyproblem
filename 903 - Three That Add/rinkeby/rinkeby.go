package rinkeby

import (
	"errors"
	"fmt"
	"github.com/camilovietnam/threethatadd/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"net/http"
)

const rinkebyNetworkURL = "fill this field ..."
const contractAddress = "fill this field ..."
const keyPassword = "fill this field ..."

func LoadContract(r *http.Request) (*contracts.ThreeThatAdd, *bind.TransactOpts, error) {
	blockchain, err := ethclient.Dial(rinkebyNetworkURL)
	if err != nil {
		fmt.Printf("[❌] dial eth client %v\n", err)
		return nil, nil, errors.New("internal server error")
	}

	chainID, err := blockchain.ChainID(r.Context())
	if err != nil {
		fmt.Printf("[❌] read chan ID %v\n", err)
		return nil, nil, errors.New("internal server error")
	}

	auth, err := bind.NewTransactorWithChainID(r.Body, keyPassword, chainID)
	if err != nil {
		fmt.Printf("[❌] create new transactor: %v\n", err)
		return nil, nil, errors.New("internal server error")
	}

	contract, err := contracts.NewThreeThatAdd(common.HexToAddress(contractAddress), blockchain)
	if err != nil {
		fmt.Printf("[❌] new contract from address: %v\n", err)
		return nil, nil, errors.New("internal server error")
	}

	return contract, auth, nil
}
