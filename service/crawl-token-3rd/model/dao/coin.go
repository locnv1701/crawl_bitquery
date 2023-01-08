package dao

import (
	"crawl-token-3rd-service/pkg/db"
	"crawl-token-3rd-service/pkg/utils"
	"encoding/json"
)

// type Coin struct {
// 	CoinId        string         `json:"coinId"`
// 	Type          string         `json:"type"`    //coin or token
// 	Address       string         `json:"address"` //lowcase
// 	ChainId       string         `json:"chainId"`
// 	Symbol        string         `json:"symbol"` //upcase
// 	Name          string         `json:"name"`
// 	Tag           string         `json:"tag"`
// 	Decimals      uint8          `json:"decimals"`
// 	TotalSupply   string         `json:"totalSupply"`
// 	MaxSupply     string         `json:"maxSupply"`
// 	Marketcap     string         `json:"marketcap"`
// 	VolumeTrading string         `json:"volumeTrading"`
// 	Image         string         `json:"image"`
// 	Detail        map[string]any `json:"detail"`
// 	Source        string         `json:"src"`
// }

// func (coin *Coin) Count() (int, error) {
// 	count := 0
// 	query := `select count(*) from coin where src = $1 and chainId = $2;`
// 	err := db.PSQL.QueryRow(query, coin.Source, coin.ChainId).Scan(&count)

// 	return count, err
// }

// func (coin *Coin) InsertCoin() error {
// 	query := `INSERT INTO public.coin
// 	(coinId, "type", address, chainid, symbol, "name", tag, decimals, totalsupply, maxsupply, marketcap, volumetrading, image, detail, src)
// 	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);
// 	;
// 	`

// 	coinDetailJSONB, err := json.Marshal(coin.Detail)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = db.PSQL.Exec(query, coin.CoinId, coin.Type, coin.Address, coin.ChainId, coin.Symbol, coin.Name, coin.Tag, coin.Decimals, coin.TotalSupply, coin.MaxSupply, coin.Marketcap, coin.VolumeTrading, coin.Image, coinDetailJSONB, coin.Source)
// 	return err
// }

// type Detail struct {
// 	Community struct {
// 		Telegram      []string `json:"telegram"`
// 		Twitter       []string `json:"twitter"`
// 		Youtube       []string `json:"youtube"`
// 		Facebook      []string `json:"facebook"`
// 		Discord       []string `json:"discord"`
// 		Instagram     []string `json:"instagram"`
// 		Email         []string `json:"email"`
// 		Blog          []string `json:"blog"`
// 		CoinMarketCap []string `json:"coinmarketcap"`
// 		CoinGecko     []string `json:"coingecko"`
// 		// các thông tin khác lưu tương tự dưới dạng key []string
// 	} `json:"community"`
// 	Decimals struct {
// 		Chainname struct {
// 			ContractAddress string `json:"contract_address"`
// 			DecimalPlace    int    `json:"decimal_place"`
// 		} `json:"chainName"`
// 	} `json:"decimals"`
// 	Marketcap  string `json:"marketcap"`
// 	MaxSupply  string `json:"maxSupply"`
// 	SourceCode struct {
// 		Github     []string `json:"github"`
// 		Whitepaper []string `json:"whitepaper"`
// 	} `json:"sourceCode"`
// 	TotalSupply   int `json:"totalSupply"`
// 	VolumeTrading int `json:"volumeTrading"`
// 	Website       struct {
// 		OfficialSite []string `json:"official_site"`
// 	} `json:"website"`
// 	// các thông tin khác lưu dưới dạng key []string
// }

type Crypto struct {
	Id                []uint8
	CryptoId          string //gear5_token_ethereum_{address}
	CryptoSrc         string //bitquery
	CryptoCode        string //name.(lowercase + replace " " - "-")
	Name              string //!!
	Symbol            string //!!
	Decimal           int    //!!
	Address           string //lowercase
	ContractCreator   string
	Thumblogo         string
	SmallLogo         string
	BigLogo           string
	ChainId           string //!!
	ChainName         string //!!
	Category          string //!!
	SubCategory       string
	Description       string
	Rank              string
	Socials           map[string]any
	TotalReviews      string //0
	TotalIsScam       string //0
	TotalNotScam      string //0
	ScamDate          string //nil
	IsScam            bool   //false
	IsVerifiedByAdmin bool   //false
	IsShow            bool   //
	IsWarning         bool   //false
	IsProxy           bool   //false
	ObjectReference   map[string]any
	Proof             map[string]any
	SourceUrl         string //https://graphql.bitquery.io/ide#
	CreatedDate       string
	UpdatedDate       string
}

func (crypto *Crypto) Insert() error {
	query := `INSERT INTO public.crypto
	(cryptoid, cryptosrc,cryptocode, name, symbol, "decimal", address, chainid, chainname, category, subCategory, sourceurl, createddate, updateddate,
	totalReviews, totalIsScam, totalNotScam, IsScam, IsVerifiedByAdmin, IsWarning, IsProxy	)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)`

	_, err := db.PSQL.Exec(query, crypto.CryptoId, crypto.CryptoSrc, crypto.CryptoCode, crypto.Name, crypto.Symbol, crypto.Decimal, crypto.Address, crypto.ChainId, crypto.ChainName, crypto.Category, crypto.SubCategory, crypto.SourceUrl, utils.Timestamp(), utils.Timestamp(),
		crypto.TotalReviews, crypto.TotalIsScam, crypto.TotalNotScam, crypto.IsScam, crypto.IsVerifiedByAdmin, crypto.IsWarning, crypto.IsProxy)
	if err != nil {
		return err
	}
	return nil
}

func (crypto *Crypto) InsertCrypto() error {
	query := `INSERT INTO public.crypto
	(cryptoid, cryptosrc,cryptocode, name, symbol, "decimal", address, chainname, category, subCategory, sourceurl, smallLogo, description,socials, createddate, updateddate,
	isShow, totalReviews, totalIsScam, totalNotScam, IsScam, IsVerifiedByAdmin, IsWarning, IsProxy	)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24) returning id`

	socialsJSONB, err := json.Marshal(crypto.Socials)
	if err != nil {
		return err
	}

	err = db.PSQL.QueryRow(query, crypto.CryptoId, crypto.CryptoSrc, crypto.CryptoCode, crypto.Name, crypto.Symbol, crypto.Decimal, crypto.Address, crypto.ChainName, crypto.Category, crypto.SubCategory, crypto.SourceUrl, crypto.SmallLogo, crypto.Description, socialsJSONB, utils.Timestamp(), utils.Timestamp(),
		crypto.IsShow, crypto.TotalReviews, crypto.TotalIsScam, crypto.TotalNotScam, crypto.IsScam, crypto.IsVerifiedByAdmin, crypto.IsWarning, crypto.IsProxy).Scan(&crypto.Id)
	if err != nil {
		return err
	}
	return nil
}

type CryptoRank struct {
	Id           []uint8
	CryptoUUID   []uint8
	CryptoId     string //gear5_token_ethereum_{address}
	MarketCapUSD float64
	TotalSupply  string
	PriceUSD     float64
	Holders      string
	Transfers    string

	IsCoingecko     bool
	IsCoinmarketcap bool
	IsBinance       bool
	IsCoinbase      bool
	IsPancakeswap   bool
	IsUniswap       bool

	CreatedDate string
	UpdatedDate string
}

func (cryptoRank *CryptoRank) Insert() error {
	query := `INSERT INTO public.crypto_rank
	(cryptouuid, cryptoid, marketcapusd, totalsupply, priceusd, holders, iscoingecko, createddate, updateddate)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`

	_, err := db.PSQL.Exec(query, cryptoRank.CryptoUUID, cryptoRank.CryptoId, cryptoRank.MarketCapUSD, cryptoRank.TotalSupply, cryptoRank.PriceUSD, cryptoRank.Holders, cryptoRank.IsCoingecko, utils.Timestamp(), utils.Timestamp())

	return err

}
