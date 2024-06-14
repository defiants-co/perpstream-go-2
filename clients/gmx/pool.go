package gmx

import (
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/defiants-co/perpstream-go-2/clients/common/utils"
	abis "github.com/defiants-co/perpstream-go-2/clients/gmx/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type gmxConnectionPool struct {
	connections  []*abis.MainCaller
	currentIndex int
	mu           sync.Mutex
}

func newGmxConnectionPool(rpcUrls []string) (*gmxConnectionPool, error) {
	callers := returnCallers(rpcUrls)
	if len(callers) == 0 {
		return nil, errors.New("no rpcs provided")
	}

	return &gmxConnectionPool{
		connections:  callers,
		currentIndex: 0,
	}, nil
}

func returnCallers(rpcUrls []string) []*abis.MainCaller {
	var callers []*abis.MainCaller
	readerAddress := common.HexToAddress(gmxReaderContractAddress)
	for _, url := range rpcUrls {
		web3, err := ethclient.Dial(url)
		if err != nil {
			continue
		}
		client, err := abis.NewMainCaller(readerAddress, web3)
		if err != nil {
			continue
		}
		callers = append(callers, client)
	}

	return callers
}

func (p *gmxConnectionPool) numCallers() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return len(p.connections)
}

func (p *gmxConnectionPool) nextCaller() *abis.MainCaller {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.connections) == 0 {
		return nil
	}

	caller := p.connections[p.currentIndex]
	p.currentIndex = (p.currentIndex + 1) % len(p.connections)
	return caller
}

func (p *gmxConnectionPool) getPositions(userAddress common.Address, waitSeconds float64) []abis.PositionProps {
	dataStoreAddress := common.HexToAddress(gmxDataStoreContractAddress)

	for {
		caller := p.nextCaller()
		if caller == nil {
			time.Sleep(time.Duration(waitSeconds) * time.Second)
			continue
		}

		positionProps, err := caller.GetAccountPositions(
			&bind.CallOpts{Pending: false},
			dataStoreAddress,
			userAddress,
			&big.Int{},
			utils.MaxBigint(),
		)
		if err == nil {
			return positionProps
		}
		time.Sleep(time.Duration(waitSeconds) * time.Second)
	}
}
