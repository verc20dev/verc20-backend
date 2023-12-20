package cmd

import (
	"ethsyncer/pkg/orm"
	"ethsyncer/pkg/syncer"
	"ethsyncer/pkg/web3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync the transactions",
	Long:  `Sync the transactions from ethereum to database`,
	Run:   sync,
}

func init() {
	syncCmd.PersistentFlags().BoolP("watch", "w", false, "watch mode")
	err := viper.BindPFlag("syncerWatch", syncCmd.PersistentFlags().Lookup("watch"))
	if err != nil {
		log.Fatal(err)
	}

	syncCmd.PersistentFlags().Int("interval", 10, "sync interval")
	err = viper.BindPFlag("syncInterval", syncCmd.PersistentFlags().Lookup("interval"))
	if err != nil {
		log.Fatal(err)
	}
}

func sync(cmd *cobra.Command, args []string) {
	rpcUrl := viper.GetString("rpcUrl")

	wc, err := web3.GetWeb3Client(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbHost := viper.GetString("dbHost")
	dbPort := viper.GetString("dbPort")
	dbUser := viper.GetString("dbUser")
	dbPassword := viper.GetString("dbPassword")
	dbName := viper.GetString("dbName")
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort
	err = orm.InitDbClient(dsn)
	if err != nil {
		log.Fatal(err)
	}
	dbc := orm.GetDbClient()

	sm := syncer.NewSyncManager(wc, dbc)

	if viper.GetBool("syncerWatch") {
		sm.Start(viper.GetInt("syncInterval"))
	} else  {
		err := sm.SyncOnce()
		if err != nil {
			log.Fatal(err)
		}
	}

}
