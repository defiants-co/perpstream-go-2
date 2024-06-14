package hegic

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/defiants-co/perpstream-go-2/clients/common"
	"github.com/defiants-co/perpstream-go-2/clients/common/models"
	"github.com/defiants-co/perpstream-go-2/clients/common/utils"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

type HegicClient struct {
	activeOptions   *utils.ConcurrentMap[[]models.Option]
	inactiveOptions *utils.ConcurrentMap[[]models.Option]
	activeStreams   *utils.ConcurrentMap[chan struct{}]
	orderFlowChan   chan []callbackPayload
}

func NewHegicClient(fetchOnInit bool) (*HegicClient, error) {
	client := &HegicClient{
		activeOptions:   utils.NewConcurrentMap[[]models.Option](),
		inactiveOptions: utils.NewConcurrentMap[[]models.Option](),
		activeStreams:   utils.NewConcurrentMap[chan struct{}](),
		orderFlowChan:   nil,
	}
	if fetchOnInit {
		err := client.updateCache()
		if err != nil {
			return nil, err
		}
	} else {
		go client.updateCache()
	}

	return client, nil
}

func (h *HegicClient) StreamUpdates(waitSeconds float64) {
	go func() {
		for {
			h.updateCache()
			time.Sleep(time.Duration(waitSeconds) * time.Second)
		}
	}()
}

func (h *HegicClient) updateCache() error {
	active, inactive, err := getOptions()
	if err != nil {
		return errors.New("error fetching options: " + err.Error())
	}

	if h.orderFlowChan != nil {
		difference := getDifference(h.activeOptions.Values(), active)
		if len(difference) > 0 {
			h.orderFlowChan <- difference
		}
	}

	h.activeOptions.SetAll(active)
	h.inactiveOptions.SetAll(inactive)
	return nil
}

func (hc *HegicClient) GetLeaderboard(limit int) ([]models.User, error) {
	var leaderboard []hegicUserStats

	err := utils.Fetch("leaderboard", apiUrl, nil, &leaderboard)

	if err != nil {
		return nil, err
	}

	var users []models.User

	for _, stat := range leaderboard {
		users = append(users, resolveUser(stat, hc.inactiveOptions))
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].PnlUsd > users[j].PnlUsd
	})

	if len(users) < limit {
		limit = len(users)
	}

	return users[0:limit], nil
}

func (hc *HegicClient) FetchPositions(userId string) ([]models.Option, error) {
	if !ethCommon.IsHexAddress(userId) {
		return nil, errors.New(userId + " is not a valid hex address")
	}

	positions, exists := hc.activeOptions.Get(userId)
	if !exists {
		return make([]models.Option, 0), nil
	}

	return positions, nil
}

func (hc *HegicClient) StartPositionStream(userId string, callback common.OptionsPositionCallback, waitSeconds float64) error {
	lastPositions, err := hc.FetchPositions(userId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if hc.activeStreams.Contains(userId) {
		return errors.New(userId + "is already streaming")
	}

	cancelChan := make(chan struct{})
	hc.activeStreams.Set(userId, cancelChan)

	go callback(lastPositions, lastPositions, true, dataSourceName, userId)

	for {
		select {
		case <-cancelChan:
			hc.activeStreams.Delete(userId)
			return nil

		default:
			newPositions, _ := hc.FetchPositions(userId)
			if !utils.OptionPositionSetsAreEqual(lastPositions, newPositions) {
				go callback(lastPositions, newPositions, false, dataSourceName, userId)
				lastPositions = newPositions
			}
			if waitSeconds > 0 {
				time.Sleep(time.Duration(waitSeconds) * time.Second)
			}
		}
	}
}

func (hc *HegicClient) CancelPositionStream(userId string) error {

	if !hc.activeStreams.Contains(userId) {
		return errors.New(userId + " is not streaming")
	}

	cancelChan, _ := hc.activeStreams.Get(userId)
	cancelChan <- struct{}{}

	return nil
}

func (hc *HegicClient) StartOrderFlowStream(callback common.OptionsOrderCallback) error {
	// Implementation for starting the order flow stream
	if hc.orderFlowChan != nil {
		return errors.New("stream already running")
	}

	hc.orderFlowChan = make(chan []callbackPayload) // Initialize the channel

	go func() {
		for {
			select {
			case o := <-hc.orderFlowChan:
				if o == nil {
					hc.orderFlowChan = nil
					fmt.Println("cancelling")
					return
				} else {
					for _, i := range o {
						go callback(i.Option, dataSourceName, i.UserId)
					}
				}
			}
		}
	}()
	return nil
}

func (hc *HegicClient) CancelOrderFlowStream() error {
	if hc.orderFlowChan == nil {
		return errors.New("stream is not running")
	}

	close(hc.orderFlowChan) // Close the channel to signal cancellation
	hc.orderFlowChan = nil

	return nil
}

func (hc *HegicClient) MockRequest() {
	if hc.orderFlowChan != nil {
		hc.orderFlowChan <- []callbackPayload{
			{
				Option: models.Option{
					Strategy:       "Call",
					Asset:          "ETH",
					IsInversed:     false,
					Amount:         1.0,
					AmountUsd:      3000.0,
					PremiumPaid:    100.0,
					PnlUsd:         200.0,
					PnlPercent:     6.67,
					PurchaseDate:   time.Now().AddDate(0, 0, -1),
					ExpirationDate: time.Now().AddDate(0, 0, 29),
					DurationDays:   30,
				},
				UserId: "amrith",
			},
			{
				Option: models.Option{
					Strategy:       "Put",
					Asset:          "BTC",
					IsInversed:     false,
					Amount:         0.5,
					AmountUsd:      25000.0,
					PremiumPaid:    500.0,
					PnlUsd:         1500.0,
					PnlPercent:     6.0,
					PurchaseDate:   time.Now().AddDate(0, 0, -2),
					ExpirationDate: time.Now().AddDate(0, 0, 28),
					DurationDays:   30,
				},
				UserId: "john_doe",
			},
		}
	}
}
