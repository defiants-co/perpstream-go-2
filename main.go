package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/defiants-co/perpstream-go-2/clients/common"
	"github.com/defiants-co/perpstream-go-2/clients/common/models"
	"github.com/defiants-co/perpstream-go-2/clients/hegic"
)

func main() {
	hc, err := hegic.NewHegicClient()
	if err != nil {
		panic(err)
	}

	hc.StreamUpdates(10)
	hc.StartOrderFlowStream(Callback2)

	select {}

	// var options []models.Option

	// leader, _ := hc.GetLeaderboard(10000)
	// for _, x := range leader {
	// 	p, _ := hc.FetchPositions(x.Id)
	// 	options = append(options, p...)
	// }

	// writeToFile[models.Option](options, "options.json")

	// var Rpcs []string = []string{
	// 	"https://arbitrum.llamarpc.com",
	// 	"https://rpc.ankr.com/arbitrum",
	// 	"https://arbitrum-one-rpc.publicnode.com",
	// 	"https://arb-pokt.nodies.app",
	// 	"https://arb-mainnet-public.unifra.io",
	// 	"https://arbitrum.meowrpc.com",
	// 	"https://arbitrum-one.publicnode.com",
	// 	"https://arbitrum.rpc.subquery.network/public",
	// 	"https://1rpc.io/arb",
	// }

	// gc, _ := gmx.NewGmxClient(Rpcs)

	// var wg sync.WaitGroup

	// leaderBoard, _ := gc.GetLeaderboard(10000)
	// var filteredLeaderBoard []models.User
	// for _, l := range leaderBoard {
	// 	if l.PnlUsd > 10_000 && l.PnlPercent > 100 {
	// 		filteredLeaderBoard = append(filteredLeaderBoard, l)
	// 	}
	// }
	// var userCount int = len(filteredLeaderBoard)
	// positionList := utils.NewConcurrentList[models.Future]()
	// for _, l := range filteredLeaderBoard {
	// 	wg.Add(1)
	// 	go func(l models.User) {
	// 		defer wg.Done()
	// 		defer func() {
	// 			userCount--
	// 			fmt.Printf("Users left: %d\n", userCount)
	// 		}()
	// 		positions, err := gc.FetchPositions(l.Id)
	// 		if err == nil {
	// 			positionList.Append(positions...)
	// 		}
	// 	}(l)
	// }
	// wg.Wait()
	// writeToFile[models.Future](positionList.GetCopy(), "futures.json")
}

var Callback common.FuturesPositionCallback = func(oldPositions []models.Future, newPositions []models.Future, isInitCallback bool, dataSource string, userId string) {
	fmt.Println("started stream " + userId)
}

var Callback2 common.OptionsOrderCallback = func(order models.Option, dataSource string, userId string) {
	fmt.Println(userId, order.Asset)
}

var Callback3 common.OptionsPositionCallback = func(old []models.Option, new []models.Option, init bool, data string, user string) {
	fmt.Println("started stream " + user)
}

func writeToFile[T any](data []T, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	errCh := make(chan error, 1)

	go func() {
		errCh <- encoder.Encode(data)
	}()

	select {
	case err = <-errCh:
		if err != nil {
			return err
		}
	}

	return nil
}
