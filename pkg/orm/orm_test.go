package orm

import (
	"testing"
)

func TestGetDbClient(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	err := InitDbClient(dsn)
	if err != nil {
		t.Fatal(err)
	}
	db := GetDbClient()
	t.Log(db)
	var tx = TxModel{}
	err = db.First(&tx, 1).Error
	if err != nil {
		t.Log(err)
	}
	t.Log(tx)
	type Result struct {
		Name    string
		Holders int32
	}
	var res []Result
	err = db.
		Model(&TokenHolderModel{}).
		Select("token_name as name, count(distinct address) as holders").
		Group("token_name").
		Find(&res).
		Error
	if err != nil {
		t.Log(err)
	}
	t.Log(res)

	type TokenHistoryInfo struct {
		Method      string `json:"method,omitempty"`
		From        string `json:"from,omitempty"`
		To          string `json:"to,omitempty"`
		BlockNumber int32  `json:"block_number,omitempty"`
		Quantity    string `json:"quantity,omitempty"`
		CreatedAt   string `json:"created_at,omitempty"`
		CreationTx  string `json:"creation_tx,omitempty"`
	}

	var histories []TokenHistoryInfo

	err = db.
		Order("tx_id desc").
		Model(&HistoryModel{}).
		Select("histories.\"type\" as method, t.\"from\", t.\"to\", block_num, amount as quantity, timestamp as created_at, hash as creation_tx").
		Joins("left join transactions t on t.id = histories.tx_id").
		Scan(&histories).
		Error
	if err != nil {
		t.Log(err)
	}
	t.Log(histories)
}


