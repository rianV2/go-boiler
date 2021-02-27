// +build integration

package gormrepo_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/remnv/go-boiler/internal/config"
	"github.com/remnv/go-boiler/internal/helpers"
	"github.com/remnv/go-boiler/internal/storage"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

var dbName string

func TestMain(m *testing.M) {
	err := config.Load()
	if err != nil {
		fmt.Printf("Config error: %s\n", err.Error())
		os.Exit(1)
	}

	err = initLogging()
	if err != nil {
		fmt.Printf("Logging error: %s\n", err.Error())
		os.Exit(1)
	}

	conn, err := prepareDB()
	if err != nil {
		fmt.Printf("Prepare db error: %s", err.Error())
		os.Exit(1)
	}
	defer dropDB(conn)

	retCode := m.Run()
	os.Exit(retCode)
}

func initLogging() error {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
	log := logrus.StandardLogger()
	level, err := logrus.ParseLevel(config.Instance().LogLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)

	return err
}

func cleanDB(t *testing.T, db *gorm.DB) {
	defer func(t *testing.T) {
		err := db.Close()
		require.NoError(t, err)
	}(t)
	defer func(t *testing.T) {
		err := storage.TruncateNonRefTables(db)
		require.NoError(t, err)
	}(t)
}

func prepareDB() (dbConn *gorm.DB, err error) {
	dbName = "t_" + helpers.RandomString(10)
	err = storage.CreateMySqlDb(dbName)
	if err != nil {
		return
	}

	dbConn = storage.MySqlDbConn(&dbName)
	err = storage.MigrateMysqlDb(dbConn.DB(), nil, false, -1)
	if err != nil {
		return
	}

	return
}

func dropDB(dbConn *gorm.DB) error {
	defer dbConn.Close()
	err := dbConn.Exec(fmt.Sprintf("DROP DATABASE %s", dbName)).Error
	return err
}
