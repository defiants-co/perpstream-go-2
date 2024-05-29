package common

import "github.com/defiants-co/perpstream-go-2/clients/common/models"

// Client is a generic interface for different types of trading clients.
type Client[T any] interface {
	GetLeaderboard(limit int) ([]models.User, error)
	FetchPositions(userId string) ([]T, error)
	StartPositionStream(userId string, callback StreamPositionsCallback[T], waitSeconds float64) error
	CancelPositionStream(userId string) error
	StartOrderFlowStream(callback StreamOrdersCallback[T]) error
	CancelOrderFlowStream() error
}

// StreamCallback defines the callback function for position streams.
type StreamPositionsCallback[T any] func(
	oldPositions []T,
	newPositions []T,
	isInitCallback bool,
	dataSource string,
	userId string,
)

// StreamOrdersCallback defines the callback function for order flow streams.
type StreamOrdersCallback[T any] func(
	position T,
	dataSource string,
	userId string,
)

// FuturesClient represents a client for futures trading.
type FuturesClient = Client[models.Future]

// OptionsClient represents a client for options trading.
type OptionsClient = Client[models.Option]

type FuturesPositionCallback = StreamPositionsCallback[models.Future]

type OptionsPositionCallback = StreamPositionsCallback[models.Option]

type FuturesOrderCallback = StreamOrdersCallback[models.Future]

type OptionsOrderCallback = StreamOrdersCallback[models.Option]
