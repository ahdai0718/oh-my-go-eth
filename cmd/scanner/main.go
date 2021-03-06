package main

import (
	"flag"

	"github.com/ahdai0718/oh-my-go-eth/internal/app/scanner"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var _ = flag.String("run_mode", "release", "Game server run mode (dev|debug|release|test)")

var _ = flag.String("database_host", "", "Database host")
var _ = flag.String("database_port", "", "Database port")
var _ = flag.String("database_schema", "", "Database schema")
var _ = flag.String("database_username", "", "Database username")
var _ = flag.String("database_password", "", "Database password")

var _ = flag.String("redis_host", "localhost", "Redis host")
var _ = flag.String("redis_port", "6379", "Redis host")

var _ = flag.String("eth_data_seed_url", "https://data-seed-prebsc-2-s3.binance.org:8545/", "ETH API endpoint")

const (
	ScanLimit = 20
)

func main() {
	flag.Parse()
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	scanner.Init(viper.GetString("eth_data_seed_url"))
	scanner.SetScanLimit(ScanLimit)
	scanner.Scan()

	for {
		select {}
	}
}
