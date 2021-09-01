package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/ahdai0718/oh-my-go-eth/internal/app/server/eth"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var _ = flag.String("run_mode", "release", "Game server run mode (dev|debug|release|test)")
var _ = flag.String("gin_mode", "release", "Gin http server run mode (debug|release|test)")

var _ = flag.String("server_host", "0.0.0.0", "Server host")
var _ = flag.String("server_port", "40001", "Server host port")

var _ = flag.String("database_host", "", "Database host")
var _ = flag.String("database_port", "", "Database port")
var _ = flag.String("database_schema", "", "Database schema")
var _ = flag.String("database_username", "", "Database username")
var _ = flag.String("database_password", "", "Database password")

var _ = flag.String("eth_data_seed_url", "https://data-seed-prebsc-2-s3.binance.org:8545/", "ETH API endpoint")

// @title ETH Service API
// @version 1.0
// @description

// @BasePath /api/v1
func main() {
	flag.Parse()
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	eth.Init()

	serverHost := viper.GetString("server_host")
	serverPort := viper.GetInt("server_port")

	glog.Info(viper.AllSettings())

	gin.SetMode(viper.GetString("gin_mode"))

	eth.SetDataSeedURL(viper.GetString("eth_data_seed_url"))

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := router.Group("/api/v1")
	{

		ethGroup := api.Group("/eth")
		ethGroup.Use()
		{
			blockGroup := ethGroup.Group("/blocks")
			blockGroup.Use()
			{
				blockGroup.GET("", eth.ServerHandlerBlockList)
				blockGroup.GET("/:id", eth.ServerHandlerBlock)
			}

			transactionGroup := ethGroup.Group("/transaction")
			transactionGroup.Use()
			{
				transactionGroup.GET("/:txHash", eth.ServerHandlerTransaction)
			}
		}
	}

	router.Use(static.Serve("/swagger-static", static.LocalFile("../../../api/docs", true)))
	url := ginSwagger.URL(fmt.Sprintf("http://%s:%d/swagger-static/swagger.json", serverHost, serverPort))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", serverHost, serverPort),
		Handler:        router,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	glog.Info(server)

	err := server.ListenAndServe()
	glog.Error(err)
}
