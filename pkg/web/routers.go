package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Route is the information for every URI.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	router := gin.Default()
	// enable cors
	router.Use(cors.Default())
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// DefaultHandleFunc Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the HolderAPI part of the API
	HolderAPI HolderAPI
	// Routes for the TokenAPI part of the API
	TokenAPI TokenAPI
	// Routes for the MarketAPI part of the API
	MarketAPI MarketAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{
		{
			"GetHolder",
			http.MethodGet,
			"/holders/:address",
			handleFunctions.HolderAPI.GetHolder,
		},
		{
			"GetHolder",
			http.MethodGet,
			"/holders/:address/histories",
			handleFunctions.HolderAPI.ListHolderHistories,
		},
		{
			"GetToken",
			http.MethodGet,
			"/tokens/:name",
			handleFunctions.TokenAPI.GetToken,
		},
		{
			"ListToken",
			http.MethodGet,
			"/tokens",
			handleFunctions.TokenAPI.ListToken,
		},
		{
			"ListTokenHistories",
			http.MethodGet,
			"/tokens/:name/histories",
			handleFunctions.TokenAPI.ListTokenHistories,
		},
		{
			"ListTokenHolders",
			http.MethodGet,
			"/tokens/:name/holders",
			handleFunctions.TokenAPI.ListTokenHolders,
		},
		{
			"GetStatus",
			http.MethodGet,
			"/status",
			handleFunctions.TokenAPI.GetStatus,
		},
		{
			"ListMarketTokens",
			http.MethodGet,
			"/market/tokens",
			handleFunctions.MarketAPI.ListMarketTokens,
		},
		{
			"GetMarketTokenDetail",
			http.MethodGet,
			"/market/tokens/:name",
			handleFunctions.MarketAPI.GetMarketTokenDetail,
		},
		{
			"ListOrders",
			http.MethodGet,
			"/market/orders",
			handleFunctions.MarketAPI.ListOrders,
		},
		{
			"CreateOrder",
			http.MethodPost,
			"/market/orders",
			handleFunctions.MarketAPI.CreateOrder,
		},
		{
			"GetOrderDetail",
			http.MethodGet,
			"/market/orders/:id",
			handleFunctions.MarketAPI.GetOrderDetail,
		},
		{
			"ExecuteOrder",
			http.MethodPost,
			"/market/orders/:id/execute",
			handleFunctions.MarketAPI.ExecuteOrder,
		},
		{
			"FreezeOrder",
			http.MethodPost,
			"/market/orders/:id/freeze",
			handleFunctions.MarketAPI.FreezeOrder,
		},
		{
			"CancelOrder",
			http.MethodPost,
			"/market/orders/:id/cancel",
			handleFunctions.MarketAPI.CancelOrder,
		},
		//{
		//	"UpdateOrder",
		//	http.MethodPost,
		//	"/orders/:id/update",
		//	handleFunctions.OrderAPI.UpdateOrder,
		//},
		{
			"ListTradingActivities",
			http.MethodGet,
			"/market/activities",
			handleFunctions.MarketAPI.ListTradingActivities,
		},
		{
			"GetMarketTokenPrice",
			http.MethodGet,
			"/market/tokens/:name/price",
			handleFunctions.MarketAPI.GetMarketTokenPrice,
		},
	}
}
