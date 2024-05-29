package gmx

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"strconv"

	"github.com/defiants-co/perpstream-go-2/clients/common/utils"
)

// priceCache holds a concurrent map to store prices.
type priceCache struct {
	internalMap utils.ConcurrentMap[float64]
}

// NewPriceCache initializes a new PriceCache and updates prices.
func newPriceCache() *priceCache {
	pc := &priceCache{
		internalMap: *utils.NewConcurrentMap[float64](),
	}
	pc.UpdatePrices()
	return pc
}

// GetPrice retrieves the price of a given token from the cache.
// If the token does not exist in the cache, it returns 0.
func (p *priceCache) GetPrice(token string) float64 {
	v, exists := p.internalMap.Get(token)
	if exists {
		return v
	}

	return 0
}

// GetAllPrices returns all the prices stored in the cache.
func (p *priceCache) GetAllPrices() map[string]float64 {
	return p.internalMap.Values()
}

// UpdatePrices fetches the latest prices from an external API and updates the cache.
func (pc *priceCache) UpdatePrices() {
	resp, err := http.Get(pricesUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var apiResponse apiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return
	}

	for _, priceData := range apiResponse.SignedPrices {
		priceFloat, err := strconv.ParseFloat(priceData.MaxPriceFull, 64)
		if err != nil {
			continue
		}

		priceFactor := 6
		if !utils.Contains([]string{"USDC", "USDC.e", "USDT"}, priceData.TokenSymbol) {
			priceFactor = gmxMarketToDecimals[priceData.TokenSymbol]
		}

		price := priceFloat / math.Pow(10, float64(30-priceFactor))
		if utils.Contains(utils.GetKeys(gmxMarketToDecimals), priceData.TokenSymbol) {
			pc.internalMap.Set(priceData.TokenSymbol, math.Round(price*10000)/10000)
		}
	}
}

// SignedPrice represents the price data for a token.
type signedPrice struct {
	TokenSymbol  string `json:"tokenSymbol"`
	MaxPriceFull string `json:"maxPriceFull"`
	MinPriceFull string `json:"minPriceFull"`
}

// ApiResponse represents the structure of the API response containing price data.
type apiResponse struct {
	SignedPrices []signedPrice `json:"signedPrices"`
}
