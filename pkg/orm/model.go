package orm

type Tabler interface {
	TableName() string
}

type TxModel struct {
	ID        uint64 `gorm:"primaryKey"`
	Hash      string
	From      string
	To        string
	BlockNum  uint64
	Index     uint64
	Timestamp uint64
	Input     string
}

func (TxModel) TableName() string {
	return "transactions"
}

type StatusModel struct {
	ID                   uint64 `gorm:"primaryKey"`
	SyncedBlockNum       uint64
	ImportedBlockNum     uint64
	EventFetchedBlockNum uint64
}

func (StatusModel) TableName() string {
	return "status"
}

type TokenInfoModel struct {
	ID          uint64 `gorm:"primaryKey"`
	Name        string
	TotalSupply string
	Decimals    int
	TotalMinted string
	TotalTxs    uint64
	Limit       string
	CreatedBy   string
	CreationTx  string
	CreatedAt   uint64
	MintedOutAt uint64
	StartBlock  uint64
	Type        string
	Duration    uint64
	TotalVolume string
	DailyVolume string
	FloorPrice  string
	SortWeight  int
	IsOfficial  bool
	IsVerified  bool
}

func (TokenInfoModel) TableName() string {
	return "token_info"
}

type TokenHolderModel struct {
	ID        uint64 `gorm:"primaryKey"`
	TokenID   uint64
	TokenName string
	Address   string
	Balance   string
}

func (TokenHolderModel) TableName() string {
	return "token_holders"
}

type HistoryModel struct {
	ID        uint64 `gorm:"primaryKey"`
	TokenName string
	TxId      uint64
	Type      string
	Amount    string
}

func (HistoryModel) TableName() string {
	return "histories"
}

type OrderModel struct {
	ID             uint64 `gorm:"primaryKey"`
	Maker          string
	TokenName      string
	Amount         string
	UnitPrice      string
	Tx             string
	CreationTime   uint64
	ExpirationTime uint64
	Signature      string
	Canceled       bool
	CanceledAt     uint64
	Executed       bool
	ExecutedAt     uint64
	Taker          string
	Input          string
}

func (OrderModel) TableName() string {
	return "orders"
}

type TradingActivityModel struct {
	ID        uint64 `gorm:"primaryKey"`
	TokenName string
	OrderId   uint64
	Type      string
	Amount    string
	UnitPrice string
	CreatedAt uint64
	From      string
	To        string
	TxHash    string
}

func (TradingActivityModel) TableName() string {
	return "trading_activities"
}
