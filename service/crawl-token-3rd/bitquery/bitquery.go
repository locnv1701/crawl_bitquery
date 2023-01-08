package bitquery

import (
	"crawl-token-3rd-service/pkg/log"
	"crawl-token-3rd-service/pkg/server"
	"crawl-token-3rd-service/pkg/utils"
	"crawl-token-3rd-service/service/crawl-token-3rd/model/dao"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func Call() {

	url := "https://graphql.bitquery.io"
	method := "POST"

	chainName := server.Config.GetString("CHAIN_NAME")
	chainId := server.Config.GetString("CHAIN_ID")

	fmt.Println(chainName, chainId)

	// start := time.Date(2021, 10, 28, 18, 0, 0, 0, time.Local)

	listFail := ListFailBscBitquery{}
	err := listFail.GetList()
	if err != nil {
		log.Println(log.LogLevelError, `listFail.GetList()`, err.Error())
	}
	fmt.Println(len(listFail.List))
	//2021-10-28T18:00:00
	for _, failTime := range listFail.List {
		start, error := time.Parse("2006-01-02T15:04:05", failTime.Start)
		// fmt.Println(failTime.Start, "===========================>", failTime.End)
		for {

			if error != nil {
				fmt.Println(error)
				return
			}
			bodyReq := `{"query":"{\n  ethereum(network: bsc) {\n    smartContractCalls(\n      smartContractMethod: {is: \"Contract Creation\"}\n      smartContractType: {is: Token}\n      time: {after: \"`

			timeStart := start.Format("2006-01-02T15:04:05")

			// fmt.Println(timeStart)

			end := start.Add(10 * time.Minute)

			timeEnd := end.Format("2006-01-02T15:04:05")
			// fmt.Println(timeStart, "--->", timeEnd)

			start = end

			if start.Format("2006-01-02T15:04:05") > failTime.End {
				break
			}

			// fmt.Println(timeEnd)

			bodyReq += timeStart

			bodyReq += `\", before:\"`

			bodyReq += timeEnd

			bodyReq += `\" } ) {\n      block {\n        height\n        timestamp {\n          time\n        }\n      }\n      smartContract {\n        contractType\n        address {\n          address\n        }\n        currency {\n          name\n          symbol\n          decimals\n          tokenType\n        }\n      }\n    }\n  }\n}","variables":"{}"}`

			// fmt.Println(bodyReq)

			payload := strings.NewReader(bodyReq)
			time.Sleep(1000 * time.Millisecond)

			client := &http.Client{}
			req, err := http.NewRequest(method, url, payload)
			if err != nil {
				log.Println(log.LogLevelError, `http.NewRequest(method, url, payload)`, err.Error())
				log.Println(log.LogLevelError, fmt.Sprintf("Fail start: %s end: %s ", timeStart, timeEnd), "")
				failBscBitquery := FailBscBitquery{
					Start:       timeStart,
					End:         timeEnd,
					Err:         err.Error(),
					CreatedDate: utils.Timestamp(),
				}

				err := failBscBitquery.Insert()
				if err != nil {
					log.Println(log.LogLevelError, `failBscBitquery.Insert()`, err.Error())
				}

				start = end
				continue
			}
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("X-API-KEY", "BQYnlEzOfDXVqXGjzGA865GqPTxp57Ec")

			res, err := client.Do(req)
			if err != nil {
				log.Println(log.LogLevelError, `client.Do(req)`, err.Error())
				log.Println(log.LogLevelError, fmt.Sprintf("Fail start: %s end: %s ", timeStart, timeEnd), "")
				failBscBitquery := FailBscBitquery{
					Start:       timeStart,
					End:         timeEnd,
					Err:         err.Error(),
					CreatedDate: utils.Timestamp(),
				}

				err := failBscBitquery.Insert()
				if err != nil {
					log.Println(log.LogLevelError, `failBscBitquery.Insert()`, err.Error())
				}
				start = end
				continue
			}
			defer res.Body.Close()

			response := Response{}
			body, err := io.ReadAll(res.Body)
			if err != nil {
				log.Println(log.LogLevelError, `io.ReadAll(res.Body)`, err.Error())
				log.Println(log.LogLevelError, fmt.Sprintf("Fail start: %s end: %s ", timeStart, timeEnd), "")
				failBscBitquery := FailBscBitquery{
					Start:       timeStart,
					End:         timeEnd,
					Err:         err.Error(),
					CreatedDate: utils.Timestamp(),
				}

				err := failBscBitquery.Insert()
				if err != nil {
					log.Println(log.LogLevelError, `failBscBitquery.Insert()`, err.Error())
				}
				start = end
				continue
			}

			// fmt.Println(string(body))

			err = json.Unmarshal(body, &response)
			if err != nil {

				log.Println(log.LogLevelError, `json.Unmarshal(body, &response)`, err.Error())
				log.Println(log.LogLevelError, fmt.Sprintf("Fail start: %s end: %s ", timeStart, timeEnd), "")
				failBscBitquery := FailBscBitquery{
					Start:       timeStart,
					End:         timeEnd,
					Err:         err.Error(),
					CreatedDate: utils.Timestamp(),
				}

				err := failBscBitquery.Insert()
				if err != nil {
					log.Println(log.LogLevelError, `failBscBitquery.Insert()`, err.Error())
				}

				start = end
				continue
			}

			// fmt.Println(len(response.Data.Ethereum.SmartContractCalls))
			// if len(response.Data.Ethereum.SmartContractCalls) == 0 {
			// 	continue
			// }

			count := 0

			for _, ele := range response.Data.Ethereum.SmartContractCalls {
				if ele.SmartContract.Address.Address == "0x0000000000000000000000000000000000000000" {
					continue
				}
				if ele.SmartContract.Currency.TokenType != "ERC20" {
					continue
				}
				if ele.SmartContract.Currency.Name == "" {
					continue
				}
				if ele.SmartContract.Currency.Symbol == "" {
					continue
				}
				token := dao.Crypto{
					CryptoId:    "gear5_token_" + chainName + "_" + strings.ToLower(ele.SmartContract.Address.Address),
					CryptoSrc:   "bitquery",
					CryptoCode:  GenCryptoCodeByName(ele.SmartContract.Currency.Name),
					Name:        ele.SmartContract.Currency.Name,
					Symbol:      ele.SmartContract.Currency.Symbol,
					Decimal:     ele.SmartContract.Currency.Decimals,
					Address:     strings.ToLower(ele.SmartContract.Address.Address),
					ChainId:     chainId,
					ChainName:   chainName,
					Category:    "Crypto Projects",
					SubCategory: GenSubcategoryByChainname(chainName),
					SourceUrl:   "https://graphql.bitquery.io",

					TotalReviews:      "0",
					TotalIsScam:       "0",
					TotalNotScam:      "0",
					IsScam:            false,
					IsVerifiedByAdmin: false,
					IsWarning:         false,
					IsProxy:           false,
				}
				err := token.Insert()
				if err != nil {
					log.Println(log.LogLevelError, `Insert fail token.Insert()`+token.Address, err.Error())
				}
				count += 1
			}

			// fmt.Println("timeDif: ", time.Since(start))

			log.Println(log.LogLevelInfo, fmt.Sprintf("Succes %d start: %s end: %s ", count, timeStart, timeEnd), time.Since(start))

			start = end

			if start.After(time.Now()) {
				break
			}
		}
		// log.Println(log.LogLevelInfo, fmt.Sprintf("Succes binance %s", start.Format("2006-01-02T15:04:05")), time.Since(start))
	}

	// log.Println(log.LogLevelInfo, fmt.Sprintf("Succes binance %s", start.Format("2006-01-02T15:04:05")), time.Since(start))
}

func GenCryptoCodeByName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	name = strings.Join(strings.Split(strings.Join(strings.Fields(name), " "), " "), "-")

	return name
}

func GenSubcategoryByChainname(chainname string) string {
	chainname = strings.TrimSpace(chainname)

	subcategory := strings.ToUpper(chainname[:1]) + strings.ToLower(chainname[1:]) + " " + "Ecosystem"

	return subcategory
}
