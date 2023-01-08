package solana

import (
	"crawl-token-3rd-service/pkg/log"
	"crawl-token-3rd-service/service/crawl-token-3rd/bitquery"
	"crawl-token-3rd-service/service/crawl-token-3rd/model/dao"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Call() {

	url := "https://api.solscan.io/tokens?offset=%d&limit=50"

	offset := 0
	for {

		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf(url, offset), nil)

		if err != nil {
			log.Println(log.LogLevelError, `Insert fail http.NewRequest`, err.Error())

			return
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("X-API-KEY", "BQYc4ulTHt7vDgCMKWAJNgAVn1izHiNA")

		res, err := client.Do(req)
		if err != nil {
			log.Println(log.LogLevelError, `Insert fail client.Do`, err.Error())

			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(log.LogLevelError, `Insert fail io.ReadAll`, err.Error())

			return
		}
		// fmt.Println(string(body))

		response := TokenSolanaResponse{}

		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Println(log.LogLevelError, `Insert fail json.Unmarshal`, err.Error())

		}

		// fmt.Println(len(response.Data.Tokens))

		for _, ele := range response.Data.Tokens {
			token := dao.Crypto{
				CryptoId:    "gear5_token_" + "solana" + "_" + ele.Address,
				CryptoSrc:   "solscan",
				CryptoCode:  bitquery.GenCryptoCodeByName(ele.TokenName),
				Name:        ele.TokenName,
				Symbol:      ele.TokenSymbol,
				Decimal:     ele.Decimals,
				Address:     ele.Address,
				ChainName:   "solana",
				Category:    "Crypto Projects",
				SubCategory: bitquery.GenSubcategoryByChainname("solana"),
				SourceUrl:   fmt.Sprintf(url, offset),
				SmallLogo:   ele.Icon,
				Description: ele.Extensions.Description,

				IsShow:            true,
				TotalReviews:      "0",
				TotalIsScam:       "0",
				TotalNotScam:      "0",
				IsScam:            false,
				IsVerifiedByAdmin: false,
				IsWarning:         false,
				IsProxy:           false,
			}

			mapSocial := make(map[string]any)
			if ele.Extensions.CoingeckoID != "" {
				mapSocial["coingeckoId"] = ele.Extensions.CoingeckoID
			}
			if ele.Extensions.Instagram != "" {
				mapSocial["instagram"] = ele.Extensions.Instagram
			}
			if ele.Extensions.Discord != "" {
				mapSocial["discord"] = ele.Extensions.Discord
			}
			if ele.Extensions.Github != "" {
				mapSocial["github"] = ele.Extensions.Github
			}
			if ele.Extensions.Medium != "" {
				mapSocial["medium"] = ele.Extensions.Medium
			}
			if ele.Extensions.Telegram != "" {
				mapSocial["telegram"] = ele.Extensions.Telegram
			}
			if ele.Extensions.Twitter != "" {
				mapSocial["twitter"] = ele.Extensions.Twitter
			}
			if ele.Extensions.Blog != "" {
				mapSocial["blog"] = ele.Extensions.Blog
			}
			if ele.Extensions.Website != "" {
				mapSocial["website"] = ele.Extensions.Website
			}
			if ele.Extensions.Coinmarketcap != "" {
				mapSocial["coinmarketcap"] = ele.Extensions.Coinmarketcap
			}

			token.Socials = mapSocial

			err := token.InsertCrypto()
			if err != nil {
				log.Println(log.LogLevelError, `Insert fail token.InsertCrypto()`+token.Address, err.Error())
				continue
			}

			// fmt.Println(offset, ele.TokenName)

			cryptoRank := dao.CryptoRank{
				CryptoUUID: token.Id,
				CryptoId:   "gear5_token_" + "solana" + "_" + ele.Address,
				PriceUSD:   ele.PriceUst,
				Holders:    fmt.Sprintf("%d", ele.Holder),
			}

			if ele.CoingeckoInfo != nil {
				cryptoRank.IsCoingecko = true
				cryptoRank.MarketCapUSD = ele.CoingeckoInfo.MarketData.MarketCap
				cryptoRank.TotalSupply = fmt.Sprintf("%f", ele.CoingeckoInfo.MarketData.TotalSupply)
			}

			err = cryptoRank.Insert()
			if err != nil {
				log.Println(log.LogLevelError, `Insert fail cryptoRank.Insert()`+token.Address, err.Error())
			}

		}
		offset += 50
		if offset >= 13100 {
			break
		}
	}

}
