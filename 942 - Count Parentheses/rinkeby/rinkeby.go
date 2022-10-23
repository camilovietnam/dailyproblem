package rinkeby

import (
	"context"
	"fmt"
	"github.com/dailyproblem/942-count-parentheses/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"net/http"
)

const goerliNetworkURL = "..."
const passphrase = "...."
const contractAddress = "..."

func LoadContract(ctx context.Context, r *http.Request) (*contracts.Parenthesis, *bind.TransactOpts, error) {
	blockchain, err := ethclient.Dial(goerliNetworkURL)
	if err != nil {
		fmt.Printf("[❌] dial eth client %v\n", err)
		return nil, nil, err
	}
	chainID, err := blockchain.ChainID(ctx)
	if err != nil {
		fmt.Printf("[❌] read chan ID %v\n", err)
		return nil, nil, err
	}

	auth, err := bind.NewTransactorWithChainID(r.Body, passphrase, chainID)
	if err != nil {
		fmt.Printf("[❌] create new transactor: %v\n", err)
		return nil, auth, err
	}

	contract, err := contracts.NewParenthesis(common.HexToAddress(contractAddress), blockchain)
	if err != nil {
		fmt.Printf("[❌] new contract from address: %v\n", err)
		return nil, auth, err
	}

	return contract, auth, nil
}
