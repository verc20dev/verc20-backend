package syncer

import (
	"ethsyncer/pkg/orm"
	"ethsyncer/pkg/web3"
	"github.com/ethereum/go-ethereum/core/types"
	"net/http"
	"testing"
)

func TestSyncer_Sync(t *testing.T) {
	rpcUrl := "https://eth-goerli.g.alchemy.com/v2/UNcjTwO5pnN-6ua_dCM3vSqxaLdlWy8K"
	wc, err := web3.GetWeb3Client(rpcUrl)
	if err != nil {
		t.Fatal(err)
	}
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	err = orm.InitDbClient(dsn)
	if err != nil {
		t.Fatal(err)
	}
	dbc := orm.GetDbClient()

	syncer := NewSyncer(
		wc,
		dbc,
		10230006,
		10230008,
		0,
		false,
		func(tx *types.Transaction) bool {
			return http.DetectContentType(tx.Data()) != "application/octet-stream"
		},
	)
	if syncer.Sync() != nil {
		t.Fatal(err)
	}
}
