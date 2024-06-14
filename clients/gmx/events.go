package gmx

import (
	"errors"
	"fmt"

	"github.com/defiants-co/perpstream-go-2/clients/common/utils"
	abis "github.com/defiants-co/perpstream-go-2/clients/gmx/abi"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type GmxEventClient struct {
	eventCaller *abis.AbisFilterer
}

// func NewGmxEventClient(rpcWebsocketUrl string) (*gmxEventClient, error) {
func NewGmxEventClient(rpcWebsocketUrl string) (*GmxEventClient, error) {
	if !utils.IsValidWSSUrl(rpcWebsocketUrl) {
		return nil, errors.New("Websocket client required")
	}

	client, err := ethclient.Dial(rpcWebsocketUrl)
	if err != nil {
		return nil, errors.New("Websocket error: " + err.Error())
	}

	emitterAddress := ethCommon.HexToAddress(gmxEventEmitterContract)

	contractClient, err := abis.NewAbisFilterer(emitterAddress, client)
	if err != nil {
		return nil, errors.New("issue building filterer")
	}

	// return &gmxEventClient{eventCaller: contractClient}, nil
	return &GmxEventClient{eventCaller: contractClient}, nil
}

func (c *GmxEventClient) StartStreaming() error {
	logs := make(chan *abis.AbisEventLog)
	fmt.Println("starting")
	sub, err := c.eventCaller.WatchEventLog(nil, logs, nil)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	for {
		select {

		case err := <-sub.Err():
			fmt.Println(err)
			return err
		case log := <-logs:
			// Process the log
			fmt.Println(log)
		}
	}
}
