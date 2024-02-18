package web

type TokenHolding struct {
	Name    string `json:"name,omitempty"`
	Balance string `json:"balance,omitempty"`
}

type HolderDetail struct {
	Tokens []TokenHolding `json:"tokens,omitempty"`
}

type ListTokenHistoryResponse struct {
	Total int64              `json:"total,omitempty"`
	Data  []TokenHistoryInfo `json:"data,omitempty"`
}

type ListTokenHolderResponse struct {
	Total int64             `json:"total,omitempty"`
	Data  []TokenHolderInfo `json:"data,omitempty"`
}

type ListTokenResponse struct {
	Total int64         `json:"total,omitempty"`
	Data  []TokenDetail `json:"data,omitempty"`
}

type TokenDetail struct {
	Name         string `json:"name,omitempty"`
	TotalSupply  string `json:"total_supply,omitempty"`
	Decimals     int    `json:"decimals,omitempty"`
	Minted       string `json:"minted"`
	Holders      int64  `json:"holders"`
	Transactions uint64 `json:"transactions"`
	CreationTx   string `json:"creation_tx,omitempty"`
	CreatedAt    uint64 `json:"created_at,omitempty"`
	CreatedBy    string `json:"created_by,omitempty"`
	MintedOutAt  uint64 `json:"minted_out_at,omitempty"`
	Limit        string `json:"limit,omitempty"`
	StartBlock   uint64 `json:"start_block,omitempty"`
	Type         string `json:"type,omitempty"`
	Duration     uint64 `json:"duration,omitempty"`
	IsVerified   bool   `json:"isVerified"`
	IsOfficial   bool   `json:"isOfficial"`
}

type TokenHistoryInfo struct {
	Name        string `json:"name,omitempty"`
	Method      string `json:"method,omitempty"`
	From        string `json:"from,omitempty"`
	To          string `json:"to,omitempty"`
	BlockNumber int32  `json:"block_number,omitempty"`
	Quantity    string `json:"quantity,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	CreationTx  string `json:"creation_tx,omitempty"`
}

type TokenHolderInfo struct {
	Rank    int32  `json:"rank"`
	Address string `json:"address,omitempty"`
	Balance string `json:"balance,omitempty"`
}

type TokenInfo struct {
	Name         string `json:"name,omitempty"`
	TotalSupply  string `json:"total_supply,omitempty"`
	Minted       string `json:"minted,omitempty"`
	Holders      int32  `json:"holders"`
	Transactions int32  `json:"transactions,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	IsVerified   bool   `json:"isVerified"`
	IsOfficial   bool   `json:"isOfficial"`
}

type Status struct {
	LatestSyncedBlockNumber   uint64 `json:"latest_synced_block_number,omitempty"`
	LatestImportedBlockNumber uint64 `json:"latest_imported_block_number,omitempty"`
}

type OrderDetail struct {
	ID             uint64 `json:"id,omitempty"`
	Tick           string `json:"tick,omitempty"`
	Quantity       string `json:"quantity,omitempty"`
	UnitPrice      string `json:"unit_price,omitempty"`
	Owner          string `json:"owner,omitempty"`
	Status         string `json:"status,omitempty"`
	CreatedAt      uint64 `json:"created_at,omitempty"`
	ExpirationTime uint64 `json:"expiration_time,omitempty"`
	Taker          string `json:"taker,omitempty"`
	Input          string `json:"input,omitempty"`
	Signature      string `json:"signature,omitempty"`
	Sell           bool   `json:"sell,omitempty"`
}

type ListOrdersResponse struct {
	Total int64         `json:"total,omitempty"`
	Data  []OrderDetail `json:"data,omitempty"`
}

type CreateOrderRequest struct {
	Tick           string `json:"tick,omitempty" binding:"required"`
	Owner          string `json:"owner,omitempty" binding:"required"`
	Quantity       string `json:"quantity,omitempty" binding:"required"`
	UnitPrice      string `json:"unit_price,omitempty" binding:"required"`
	Tx             string `json:"tx,omitempty" binding:"required"`
	CreationTime   uint64 `json:"creation_time,omitempty" binding:"required"`
	ExpirationTime uint64 `json:"expiration_time,omitempty" binding:"required"`
	Signature      string `json:"signature,omitempty" binding:"required"`
	Input          string `json:"input,omitempty" binding:"required"`
	Sell           *bool  `json:"sell" binding:"required"`
}

type CancelOrderRequest struct {
	Tx string `json:"tx,omitempty" binding:"required"`
}

type TradingActivity struct {
	ID        uint64 `json:"id,omitempty"`
	Tick      string `json:"tick,omitempty"`
	Type      string `json:"type,omitempty"`
	Quantity  string `json:"quantity,omitempty"`
	UnitPrice string `json:"unit_price,omitempty"`
	CreatedAt uint64 `json:"created_at,omitempty"`
	From      string `json:"from,omitempty"`
	To        string `json:"to,omitempty"`
	Tx        string `json:"tx,omitempty"`
}

type ListTradingActivitiesResponse struct {
	Total int64             `json:"total,omitempty"`
	Data  []TradingActivity `json:"data,omitempty"`
}

type ToBeSignedOrder struct {
	Seller         string `json:"seller,omitempty"`
	ListTx         string `json:"listTx,omitempty"`
	Tick           string `json:"tick,omitempty"`
	Amount         string `json:"amount,omitempty"`
	Price          string `json:"price,omitempty"`
	ListingTime    uint64 `json:"listingTime,omitempty"`
	ExpirationTime uint64 `json:"expirationTime,omitempty"`
}

type MarketTokenDetail struct {
	Tick        string `json:"tick,omitempty"`
	TotalSupply string `json:"totalSupply,omitempty"`
	Holders     int64  `json:"holders"`
	FloorPrice  string `json:"floorPrice,omitempty"`
	TotalVolume string `json:"totalVolume,omitempty"`
	DailyVolume string `json:"dailyVolume,omitempty"`
	IsVerified  bool   `json:"isVerified"`
	IsOfficial  bool   `json:"isOfficial"`
}

type ListMarketTokensResponse struct {
	Total int64               `json:"total,omitempty"`
	Data  []MarketTokenDetail `json:"data,omitempty"`
}

type ExecuteOrderRequest struct {
	Tx string `json:"tx,omitempty" binding:"required"`
}

type FreezeOrderRequest struct {
	Signature string `json:"signature,omitempty" binding:"required"`
	Address   string `json:"address,omitempty" binding:"required"`
}

type ListHolderHistoriesResponse struct {
	Total int64              `json:"total,omitempty"`
	Data  []TokenHistoryInfo `json:"data,omitempty"`
}
