package hegic

import (
	"strings"
	"time"

	"github.com/defiants-co/perpstream-go-2/clients/common/models"
	"github.com/defiants-co/perpstream-go-2/clients/common/utils"
)

func getOptions() (active map[string][]models.Option, inactive map[string][]models.Option, reqErr error) {
	var optionsResponse hegicPositionsResponse

	err := utils.Fetch[hegicPositionsResponse]("positions", apiUrl, nil, &optionsResponse)
	if err != nil {
		return nil, nil, err
	}

	active = make(map[string][]models.Option)
	inactive = make(map[string][]models.Option)

	now := time.Now()

	for _, p := range optionsResponse.Positions {
		option := resolveOption(p)
		if now.After(p.ExpirationDate) || p.CloseDate.Year() != 1 {
			inactive[p.User] = append(inactive[p.User], option)
		} else {
			active[p.User] = append(active[p.User], option)
		}
	}

	for key := range inactive {
		_, exists := active[key]
		if !exists {
			active[key] = make([]models.Option, 0)
		}
	}

	return active, inactive, nil
}

func resolveOption(hegicOption hegicPosition) models.Option {
	strikes := make([]models.Strike, 0)

	if hegicOption.StrategyStrikes.BuyCallStrike > 0 {
		strikes = append(strikes, models.Strike{
			Type:        "call",
			IsBought:    true,
			StrikePrice: utils.RoundTo2Decimals(hegicOption.StrategyStrikes.BuyCallStrike),
		})
	}
	if hegicOption.StrategyStrikes.BuyPutStrike > 0 {
		strikes = append(strikes, models.Strike{
			Type:        "put",
			IsBought:    true,
			StrikePrice: utils.RoundTo2Decimals(hegicOption.StrategyStrikes.BuyPutStrike),
		})
	}
	if hegicOption.StrategyStrikes.SellCallStrike > 0 {
		strikes = append(strikes, models.Strike{
			Type:        "call",
			IsBought:    false,
			StrikePrice: utils.RoundTo2Decimals(hegicOption.StrategyStrikes.SellCallStrike),
		})
	}
	if hegicOption.StrategyStrikes.SellPutStrike > 0 {
		strikes = append(strikes, models.Strike{
			Type:        "put",
			IsBought:    false,
			StrikePrice: utils.RoundTo2Decimals(hegicOption.StrategyStrikes.SellPutStrike),
		})
	}

	return models.NewOption(
		strings.ToLower(hegicOption.Type),
		hegicOption.Asset,
		hegicOption.IsInversed,
		hegicOption.Amount,
		utils.RoundTo2Decimals(hegicOption.AmountUsd),
		utils.RoundTo2Decimals(hegicOption.PremiumPaid),
		utils.RoundTo2Decimals(hegicOption.UserNetProfit),
		utils.RoundToNDecimals(hegicOption.UserNetProfitPercent, 4),
		hegicOption.PurchaseDate,
		hegicOption.ExpirationDate,
		hegicOption.Period,
		strikes,
	)
}

func resolveUser(userStats hegicUserStats, inactiveOptionsMap *utils.ConcurrentMap[[]models.Option]) models.User {
	total := 0
	wins := 0

	usersOptions, exists := inactiveOptionsMap.Get(userStats.User)
	if exists {
		for _, option := range usersOptions {
			if option.PnlUsd > 0 {
				wins++
			}
			total++
		}
	}

	winRate := 0.0
	if total > 0 {
		winRate = float64(wins) / float64(total)
	}

	return models.NewUser(
		userStats.User,
		userStats.Overall.PnlUsd,
		userStats.Overall.PnlPercent,
		winRate,
		userStats.Overall.TradingVolume,
	)
}

func optionToCallbackPayload(option models.Option, userId string) callbackPayload {
	return callbackPayload{
		Option: option,
		UserId: userId,
	}
}

func getDifference(originalOptionsMap map[string][]models.Option, newOptionsMap map[string][]models.Option) []callbackPayload {
	var difference []callbackPayload

	for userId, newOptions := range newOptionsMap {
		originalOptions, exists := originalOptionsMap[userId]
		if !exists {
			for _, x := range newOptions {
				difference = append(difference, optionToCallbackPayload(x, userId))
			}
			continue
		}
		for _, newOption := range newOptionsMap[userId] {
			found := false
			for _, oldOption := range originalOptions {
				if newOption.Equal(&oldOption) {
					found = true
				}
			}

			if !found {
				difference = append(difference, callbackPayload{UserId: userId, Option: newOption})
			}
		}

	}

	return difference
}
