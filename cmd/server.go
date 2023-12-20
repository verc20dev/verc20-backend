package cmd

import (
	"ethsyncer/pkg/indexer"
	"ethsyncer/pkg/orm"
	"ethsyncer/pkg/syncer"
	"ethsyncer/pkg/web"
	"ethsyncer/pkg/web3"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the all-in-one server",
	Long:  "Start the sync service, index service and the api server",
	Run:   server,
}

func init() {
	serverCmd.PersistentFlags().Int("port", 8080, "server port")
	err := viper.BindPFlag("port", serverCmd.PersistentFlags().Lookup("port"))
	if err != nil {
		log.Fatal(err)
	}
	serverCmd.PersistentFlags().Bool("no-sync", false, "don't sync the transactions")
	err = viper.BindPFlag("noSyncServer", serverCmd.PersistentFlags().Lookup("no-sync"))
	if err != nil {
		log.Fatal(err)
	}
}

func server(cmd *cobra.Command, args []string) {
	log.Info("Starting server...")

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
	smDbc, err := orm.NewDbClient(dsn)
	if err != nil {
		log.Fatal(err)
	}
	imDbc, err := orm.NewDbClient(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// start syncer and indexer
	if viper.GetBool("noSyncServer") {
		log.Info("no-sync flag is set, skip the sync service")
	} else {
		sm := syncer.NewSyncManager(wc, smDbc)
		im := indexer.NewManager(imDbc)
		go sm.Start(viper.GetInt("syncInterval"))
		go im.Start(viper.GetInt("syncInterval"))
	}


	//wssUrl := viper.GetString("wssUrl")
	//wssWc, err := web3.GetWeb3Client(wssUrl)
	//marketContractAddress := viper.GetString("marketAddress")
	//mel, err := listener.NewMarkerEventListener(
	//	wssWc, marketContractAddress,
	//	protocol_default.ProcessOrderExecuted,
	//	protocol_default.ProcessOrderCanceled,
	//)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//go mel.ListenOrderExecuted()
	//go mel.ListenOrderCanceled()

	// start api server
	routes := web.ApiHandleFunctions{}
	router := web.NewRouter(routes)

	err = router.Run(fmt.Sprintf(":%d", viper.GetInt("port")))
	if err != nil {
		log.Fatal(err)
	}
}
