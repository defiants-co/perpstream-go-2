package models

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

type Future struct {
	Position
	// Comparable Information
	Market          string `json:"market"`
	CollateralToken string `json:"collateral_token"`
	IsLong          bool   `json:"is_long"`

	// Size Information
	CollateralAmount    float64 `json:"collateral_amount"`
	CollateralAmountUsd float64 `json:"collateral_amount_usd"`
	PositionSize        float64 `json:"position_size"`
	PositionSizeUsd     float64 `json:"position_size_usd"`
	Leverage            float64 `json:"leverage"`

	// Price Information
	EntryPrice float64 `json:"entry_price"`
	MarkPrice  float64 `json:"mark_price"`

	// Profit Information
	PnlUsd     float64 `json:"pnl_usd"`
	PnlPercent float64 `json:"pnl_percent"`

	// Time Information
	LastUpdated time.Time `json:"last_updated"`
}

func (f Future) MarshalJSON() ([]byte, error) {
	type Alias Future
	return json.Marshal(&struct {
		*Alias
		UserId string `json:"-"`
	}{
		Alias: (*Alias)(&f),
	})
}
func comparablePosition(position1 Future, position2 Future) bool {
	return (position1.CollateralToken == position2.CollateralToken &&
		position1.Market == position2.Market &&
		position1.IsLong == position2.IsLong)
}

func (fp *Future) Equal(other *Future) bool {
	return (comparablePosition(*fp, *other) &&
		math.Abs(fp.CollateralAmount-other.CollateralAmount) < 0.05*fp.CollateralAmount &&
		math.Abs(fp.EntryPrice-other.EntryPrice) < 0.1*fp.EntryPrice &&
		math.Abs(fp.Leverage-other.Leverage) < 0.05*fp.Leverage &&
		math.Abs(fp.PositionSize-other.PositionSize) < 0.01*fp.PositionSize)
}

func NewFuture(market, collateralToken string, isLong bool, collateralAmount, collateralAmountUsd, positionSize, positionSizeUsd, leverage, entryPrice, markPrice, pnlUsd, pnlPercent float64, lastUpdated time.Time) Future {
	return Future{
		Market:              market,
		CollateralToken:     collateralToken,
		IsLong:              isLong,
		CollateralAmount:    collateralAmount,
		CollateralAmountUsd: collateralAmountUsd,
		PositionSize:        positionSize,
		PositionSizeUsd:     positionSizeUsd,
		Leverage:            leverage,
		EntryPrice:          entryPrice,
		MarkPrice:           markPrice,
		PnlUsd:              pnlUsd,
		PnlPercent:          pnlPercent,
		LastUpdated:         lastUpdated,
	}
}

func (f Future) String() string {
	return fmt.Sprintf("Market: %s,\nCollateralToken: %s,\nIsLong: %t,\nCollateralAmount: %f,\nCollateralAmountUsd: %f,\nPositionSize: %f,\nPositionSizeUsd: %f,\nLeverage: %f,\nEntryPrice: %f,\nMarkPrice: %f,\nPnlUsd: %f,\nPnlPercent: %f,\nLastUpdated: %s",
		f.Market, f.CollateralToken, f.IsLong, f.CollateralAmount, f.CollateralAmountUsd, f.PositionSize, f.PositionSizeUsd, f.Leverage, f.EntryPrice, f.MarkPrice, f.PnlUsd, f.PnlPercent, f.LastUpdated.Format(time.RFC3339))
}
