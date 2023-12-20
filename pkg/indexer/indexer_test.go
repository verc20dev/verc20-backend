package indexer

import (
	"ethsyncer/pkg/orm"
	"testing"
)

// test the indexer
func TestIndexer(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	err := orm.InitDbClient(dsn)
	if err != nil {
		t.Fatal(err)
	}
	dbc := orm.GetDbClient()

	// create a new indexer
	indexer := NewIndexer(dbc, 10266204, 10266220)
	err = indexer.Index()
	if err != nil {
		t.Fatal(err)
	}
}

// test the manager
func TestManager(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	err := orm.InitDbClient(dsn)
	if err != nil {
		t.Fatal(err)
	}
	dbc := orm.GetDbClient()

	// create a new manager
	manager := NewManager(dbc)
	err = manager.IndexOnce()
	if err != nil {
		t.Fatal(err)
	}
}