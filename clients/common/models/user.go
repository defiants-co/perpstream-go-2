package models

type User struct {
	Id          string  `json:"id"`
	PnlUsd      float64 `json:"pnl_usd"`
	PnlPercent  float64 `json:"pnl_percent"`
	WinRate     float64 `json:"win_rate"`
	TotalVolume float64 `json:"total_volume"`
}

func NewUser(id string, pnlUsd, pnlPercent, winRate, totalVolume float64) User {
	return User{
		Id:          id,
		PnlUsd:      pnlUsd,
		PnlPercent:  pnlPercent,
		WinRate:     winRate,
		TotalVolume: totalVolume,
	}
}
