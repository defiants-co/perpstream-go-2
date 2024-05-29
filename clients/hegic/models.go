package hegic

import (
	"time"

	"github.com/defiants-co/perpstream-go-2/clients/common/models"
)

type hegicProjectedPnl struct {
	ProjectedPnlCurrentEpoch                           []hegicScenarioData `json:"projectedPnl_CurrentEpoch"`
	ManualRiskAnalysis                                 []hegicScenarioData `json:"manualRiskAnalysis"`
	ProjectedPnl7_14                                   []hegicScenarioData `json:"projectedPnl_7_14"`
	ProjectedPnl14_30                                  []hegicScenarioData `json:"projectedPnl_14_30"`
	ProjectedPnl30_60                                  []hegicScenarioData `json:"projectedPnl_30_60"`
	ProjectedPnl60_90                                  []hegicScenarioData `json:"projectedPnl_60_90"`
	OpenInterestBySentiment                            hegicSentimentData  `json:"openInterestBySentiment"`
	ExpectedGrossMarginDates                           hegicDatesData      `json:"expectedGrossMarginByDateAndSentiment"`
	ExpectedGrossMargin                                []hegicScenarioData `json:"ExpectedGrossMarginForCurrentEpoch"`
	ExpectedGrossMarginByPrice4                        []hegicScenarioData `json:"ExpectedGrossMarginByPrice_4"`
	ExpectedGrossMarginByPrice3                        []hegicScenarioData `json:"ExpectedGrossMarginByPrice_3"`
	ExpectedGrossMarginByPrice2                        []hegicScenarioData `json:"ExpectedGrossMarginByPrice_2"`
	ExpectedGrossMarginByPrice1                        []hegicScenarioData `json:"ExpectedGrossMarginByPrice_1"`
	ExpectedGrossMarginByPrice                         []hegicScenarioData `json:"expectedGrossMarginByPrice"`
	ExpectedGrossMarginByPricesWithoutInversionProduct []hegicScenarioData `json:"expectedGrossMarginByPricesWithoutInversionProduct"`
	CountContracts                                     int                 `json:"countContracts"`
	RealizedRevenue                                    float64             `json:"realizedRevenue"`
	UnrealizedRevenue                                  float64             `json:"unrealizedRevenue"`
	LockedCollateral                                   float64             `json:"lockedCollateral"`
	RealizedPayOff                                     float64             `json:"realizedPayOff"`
	UnrealizedPayOff                                   float64             `json:"unrealizedPayOff"`
	RealizedGrossMargin                                float64             `json:"realizedGrossMargin"`
	UnrealizedGrossMargin                              float64             `json:"unrealizedGrossMargin"`
	TotalTradingVolume                                 float64             `json:"totalTradingVolume"`
	TotalOpenInterest                                  float64             `json:"totalOpenInterest"`
	EthOpenInterest                                    float64             `json:"ethOpenInterest"`
	BtcOpenInterest                                    float64             `json:"btcOpenInterest"`
	UnrealizedGrossMarginForCurrentEpoch               float64             `json:"unrealizedGrossMarginForCurentEpoch"`
	RealizedGrossMarginForCurrentEpoch                 float64             `json:"realizedGrossMarginForCurrentEpoch"`
	EstimatedGrossMarginByEndOfCurrentEpoch            float64             `json:"estimatedGrossMarginByTheEndOfCurrentEpoch"`
	Solvency                                           float64             `json:"solvency"`
	OperationalBalance                                 float64             `json:"opterationalBalance"`
	PayoffPoolBalance                                  float64             `json:"payoffPoolBalance"`
	ExpectedGrossMarginByDateFormatTime                hegicDateTimeData   `json:"expectedGrossMarginByDateFormatTime"`
}

type hegicScenarioData struct {
	Scenario                 string  `json:"scenario"`
	EthPrice                 float64 `json:"ethPrice"`
	BtcPrice                 float64 `json:"btcPrice"`
	EthExpectedGrossMargin   float64 `json:"ethExpectedGrossMargin"`
	BtcExpectedGrossMargin   float64 `json:"btcExpectedGrossMargin"`
	TotalExpectedGrossMargin float64 `json:"totalExpectedGrossMargin"`
	Order                    int     `json:"order"`
	CoverPoolPnl             float64 `json:"coverPoolPnl"`
}

type hegicSentimentData struct {
	Bullish        float64 `json:"bullish"`
	Bearish        float64 `json:"bearish"`
	HighVolatility float64 `json:"highVolatility"`
	LowVolatility  float64 `json:"lowVolatility"`
}

type hegicDatesData struct {
	Dates          []hegicDateRange `json:"dates"`
	Bullish        []float64        `json:"bullish"`
	Bearish        []float64        `json:"bearish"`
	HighVolatility []float64        `json:"highVolatility"`
	LowVolatility  []float64        `json:"lowVolatility"`
}

type hegicDateRange struct {
	End   time.Time `json:"end"`
	Start time.Time `json:"start"`
}

type hegicDateTimeData struct {
	Dates          []time.Time `json:"dates"`
	Bullish        []float64   `json:"bullish"`
	Bearish        []float64   `json:"bearish"`
	HighVolatility []float64   `json:"highVolatility"`
	LowVolatility  []float64   `json:"lowVolatility"`
}

type hegicUserStats struct {
	User         string     `json:"user"`
	Overall      hegicStats `json:"overall"`
	CurrentEpoch hegicStats `json:"currentEpoch"`
}

type hegicStats struct {
	TradingVolume        float64 `json:"tradingVolume"`
	PaidPremium          float64 `json:"paidPremium"`
	PaidOut              float64 `json:"paidOut"`
	PnlUsd               float64 `json:"pnlUsd"`
	PnlPercent           float64 `json:"pnlPercent"`
	ClosedContractsCount int     `json:"closedContractsCount"`
}

type hegicPosition struct {
	State                int             `json:"state"`
	ID                   int             `json:"id"`
	User                 string          `json:"user"`
	PurchaseDate         time.Time       `json:"purchaseDate"`
	ExpirationDate       time.Time       `json:"expirationDate"`
	CloseDate            time.Time       `json:"closeDate"`
	Period               int             `json:"period"`
	Amount               float64         `json:"amount"`
	AmountUsd            float64         `json:"amountUsd"`
	Type                 string          `json:"type"`
	GeneralType          string          `json:"generalType"`
	Asset                string          `json:"asset"`
	ProfitZone           hegicProfitZone `json:"profitZone"`
	StrategyStrikes      hegicStrikes    `json:"strategyStrikes"`
	ClosedPrice          float64         `json:"closedPrice"`
	PremiumPaid          float64         `json:"premiumPaid"`
	Collateral           float64         `json:"collateral"`
	PayOff               float64         `json:"payOff"`
	GrossMargin          float64         `json:"grossMargin"`
	UserNetProfit        float64         `json:"userNetProfit"`
	UserNetProfitPercent float64         `json:"userNetProfitPercent"`
	IsInversed           bool            `json:"isInversed"`
}

type hegicTradingStats struct {
	StrategiesPurchased int     `json:"strategiesPurchased"`
	FavoriteAsset       string  `json:"favoriteAsset"`
	TradingVolume       float64 `json:"tradingVolume"`
	Winrate             float64 `json:"winrate"`
	PnlUSD              float64 `json:"pnlUSD"`
	PnlPercent          float64 `json:"pnlPercent"`
}

type hegicProfitZone struct {
	Left   hegicProfitValue  `json:"left"`
	Center *hegicProfitValue `json:"center"`
	Right  hegicProfitValue  `json:"right"`
}

type hegicProfitValue struct {
	Profit bool    `json:"profit"`
	Value  float64 `json:"value"`
}

type hegicStrikes struct {
	BuyCallStrike  float64 `json:"buyCallStrike"`
	BuyPutStrike   float64 `json:"buyPutStrike"`
	SellCallStrike float64 `json:"sellCallStrike"`
	SellPutStrike  float64 `json:"sellPutStrike"`
}

type hegicPositions struct {
	Active []hegicPosition `json:"active"`
	Closed []hegicPosition `json:"closed"`
}

type hegicReferrerStats struct {
	Referrals int     `json:"referrals"`
	Earnings  float64 `json:"earnings"`
}

type hegicStaking struct {
	UserStake  float64 `json:"userStake"`
	TotalStake float64 `json:"totalStake"`
	UserShare  float64 `json:"userShare"`
}

type hegicStatsForActivePositions struct {
	ActiveStrategiesCount      int     `json:"activeStrategiesCount"`
	ActiveStrategiesUsdPnL     float64 `json:"activeStrategiesUsdPnL"`
	ActiveStrategiesPercentPnL float64 `json:"activeStrategiesPercentPnL"`
}

type hegicStatsForClosedPositions struct {
	ClosedStrategiesCount      int     `json:"closedStrategiesCount"`
	ClosedStrategiesUsdPnL     float64 `json:"closedStrategiesUsdPnL"`
	ClosedStrategiesPercentPnL float64 `json:"closedStrategiesPercentPnL"`
}

type hegicPositionsResponse struct {
	Positions []hegicPosition `json:"positions"`
}

type userData struct {
	TradingStats            hegicTradingStats            `json:"tradingStats"`
	Positions               hegicPositions               `json:"positions"`
	ReferrerStats           hegicReferrerStats           `json:"referrerStats"`
	LastActivity            time.Time                    `json:"lastActivity"`
	Staking                 hegicStaking                 `json:"staking"`
	StatsForActivePositions hegicStatsForActivePositions `json:"statsForActivePositions"`
	StatsForClosedPositions hegicStatsForClosedPositions `json:"statsForClosedPositions"`
}

type callbackPayload struct {
	UserId string
	Option models.Option
}
