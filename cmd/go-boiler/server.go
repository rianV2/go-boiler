package main

import (
	"github.com/remnv/go-boiler/internal/config"
	"github.com/remnv/go-boiler/internal/storage"
	"github.com/remnv/go-boiler/internal/web"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Server() *cobra.Command {
	cliCommand := &cobra.Command{
		Use:   "server",
		Short: "Run REST API server",
		Run: func(cmd *cobra.Command, args []string) {
			// Init DB
			db := storage.GetMySqlDb()
			defer storage.CloseDB(db)

			// Start Http Server
			err := web.NewHttpServer(db, config.Instance()).Start()
			if err != nil {
				logrus.WithError(err).Error("Error starting web server")
			}
		},
	}

	return cliCommand
}
