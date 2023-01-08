package handleduplicate

import (
	"crawl-token-3rd-service/pkg/db"
	"crawl-token-3rd-service/pkg/log"
	"fmt"
)

type TokenDuplicate struct {
	Count     int
	Address   string
	Chainname string
}

type ListTokenDuplicate struct {
	Tokens []TokenDuplicate
}

func HandleDuplicateToken() {
	listTokenDuplicate := ListTokenDuplicate{}
	err := listTokenDuplicate.GetListTokenDuplicate()
	if err != nil {
		log.Println(log.LogLevelError, `listTokenDuplicate.GetListTokenDuplicate()`, err.Error())
	}

	fmt.Println(len(listTokenDuplicate.Tokens))

	for _, token := range listTokenDuplicate.Tokens {
		err := token.Delete()
		if err != nil {
			log.Println(log.LogLevelError, `token.Delete()`, err.Error())
		}
	}

}

func (token TokenDuplicate) Delete() error {
	query := `delete from crypto where cryptosrc = 'dex' and address = $1 and chainname = $2;`

	_, err := db.PSQL.Exec(query, token.Address, token.Chainname)
	return err
}

func (listTokenDuplicate *ListTokenDuplicate) GetListTokenDuplicate() error {
	query := `select count(*), address, chainname 
	from crypto where  cryptosrc = 'dex' or cryptosrc= 'bitquery' 
	group by (address, chainname) having count(*) >= 2;`

	rows, err := db.PSQL.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		token := TokenDuplicate{}

		err := rows.Scan(&token.Count, &token.Address, &token.Chainname)
		if err != nil {
			return err
		}

		listTokenDuplicate.Tokens = append(listTokenDuplicate.Tokens, token)
	}

	return nil

}
