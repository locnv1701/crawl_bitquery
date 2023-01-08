package service

import (
	"crawl-token-3rd-service/pkg/router"
	crawl3rd "crawl-token-3rd-service/service/crawl-token-3rd"
	"crawl-token-3rd-service/service/index"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {

	// Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)

	router.Router.Mount(router.RouterBasePath+"/crawl-token-3rd", crawl3rd.CrawlToken3rdServiceSubRoute)

}
