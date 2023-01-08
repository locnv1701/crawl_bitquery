package constant

import "time"

const (
	RESP_SUCCESS_STATUS_CODE      = 200
	RESP_TOO_MANY_REQ_STATUS_CODE = 429
	RESP_NOT_FOUND_STATUS_CODE    = 404
	WAIT_DURATION_WHEN_RATE_LIMIT = 5 * time.Second

	ETHEREUM_ID            = "1"
	BINANCE_SMART_CHAIN_ID = "56"
	FANTOM_ID              = "250"
	CELO_ID                = "42220"
	POLYGON_ID             = "137"
	AVALANCHE_C_CHAIN_ID   = "43114"
	OPTIMISM_ID            = "10"
	ARBITRUM_ID            = "42161"
	MOONBEAM_ID            = "1284"
	KAVA_ID                = "2222"
	CRONOS_ID              = "25"

	Ethereum_Chain_Call_Transfers    = "https://etherscan.io/token/generic-tokentxns2?m=normal&contractAddress=%s&a=&sid=567d9322b5b08f81a26d55090301c2f0&p=1"
	BinanceSmartChain_Call_Transfers = "https://bscscan.com/token/generic-tokentxns2?m=normal&contractAddress=%s&a=&sid=507d4057ff01310fc11de79bcb2be7cf&p=1"
	Fantom_Call_Transfers            = ""
	Celo_Call_Transfers              = "https://celoscan.io/token/generic-tokentxns2?m=normal&contractAddress=%s&a=&sid=c299886bb864569910846521b639af89&p=1"
	Polygon_Call_Transfers           = "https://polygonscan.com/token/generic-tokentxns2?m=normal&contractAddress=%s&a=&sid=9cff6e2986726253c82831f3465e1e29&p=1"
	AvalancheCChain_Call_Transfers   = "https://snowtrace.io/token/generic-tokentxns2?m=normal&contractAddress=%s&a=&sid=f8d44d8d3046eb24777179c3cda5ee59&p=1"
	Optimism_Call_Transfers          = "https://optimistic.etherscan.io/token/generic-tokentxns2?m=normal&contractAddress=%s&a=&sid=8cbb09b8b8dccb86adb18927ee119fbb&p=1"
	Arbitrum_Call_Transfers          = "https://arbiscan.io/token/generic-tokentxns2?m=normal&contractAddress=%s&a=&sid=61dbb59a70de37f0bca8ff8b714cfb43&p=1"
	Moonbeam_Call_Transfers          = "https://moonscan.io/token/generic-tokentxns2?m=normal&contractAddress=%s&a=&sid=7e4ce09adc26b1846d55bee124b13378&p=1"
	Kava_Call_Transfers              = ""
	Cronos_Call_Transfers            = "https://cronoscan.com/token/generic-tokentxns2?m=normal&contractAddress=%s&a=&sid=b866220b09f813c1d13b72bf3be060d8&p=1"
)
