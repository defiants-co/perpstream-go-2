package gmx

import (
	"errors"
	"net/http"
	"sort"
	"time"

	"github.com/defiants-co/perpstream-go-2/clients/common"
	"github.com/defiants-co/perpstream-go-2/clients/common/models"
	"github.com/defiants-co/perpstream-go-2/clients/common/utils"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/hasura/go-graphql-client"
)

type GmxClient struct {
	common.FuturesClient
	cache          *utils.ConcurrentMap[[]models.Future]
	prices         *priceCache
	pool           *gmxConnectionPool
	subgraphClient *graphql.Client
	activeStreams  utils.ConcurrentMap[chan int]
}

func NewGmxClient(rpcs []string) (*GmxClient, error) {
	cp, err := newGmxConnectionPool(rpcs)
	if err != nil {
		return nil, err
	}
	cache := utils.NewConcurrentMap[[]models.Future]()
	pc := newPriceCache()

	client := graphql.NewClient(subgraphUrl, &http.Client{})

	return &GmxClient{
		cache:          cache,
		prices:         pc,
		pool:           cp,
		subgraphClient: client,
		activeStreams:  *utils.NewConcurrentMap[chan int](),
	}, nil
}

func (c *GmxClient) GetLeaderboard(limit int) ([]models.User, error) {
	leaderboard, err := getLeaderboardReq(c.subgraphClient)
	if err != nil {
		return nil, err
	}

	users := make([]models.User, len(*leaderboard))

	for i, subgraphUser := range *leaderboard {
		users[i] = periodStatsToUser(&subgraphUser)
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].PnlUsd > users[j].PnlUsd
	})
	if limit > len(users) {
		limit = len(users)
	}
	return users[0:limit], nil
}

func (c *GmxClient) FetchPositions(userId string) ([]models.Future, error) {
	if !ethCommon.IsHexAddress(userId) {
		return nil, errors.New(userId + " is not a valid hex address")
	}
	address := ethCommon.HexToAddress(userId)

	abiPositions := c.pool.getPositions(address, 1)

	positions := make([]models.Future, len(abiPositions))

	for i, ap := range abiPositions {
		positions[i] = *gmxToFuturesPosition(ap, c.prices)
	}

	return positions, nil
}

func (c *GmxClient) fetchRetry(userId string) []models.Future {
	positions, err := c.FetchPositions(userId)
	if err == nil {
		return positions
	}

	return c.fetchRetry(userId)
}

func (c *GmxClient) StartPositionStream(userId string, callback common.FuturesPositionCallback, waitSeconds float64) error {
	if !ethCommon.IsHexAddress(userId) {
		return errors.New(userId + " is not a valid hex address")
	}

	if c.activeStreams.Contains(userId) {
		return errors.New(userId + " is already being streamed")
	}

	cancelChan := make(chan int)

	c.activeStreams.Set(userId, cancelChan)

	lastPositions := c.fetchRetry(userId)

	go callback(lastPositions, lastPositions, true, dataSourceName, userId)

	for {
		select {
		case <-cancelChan:
			c.activeStreams.Delete(userId)
			return nil

		case <-time.After(time.Duration(waitSeconds) * time.Second):
			newPositions := c.fetchRetry(userId)
			if !utils.FuturePositionSetsAreEqual(lastPositions, newPositions) {
				go callback(lastPositions, newPositions, false, dataSourceName, userId)
				lastPositions = newPositions
			}
		}
	}

}

func (c *GmxClient) CancelPositionStream(userId string) error {
	if !c.activeStreams.Contains(userId) {
		return errors.New(userId + " is not streaming")
	}

	cancelChan, _ := c.activeStreams.Get(userId)
	cancelChan <- 1

	return nil
}

func (c *GmxClient) StartOrderFlowStream(callback common.FuturesOrderCallback) error {
	return nil
}

func (c *GmxClient) CancelOrderFlowStream() error {
	return nil
}
