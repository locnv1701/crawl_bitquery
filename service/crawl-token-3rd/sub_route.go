package review

import (
	"crawl-token-3rd-service/service/crawl-token-3rd/bitquery"

	"github.com/go-chi/chi"
)

var CrawlToken3rdServiceSubRoute = chi.NewRouter()

func init() {

	// fmt.Println("sub route")
	// crawler.Crawl()

	bitquery.Call()
	// solana.Call()

	// handleduplicate.HandleDuplicateToken()

	CrawlToken3rdServiceSubRoute.Group(func(r chi.Router) {

	})
}
