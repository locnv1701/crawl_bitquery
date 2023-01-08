package dto

type CoinstatHref struct {
	Href   string `json:"href"`
	Img    string `json:"img"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type ListCoinstat struct {
	List []*CoinstatHref `json:"list"`
}
