package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sort"

	"github.com/go-graphql"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	graphqlHandler "github.com/graphql-go/handler"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	logger    *log.Logger
	dbPool    *sql.DB
	cachePool *redis.Pool

	cfgFile string

	handler                                      http.Handler
	schemaQuestion, schemaValidate, schemaVerify *graphql.Schema
	templates                                    *template.Template
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initLogger)
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.gareng.toml)")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	viper.SetConfigType("toml")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			panic(err)
		}
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName(".gareng")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
