package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/remnv/go-boiler/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-boiler",
	Short: "Golang boiler plate",
}

func init() {
	loadConfig()
	initLogging()
	registerCommands()
}

func main() {
	Execute()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err.Error())
		os.Exit(1)
	}
}

func loadConfig() {
	err := config.Load()
	if err != nil {
		logrus.Errorf("Config error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println(">> Configuration")
	fmt.Print(config.Instance().String())
}

func initLogging() *logrus.Logger {
	cfg := config.Instance()
	log := logrus.StandardLogger()
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
	if strings.ToLower(cfg.LogFormat) == "json" {
		log.SetFormatter(&logrus.JSONFormatter{})
	}
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)
	return log
}

func registerCommands() {
	rootCmd.AddCommand(Migrate())
	rootCmd.AddCommand(Server())
}
