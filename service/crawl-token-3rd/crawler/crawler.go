package crawler

import (
	"crawl-token-3rd-service/pkg/utils"
	"crawl-token-3rd-service/service/crawl-token-3rd/constant"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Crawl() {
	url := fmt.Sprintf(`https://etherscan.io/token/%s`, "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48") //todo
	// url := fmt.Sprintf(`https://snowtrace.io/token/%s#code`, "0xc7198437980c041c805a1edcba50c1ce5db95118") //todo
	// url := fmt.Sprintf(`https://moonscan.io/token/%s#code`, "0xefaeee334f0fd1712f9a8cc375f427d9cdd40d73") //todo
	// url := fmt.Sprintf(`https://optimistic.etherscan.io/token/%s#code`, "0x2e3d870790dc77a83dd1d18184acc7439a53f475") //todo
	// url := fmt.Sprintf(`https://cronoscan.com/token/%s#code`, "0xbc6f24649ccd67ec42342accdceccb2efa27c9d9") //todo
	fmt.Println(url)

	url = "https://ethplorer.io/vi/address/0xb8c77482e45f1f44de1745f52c74426c631bdd52#chart=candlestick"

	dom := utils.GetHtmlDomJsRenderByUrl(url)

	fmt.Println(dom)

	//reponse not equal 200(404 --> No data to crawl)
	if dom == nil {
		fmt.Println("toreto")
		return
	}

	tokenInfoExplorer := extractTokenInfoExplorerHrefByHtmlDom(dom, "1", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")

	fmt.Println(tokenInfoExplorer)

	fmt.Println("jnneie")
}

type TokenInfoExplorer struct {
	Address       string            `json:"address"` //lowcase
	ChainId       string            `json:"chainId"`
	Holders       *big.Int          `json:"holders"`
	Transfers     *big.Int          `json:"transfers"`
	OfficialSite  string            `json:"officialSite"`
	SocialProfile map[string]string `json:"socialProfile"`
	ContractABI   string            `json:"contractABI"`
	Source        string            `json:"src"`
	CreatedDate   string            `json:"createddate"`
	UpdatedDate   string            `json:"updateddate"`
}

// func extractCoinInfoByHtmlDom(dom *goquery.Document) []*dto.CoinstatHref {

// 	listCoinstatHref := make([]*dto.CoinstatHref, 0)

// 	fmt.Println("run here")

// 	fmt.Println(dom)

// 	coin := Coinstat{}
// 	// explorers := []string{}

// 	domKey := `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-3021552425 guide-body coin-page`)
// 	dom.Find(domKey).Each(func(i int, s *goquery.Selection) {

// 		domKey := `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-1466193091 item-previewer`)
// 		s.Find(domKey).Each(func(i int, s *goquery.Selection) {

// 			domKey := `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-1466193091 top-section`)
// 			s.Find(domKey).Each(func(i int, s *goquery.Selection) {

// 				domKey := `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-1466193091 item-previewer-section`)
// 				s.Find(domKey).Each(func(i int, s *goquery.Selection) {

// 					domKey := `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-1466193091 name-wrapper`)
// 					s.Find(domKey).Each(func(i int, s *goquery.Selection) {

// 						domKey = `img`
// 						s.Find(domKey).Each(func(i int, s *goquery.Selection) {
// 							attrKey := `src`
// 							srcImg, found := s.Attr(attrKey)
// 							if found {
// 								coin.Image = srcImg
// 							}
// 						})

// 						domKey = `h1`
// 						s.Find(domKey).Each(func(i int, s *goquery.Selection) {
// 							coin.Symbol = strings.Trim(s.Children().Text(), " ")
// 							coin.Symbol = coin.Symbol[1 : len(coin.Symbol)-1]
// 							coin.Name = strings.Trim(s.Children().Remove().End().Text(), " ")

// 							// coin.Symbol = strings.Split(strings.Split(s.Children().Text(), "(")[0], ")")[0]
// 							fmt.Printf("=%s%s=", coin.Name, coin.Symbol)
// 						})

// 					})
// 				})
// 			})

// 			domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-1466193091 overview-details`)
// 			s.Find(domKey).Each(func(i int, s *goquery.Selection) {
// 				fmt.Println(`jsx-1466193091 overview-details`)

// 				domKey := `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-3896788395 social-column`)
// 				s.Find(domKey).Each(func(i int, s *goquery.Selection) {

// 					fmt.Println("jsx-3896788395 social-column")

// 					domKey := `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-3896788395 social-column-row`)
// 					s.Find(domKey).Each(func(i int, s *goquery.Selection) {
// 						fmt.Println(s.Children().Text())
// 						fmt.Println("jsx-3896788395 social-column-row")

// 						title := "Social Column"
// 						coin.Detail[title] = map[string]string{}
// 						domKey := `span` + utils.ConvertClassesFormatFromBrowserToGoQuery(`table-column-title`)
// 						s.Find(domKey).Each(func(i int, s *goquery.Selection) {
// 							fmt.Println(s.Text())
// 							title = s.Text()
// 						})

// 						domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-3896788395 list`)
// 						s.Find(domKey).Each(func(i int, s *goquery.Selection) {
// 							fmt.Println("jsx-3896788395 list")

// 							socialMap := make(map[string]string)

// 							domKey := `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-3896788395 social-item active`)
// 							s.Find(domKey).Each(func(i int, s *goquery.Selection) {
// 								fmt.Println("jsx-3896788395 social-item active")

// 								domKey = `a`
// 								s.Find(domKey).Each(func(i int, s *goquery.Selection) {

// 									attrKey := `href`
// 									href, found := s.Attr(attrKey)
// 									if found {
// 										socialMap["a"] = href
// 									}
// 									coin.Detail[title] = s

// 								})
// 							})
// 						})
// 					})
// 					// domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-3896788395 list`)
// 					// s.Find(domKey).Each(func(i int, s *goquery.Selection) {

// 					// 	domKey = `a`
// 					// 	s.Find(domKey).Each(func(i int, s *goquery.Selection) {
// 					// 		attrKey := `href`
// 					// 		href, found := s.Attr(attrKey)
// 					// 		if found {
// 					// 			coin.Detail["Community"] = href
// 					// 		}
// 					// 	})
// 					// })

// 					// 	domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-3896788395 social-column-row explorers`)
// 					// 	s.Find(domKey).Each(func(i int, s *goquery.Selection) {
// 					// 		domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-3896788395 list`)
// 					// 		s.Find(domKey).Each(func(i int, s *goquery.Selection) {

// 					// 			domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-3896788395 social-item active`)
// 					// 			s.Find(domKey).Each(func(i int, s *goquery.Selection) {

// 					// 				domKey = `a`
// 					// 				s.Find(domKey).Each(func(i int, s *goquery.Selection) {
// 					// 					attrKey := `href`
// 					// 					href, found := s.Attr(attrKey)
// 					// 					if found {
// 					// 						explorers = append(explorers, href)
// 					// 					}
// 					// 				})
// 					// 				coin.Detail["explorers"] = explorers
// 					// 			})
// 					// 		})
// 					// 	})

// 					// 	// domKey := `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`jsx-3896788395 list`)
// 					// 	// s.Find(domKey).Each(func(i int, s *goquery.Selection) {

// 					// 	// 	domKey = `a`
// 					// 	// 	s.Find(domKey).Each(func(i int, s *goquery.Selection) {
// 					// 	// 		attrKey := `href`
// 					// 	// 		href, found := s.Attr(attrKey)
// 					// 	// 		if found {
// 					// 	// 			coin.Detail["website"] = href
// 					// 	// 		}
// 					// 	// 	})
// 					// 	// })

// 					// })

// 				})
// 			})

// 		})

// 	})

// 	fmt.Println(coin)

// 	fmt.Println("lenen", len(listCoinstatHref))

// 	return listCoinstatHref
// }

func extractTokenInfoExplorerHrefByHtmlDom(dom *goquery.Document, chainId string, address string) TokenInfoExplorer {

	domKey := `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`address-token-txsCount`)
	dom.Find(domKey).Each(func(i int, s *goquery.Selection) {
		fmt.Println("==============run here================")
	})
	tokenInfoExplorer := TokenInfoExplorer{
		ChainId: chainId,
		Address: address,
	}

	domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`col-md-6 mb-3 mb-md-0`)
	dom.Find(domKey).Each(func(i int, s *goquery.Selection) {

		domKey := `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`card-body`)
		s.Find(domKey).Each(func(i int, s *goquery.Selection) {

			domKey = `div#ContentPlaceHolder1_tr_tokenHolders`
			s.Find(domKey).Each(func(i int, s *goquery.Selection) {

				domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`row align-items-center`)
				s.Find(domKey).Each(func(i int, s *goquery.Selection) {

					domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`col-md-8`)
					s.Find(domKey).Each(func(i int, s *goquery.Selection) {

						domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`d-flex align-items-center`)
						s.Find(domKey).Each(func(i int, s *goquery.Selection) {

							domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`mr-3`)
							s.Find(domKey).Each(func(i int, s *goquery.Selection) {
								holders := strings.Fields(s.Children().Remove().End().Text())[0]
								holders = strings.ReplaceAll(holders, ",", "")
								// holders = alphanumericRegex.ReplaceAllString(holders, "")
								if intHolders, ok := new(big.Int).SetString(holders, 10); ok {
									tokenInfoExplorer.Holders = intHolders
								}
								fmt.Println("holders", holders, "--") //todo: holders
							})
						})

					})
				})

			})
		})
	})

	domKey = `div#ContentPlaceHolder1_tr_officialsite_1`
	dom.Find(domKey).Each(func(i int, s *goquery.Selection) {

		domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`row align-items-center`)
		s.Find(domKey).Each(func(i int, s *goquery.Selection) {

			domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`col-md-8`)
			s.Find(domKey).Each(func(i int, s *goquery.Selection) {

				domKey = `a`
				s.Find(domKey).Each(func(i int, s *goquery.Selection) {
					fmt.Println("official site", s.Text()) //todo: offcail site

					tokenInfoExplorer.OfficialSite = s.Text()
				})
			})
		})
	})

	domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`row mb-4`)
	dom.Find(domKey).Each(func(i int, s *goquery.Selection) {

		domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`col-md-6`)
		s.Find(domKey).Each(func(i int, s *goquery.Selection) {

			domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`card h-100`)
			s.Find(domKey).Each(func(i int, s *goquery.Selection) {

				domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`card-body`)
				s.Find(domKey).Each(func(i int, s *goquery.Selection) {

					domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`row align-items-center`)
					s.Find(domKey).Each(func(i int, s *goquery.Selection) {

						domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`col-md-8`)
						s.Find(domKey).Each(func(i int, s *goquery.Selection) {

							domKey = `ul` + utils.ConvertClassesFormatFromBrowserToGoQuery(`list-inline mb-0`)
							s.Find(domKey).Each(func(i int, s *goquery.Selection) {

								mapSocial := make(map[string]string)
								domKey = `li` + utils.ConvertClassesFormatFromBrowserToGoQuery(`list-inline-item mr-3`)
								s.Find(domKey).Each(func(i int, s *goquery.Selection) {

									domKey = `a` + utils.ConvertClassesFormatFromBrowserToGoQuery(`link-hover-secondary`)
									s.Find(domKey).Each(func(i int, s *goquery.Selection) {
										if platform, ok := s.Attr(`data-original-title`); ok {
											if href, ok := s.Attr(`href`); ok {
												fmt.Println(platform, href)
												mapSocial[platform] = href
											}
										}
									})
								})
								fmt.Println(mapSocial)
								tokenInfoExplorer.SocialProfile = mapSocial
							})
						})
					})
				})
			})
		})
	})

	domKey = `div` + utils.ConvertClassesFormatFromBrowserToGoQuery(`d-flex justify-content-between mb-4`)
	dom.Find(domKey).Each(func(i int, s *goquery.Selection) {

		// fmt.Println("ABI", s.Text())
		fmt.Println("bravo", s.Text())
	})

	domKey = `pre` + utils.ConvertClassesFormatFromBrowserToGoQuery(`wordwrap js-copytextarea2`)
	dom.Find(domKey).Each(func(i int, s *goquery.Selection) {
		// fmt.Println("ABI", s.Text())
		tokenInfoExplorer.ContractABI = s.Text()
	})

	transfer, err := CallGetTokenTransfers(chainId, address)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(transfer)

	return tokenInfoExplorer
}

func CallGetTokenTransfers(chainId string, address string) (big.Int, error) {

	// url := GetUrlCallTransfers(chainId)

	url := "https://etherscan.io/token/generic-tokentxns2?contractAddress=0xdac17f958d2ee523a2206206994597c13d831ec7&mode=&sid=567d9322b5b08f81a26d55090301c2f0&m=normal&p=1"

	fmt.Println("url  lllllllllllllll---------", url)

	// dom := utils.GetHtmlDomJsRenderByUrl(fmt.Sprintf(url, address))

	// fmt.Println(dom.Html())

	// domKey := `div#maindiv`
	// dom.Find(domKey).Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println(s.Text())

	// 	domKey := `p` + utils.ConvertClassesFormatFromBrowserToGoQuery(`mb-2 mb-md-0`)
	// 	s.Find(domKey).Each(func(i int, s *goquery.Selection) {
	// 		fmt.Println(s.Text())
	// 	})

	// })

	// if url == "" {
	// 	return *big.NewInt(int64(0)), nil
	// }

	// res, err := http.Get(fmt.Sprintf(url, address))
	res, err := http.Get(fmt.Sprintf(url))
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println(err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	// fmt.Println(string(data))

	fmt.Println("===================================", strings.Index(string(data), "169,137,522"))

	start := strings.Index(string(data), "var totaltxns = '") + len("var totaltxns = '")

	endTransfers := start
	for index := start; index < len(string(data)); index++ {
		if string(data)[index] == '\'' {
			endTransfers = index + 1
			break
		}
	}

	fmt.Println("dmt", string(data)[start:endTransfers])

	return *big.NewInt(int64(0)), nil
}

func GetUrlCallTransfers(chainId string) string {
	linkCallTransfers := ""

	switch chainId {
	case constant.ETHEREUM_ID:
		linkCallTransfers = constant.Ethereum_Chain_Call_Transfers
	case constant.BINANCE_SMART_CHAIN_ID:
		linkCallTransfers = constant.BinanceSmartChain_Call_Transfers
	case constant.FANTOM_ID:
		linkCallTransfers = constant.Fantom_Call_Transfers
	case constant.CELO_ID:
		linkCallTransfers = constant.Celo_Call_Transfers
	case constant.POLYGON_ID:
		linkCallTransfers = constant.Polygon_Call_Transfers
	case constant.AVALANCHE_C_CHAIN_ID:
		linkCallTransfers = constant.AvalancheCChain_Call_Transfers
	case constant.OPTIMISM_ID:
		linkCallTransfers = constant.Optimism_Call_Transfers
	case constant.ARBITRUM_ID:
		linkCallTransfers = constant.Arbitrum_Call_Transfers
	case constant.MOONBEAM_ID:
		linkCallTransfers = constant.Moonbeam_Call_Transfers
	case constant.KAVA_ID:
		linkCallTransfers = constant.Kava_Call_Transfers
	case constant.CRONOS_ID:
		linkCallTransfers = constant.Cronos_Call_Transfers
	default:
		linkCallTransfers = ""
	}

	return linkCallTransfers
}

func IsEndPage(dom *goquery.Document) bool {
	isEndPage := false
	domKey := `h2`
	dom.Find(domKey).Each(func(i int, s *goquery.Selection) {
		_txtNotifyEndPageDappRadar := `Please change the filters to explore more`
		if s.Text() == _txtNotifyEndPageDappRadar {
			isEndPage = true
		}
	})
	return isEndPage
}
