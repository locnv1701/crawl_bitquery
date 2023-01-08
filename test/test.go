package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	url := "https://graphql.bitquery.io"
	method := "POST"

	bodyReq := `{"query":"{\n  ethereum(network: bsc) {\n    smartContractCalls(\n      smartContractMethod: {is: \"Contract Creation\"}\n      smartContractType: {is: Token}\n      time: {after: \"`

	bodyReq += "2020-08-30T00:00:00"

	bodyReq += `\", before:\"`

	bodyReq += "2020-09-01T00:00:00"

	bodyReq += `\" } ) {\n      block {\n        height\n        timestamp {\n          time\n        }\n      }\n      smartContract {\n        contractType\n        address {\n          address\n        }\n        currency {\n          name\n          symbol\n          decimals\n          tokenType\n        }\n      }\n    }\n  }\n}","variables":"{}"}`

	fmt.Println(bodyReq)

	// fmt.Sprintf(`{"query":"{\n  ethereum(network: bsc) {\n    smartContractCalls(\n      smartContractMethod: {is: \"Contract Creation\"}\n      smartContractType: {is: Token}\n      time: {after: %s, before: %s    ) {\n      block {\n        height\n        timestamp {\n          time\n        }\n      }\n      smartContract {\n        contractType\n        address {\n          address\n        }\n        currency {\n          name\n          symbol\n          decimals\n          tokenType\n        }\n      }\n    }\n  }\n}","variables":"{}"}`, , "2020-09-01T00:00:00")

	payload := strings.NewReader(bodyReq)

	// return

	// payload := strings.NewReader(`{"query":"{\n  ethereum(network: bsc) {\n    smartContractCalls(\n      smartContractMethod: {is: \"Contract Creation\"}\n      smartContractType: {is: Token}\n      time: {after: \"2020-08-30T00:00:00\", before: \"2020-12-30T00:00:00\"}\n    ) {\n      block {\n        height\n        timestamp {\n          time\n        }\n      }\n      smartContract {\n        contractType\n        address {\n          address\n        }\n        currency {\n          name\n          symbol\n          decimals\n          tokenType\n        }\n      }\n    }\n  }\n}","variables":"{}"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	fmt.Println(strings.Index(string(body), "class"))
}
