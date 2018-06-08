package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sort"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/graphql-go/go-graphql"
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
