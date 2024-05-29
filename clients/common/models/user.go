package models

import "time"

type User struct {
	Id          string    `json:"id"`
	PnlUsd      float64   `json:"pnl_usd"`
	PnlPercent  float64   `json:"pnl_percent"`
	WinRate     float64   `json:"win_rate"`
	TotalVolume float64   `json:"total_volume"`
	LastBuy     time.Time `json:"last_active"`
}

func NewUser(id string, pnlUsd, pnlPercent, winRate, totalVolume float64, lastActive time.Time) User {
	return User{
		Id:          id,
		PnlUsd:      pnlUsd,
		PnlPercent:  pnlPercent,
		WinRate:     winRate,
		TotalVolume: totalVolume,
		LastBuy:     lastActive,
	}
}
