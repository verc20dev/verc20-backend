package cmd

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ethsyncer",
	Short: "ethsyncer is a tool to sync the transactions from ethereum",
	Long:  `ethsyncer is a tool to sync the transactions from ethereum`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
		if DryRun {
			fmt.Println("dry-run mode")
		}
	},
}

var DryRun bool

func init() {
	//viper.SetEnvPrefix("VERC")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("rpcUrl", "http://localhost")
	viper.SetDefault("dbHost", "localhost")
	viper.SetDefault("dbPort", "5432")
	viper.SetDefault("dbUser", "postgres")
	viper.SetDefault("dbPassword", "postgres")
	viper.SetDefault("dbName", "postgres")
	viper.SetDefault("dbSslMode", "disable")
	viper.SetDefault("dbTimezone", "Asia/Shanghai")
	viper.SetDefault("syncInterval", 10)
	viper.SetDefault("indexInterval", 10)
	viper.SetDefault("port", 8080)

	viper.AutomaticEnv()

	log.Info("DbHost: ", viper.GetString("dbHost"))
	log.Info("DbHost ENV: ", os.Getenv("dbHost"))

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.Is(err, &configFileNotFoundError) {
			log.Warning("Config file not found, using default config")
		} else {
			log.Fatal(err)
		}
	}

	log.Info("DbHost: ", viper.GetString("dbHost"))
	log.Info("DbHost ENV: ", os.Getenv("dbHost"))

	rootCmd.PersistentFlags().BoolVar(&DryRun, "dry-run", false, "run without save txs to db")
	rootCmd.PersistentFlags().String("rpc-url", "http://localhost", "ethereum rpc url")
	err := viper.BindPFlag("rpcUrl", rootCmd.PersistentFlags().Lookup("rpc-url"))
	if err != nil {
		log.Fatal(err)
	}
	rootCmd.PersistentFlags().String("db-host", "localhost", "database host")
	err = viper.BindPFlag("dbHost", rootCmd.PersistentFlags().Lookup("db-host"))
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().String("db-port", "5432", "database port")
	err = viper.BindPFlag("dbPort", rootCmd.PersistentFlags().Lookup("db-port"))
	if err != nil {
		log.Fatal(err)
	}
	rootCmd.PersistentFlags().String("db-user", "postgres", "database user")
	err = viper.BindPFlag("dbUser", rootCmd.PersistentFlags().Lookup("db-user"))
	if err != nil {
		log.Fatal(err)
	}
	rootCmd.PersistentFlags().String("db-password", "postgres", "database password")
	err = viper.BindPFlag("dbPassword", rootCmd.PersistentFlags().Lookup("db-password"))
	if err != nil {
		log.Fatal(err)
	}
	rootCmd.PersistentFlags().String("db-name", "postgres", "database name")
	err = viper.BindPFlag("dbName", rootCmd.PersistentFlags().Lookup("db-name"))
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(indexCmd)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
