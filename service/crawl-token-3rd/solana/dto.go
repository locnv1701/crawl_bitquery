package solana

import "time"

type TokenSolanaDTO struct {
	PriceUst      float64  `json:"priceUst"`
	Tag           []string `json:"tag"`
	TokenName     string   `json:"tokenName"`
	TokenSymbol   string   `json:"tokenSymbol"`
	Twitter       string   `json:"twitter"`
	Website       string   `json:"website"`
	CoingeckoInfo *struct {
		CoingeckoRank int `json:"coingeckoRank"`
		MarketCapRank int `json:"marketCapRank"`
		MarketData    struct {
			CurrentPrice                 float64     `json:"currentPrice"`
			Ath                          float64     `json:"ath"`
			AthChangePercentage          float64     `json:"athChangePercentage"`
			AthDate                      time.Time   `json:"athDate"`
			Atl                          float64     `json:"atl"`
			AtlChangePercentage          float64     `json:"atlChangePercentage"`
			AtlDate                      time.Time   `json:"atlDate"`
			MarketCap                    float64     `json:"marketCap"`
			MarketCapRank                int         `json:"marketCapRank"`
			FullyDilutedValuation        float64     `json:"fullyDilutedValuation"`
			TotalVolume                  float64     `json:"totalVolume"`
			PriceHigh24H                 float64     `json:"priceHigh24h"`
			PriceLow24H                  float64     `json:"priceLow24h"`
			PriceChange24H               float64     `json:"priceChange24h"`
			PriceChangePercentage24H     float64     `json:"priceChangePercentage24h"`
			PriceChangePercentage7D      float64     `json:"priceChangePercentage7d"`
			PriceChangePercentage14D     float64     `json:"priceChangePercentage14d"`
			PriceChangePercentage30D     float64     `json:"priceChangePercentage30d"`
			PriceChangePercentage60D     float64     `json:"priceChangePercentage60d"`
			PriceChangePercentage200D    float64     `json:"priceChangePercentage200d"`
			PriceChangePercentage1Y      float64     `json:"priceChangePercentage1y"`
			MarketCapChange24H           float64     `json:"marketCapChange24h"`
			MarketCapChangePercentage24H float64     `json:"marketCapChangePercentage24h"`
			TotalSupply                  float64     `json:"totalSupply"`
			MaxSupply                    interface{} `json:"maxSupply"`
			CirculatingSupply            float64     `json:"circulatingSupply"`
			LastUpdated                  time.Time   `json:"lastUpdated"`
		} `json:"marketData"`
	} `json:"coingeckoInfo"`
	SolAlphaVolume float64   `json:"solAlphaVolume"`
	ID             string    `json:"_id"`
	Address        string    `json:"address"`
	CreatedAt      time.Time `json:"createdAt"`
	Decimals       int       `json:"decimals"`
	Extensions     struct {
		Description   string `json:"description"`
		CoingeckoID   string `json:"coingeckoId"`
		Instagram     string `json:"instagram"`
		Discord       string `json:"discord"`
		Github        string `json:"github"`
		Medium        string `json:"medium"`
		Telegram      string `json:"telegram"`
		Twitter       string `json:"twitter"`
		Blog          string `json:"blog"`
		Website       string `json:"website"`
		Coinmarketcap string `json:"coinmarketcap"`
	} `json:"extensions"`
	Icon           string    `json:"icon"`
	IsViolate      bool      `json:"isViolate"`
	MarketCapRank  int       `json:"marketCapRank"`
	MintAddress    string    `json:"mintAddress"`
	SymbolHasLower bool      `json:"symbolHasLower"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Holder         int       `json:"holder"`
	MarketCapFD    float64   `json:"marketCapFD"`
}

type TokenSolanaResponse struct {
	Succcess bool `json:"succcess"`
	Data     Data `json:"data"`
}

type Data struct {
	Tokens []TokenSolanaDTO `json:"tokens"`
}
