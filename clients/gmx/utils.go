package gmx

import (
	"context"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/defiants-co/perpstream-go-2/clients/common/models"
	abis "github.com/defiants-co/perpstream-go-2/clients/gmx/abi"
	"github.com/hasura/go-graphql-client"
)

// TODO: fix to get more accurate numbers
func gmxToFuturesPosition(position abis.PositionProps, priceCache *priceCache) *models.Future {
	market := getMarket(position.Addresses.Market.Hex())
	token := getToken(market)
	collateralToken := getCollateralToken(position.Addresses.CollateralToken.Hex())
	oldSizeInUsd := getOldSizeInUsd(position.Numbers.SizeInUsd)
	sizeInTokens := getSizeInTokens(position.Numbers.SizeInTokens, token)
	entryPrice := getEntryPrice(oldSizeInUsd, sizeInTokens)
	collateralTokenAmount := getCollateralTokenAmount(position.Numbers.CollateralAmount, collateralToken)
	collateralUsd := getCollateralUsd(collateralToken, collateralTokenAmount, priceCache)
	tokenPrice := priceCache.GetPrice(token)
	leverage := getLeverage(tokenPrice, sizeInTokens, collateralToken, collateralTokenAmount, priceCache)
	currentSizeInUsd := tokenPrice * sizeInTokens
	pnl := getPnl(position.Flags.IsLong, tokenPrice, entryPrice, sizeInTokens)

	val := models.NewFuture(
		market,
		collateralToken,
		position.Flags.IsLong,
		collateralTokenAmount,
		collateralUsd,
		sizeInTokens,
		currentSizeInUsd,
		leverage,
		entryPrice,
		tokenPrice,
		pnl,
		math.Round(1000*(pnl/collateralUsd))/1000,
		time.Now(),
	)

	return &val
}

func getMarket(marketHex string) string {
	market := gmxAddressToMarket[marketHex]
	if market == "" {
		market = "ATOM-USD"
	}
	return market
}

func getToken(market string) string {
	return strings.Split(market, "-")[0]
}

func getCollateralToken(collateralTokenHex string) string {
	return gmxAddressToCollateralToken[collateralTokenHex]
}

func getOldSizeInUsd(sizeInUsd *big.Int) float64 {
	oldSizeInUsd, _ := sizeInUsd.Float64()
	return math.Round(oldSizeInUsd/math.Pow(10, 26)) / math.Pow(10, 4)
}

func getSizeInTokens(sizeInTokens *big.Int, token string) float64 {
	sizeInTokensFl, _ := sizeInTokens.Float64()
	return math.Round(10000*sizeInTokensFl/math.Pow(10, float64(gmxMarketToDecimals[token]))) / 10000
}

func getEntryPrice(oldSizeInUsd, sizeInTokens float64) float64 {
	return math.Round((oldSizeInUsd/sizeInTokens)*10000) / 10000
}

func getCollateralTokenAmount(collateralAmount *big.Int, collateralToken string) float64 {
	collateralTokenAmountFl, _ := collateralAmount.Float64()
	return collateralTokenAmountFl / math.Pow(10, float64(gmxCollateralTokenDecimals[collateralToken]))
}

func getCollateralUsd(collateralToken string, collateralTokenAmount float64, priceCache *priceCache) float64 {
	collateralPrice := priceCache.GetPrice(collateralToken)
	return math.Round(collateralPrice*collateralTokenAmount*10000) / 10000
}

func getLeverage(tokenPrice, sizeInTokens float64, collateralToken string, collateralTokenAmount float64, priceCache *priceCache) float64 {
	collateralPrice := priceCache.GetPrice(collateralToken)
	return math.Round(1000*(tokenPrice*sizeInTokens)/(collateralPrice*collateralTokenAmount)) / 1000
}

func getPnl(isLong bool, tokenPrice, entryPrice, sizeInTokens float64) float64 {
	if isLong {
		return math.Round((tokenPrice-entryPrice)*sizeInTokens*100) / 100
	}
	return math.Round((entryPrice-tokenPrice)*sizeInTokens*100) / 100
}

type periodAccountStatObject struct {
	ID                         string `graphql:"id"`
	ClosedCount                int    `graphql:"closedCount"`
	CumsumCollateral           string `graphql:"cumsumCollateral"`
	CumsumSize                 string `graphql:"cumsumSize"`
	Losses                     int    `graphql:"losses"`
	MaxCapital                 string `graphql:"maxCapital"`
	RealizedPriceImpact        string `graphql:"realizedPriceImpact"`
	SumMaxSize                 string `graphql:"sumMaxSize"`
	NetCapital                 string `graphql:"netCapital"`
	RealizedFees               string `graphql:"realizedFees"`
	RealizedPnl                string `graphql:"realizedPnl"`
	Volume                     string `graphql:"volume"`
	Wins                       int    `graphql:"wins"`
	StartUnrealizedPnl         string `graphql:"startUnrealizedPnl"`
	StartUnrealizedFees        string `graphql:"startUnrealizedFees"`
	StartUnrealizedPriceImpact string `graphql:"startUnrealizedPriceImpact"`
	Typename                   string `graphql:"__typename"`
}

type query struct {
	All []periodAccountStatObject `graphql:"periodAccountStats(limit: $limit, where: {maxCapital_gte: $maxCapital, from: $from})"`
}

func getLeaderboardReq(client *graphql.Client) (*[]periodAccountStatObject, error) {
	var query query
	variables := map[string]interface{}{
		"limit":      graphql.Float(100000),
		"maxCapital": graphql.String("500000000000000000000000000000000"),
		"from":       graphql.Int(0),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}

	return &query.All, nil
}

func periodStatsToUser(u *periodAccountStatObject) models.User {
	realizedPnlFloat, _ := strconv.ParseFloat(u.NetCapital, 64)
	realizedPnl := -1 * math.Round(realizedPnlFloat/math.Pow(10, 30))

	capital, _ := strconv.ParseFloat(u.MaxCapital, 64)
	capital = capital / math.Pow(10, 30)

	pnlPercent := realizedPnl / capital

	totalTrades := u.Wins + u.Losses
	winRate := math.Round(float64(u.Wins)*100/float64(totalTrades)) / 100

	vol, _ := strconv.ParseFloat(u.Volume, 64)
	totalVolume := math.Round(vol/math.Pow(10, 28)) / 100

	return models.User{
		Id:          u.ID,
		PnlUsd:      realizedPnl,
		PnlPercent:  pnlPercent,
		WinRate:     float64(winRate),
		TotalVolume: totalVolume,
	}
}
