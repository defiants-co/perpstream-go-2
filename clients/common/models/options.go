package models

import (
	"encoding/json"
	"time"
)

type Option struct {
	Position
	// Comparable Information
	Strategy   string `json:"strategy"`
	Asset      string `json:"asset"`
	IsInversed bool   `json:"is_inversed"`

	// Size Information
	Amount    float64 `json:"amount"`
	AmountUsd float64 `json:"amount_usd"`

	// Price Information
	PremiumPaid float64 `json:"premium_paid"`

	// Profit Information
	PnlUsd     float64 `json:"pnl_usd"`
	PnlPercent float64 `json:"pnl_percent"`

	// Time Information
	PurchaseDate   time.Time `json:"purchase_date"`
	ExpirationDate time.Time `json:"expiration_date"`
	DurationDays   int       `json:"duration_days"`

	// TODO: add strike prices
	Strikes []Strike `json:"strikes"`
}

type Strike struct {
	Type        string  `json:"type"`
	IsBought    bool    `json:"is_bought"`
	StrikePrice float64 `json:"strike_price"`
}

func (f Option) MarshalJSON() ([]byte, error) {
	type Alias Option
	return json.Marshal(&struct {
		*Alias
		Position Position `json:"-"`
	}{
		Alias: (*Alias)(&f),
	})
}

func (o *Option) Equal(other *Option) bool {
	return o.Strategy == other.Strategy &&
		o.Asset == other.Asset &&
		o.IsInversed == other.IsInversed &&
		o.Amount == other.Amount &&
		// o.AmountUsd == other.AmountUsd &&
		o.PremiumPaid == other.PremiumPaid &&
		// o.PnlUsd == other.PnlUsd &&
		// o.PnlPercent == other.PnlPercent &&
		// o.PurchaseDate.Equal(other.PurchaseDate) &&
		o.ExpirationDate.Equal(other.ExpirationDate) &&
		o.DurationDays == other.DurationDays
}

func NewOption(
	strategy,
	asset string,
	isInversed bool,
	amount,
	amountUsd,
	premiumPaid,
	pnlUsd,
	pnlPercent float64,
	purchaseDate,
	expirationDate time.Time,
	durationDays int,
	strikes []Strike,
) Option {
	return Option{
		Strategy:       strategy,
		Asset:          asset,
		IsInversed:     isInversed,
		Amount:         amount,
		AmountUsd:      amountUsd,
		PremiumPaid:    premiumPaid,
		PnlUsd:         pnlUsd,
		PnlPercent:     pnlPercent,
		PurchaseDate:   purchaseDate,
		ExpirationDate: expirationDate,
		DurationDays:   durationDays,
		Strikes:        strikes,
	}
}
