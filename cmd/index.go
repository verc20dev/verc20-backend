package cmd

import (
	"ethsyncer/pkg/indexer"
	"ethsyncer/pkg/orm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Index the transactions",
	Long:  "Index the transactions to update token and balance info",
	Run:   index,
}

func init() {
	indexCmd.PersistentFlags().BoolP("watch", "w", false, "watch mode")
	err := viper.BindPFlag("indexerWatch", indexCmd.PersistentFlags().Lookup("watch"))
	if err != nil {
		log.Fatal(err)
	}

	indexCmd.PersistentFlags().Int("interval", 10, "sync interval")
	err = viper.BindPFlag("indexInterval", indexCmd.PersistentFlags().Lookup("interval"))
	if err != nil {
		log.Fatal(err)
	}
}

func index(cmd *cobra.Command, args []string) {
	dbHost := viper.GetString("dbHost")
	dbPort := viper.GetString("dbPort")
	dbUser := viper.GetString("dbUser")
	dbPassword := viper.GetString("dbPassword")
	dbName := viper.GetString("dbName")
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort
	err := orm.InitDbClient(dsn)
	if err != nil {
		log.Fatal(err)
	}
	dbc := orm.GetDbClient()

	im := indexer.NewManager(dbc)

	if viper.GetBool("indexerWatch") {
		im.Start(viper.GetInt("syncInterval"))
	} else {
		err := im.IndexOnce()
		if err != nil {
			log.Fatal(err)
		}
	}
}
