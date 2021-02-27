package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
	"github.com/remnv/go-boiler/internal/config"
	"github.com/remnv/go-boiler/internal/helpers"
	"github.com/remnv/go-boiler/internal/repository/gormrepo"
	"github.com/sirupsen/logrus"
)

func GetMySqlDb() *gorm.DB {
	dbName := config.Instance().DB.Database
	return MySqlDbConn(&dbName)
}

func MySqlDbConn(dbName *string) *gorm.DB {
	dbURL := getMysqlUrl(dbName)
	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic(fmt.Sprintf("error: %v for %v", err.Error(), dbURL))
	}
	if config.Instance().DB.Debug {
		db.LogMode(true)
	}
	db.SetLogger(customDbLogger{})

	db.DB().SetConnMaxLifetime(time.Second * time.Duration(config.Instance().DB.MaxConnLifeTime))
	db.DB().SetMaxOpenConns(config.Instance().DB.MaxOpenConnections)
	db.DB().SetMaxIdleConns(config.Instance().DB.MaxIdleConnections)

	return db
}

func CreateMySqlDb(dbName string) error {
	dbConn := MySqlDbConn(nil)
	defer dbConn.Close()
	return dbConn.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
}

func MigrateMysqlDb(db *sql.DB, migrationFolder *string, rollback bool, versionToForce int) error {
	dbConfig := config.Instance().DB

	var validMigrationFolder = dbConfig.Migration.Path
	if !helpers.IsZero(migrationFolder) {
		validMigrationFolder = *migrationFolder
	}

	if validMigrationFolder == "" {
		return fmt.Errorf("empty migration folder")
	}
	logrus.Infof("Migration folder: %s", validMigrationFolder)

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		logrus.WithError(err).Warning("Error when instantiating driver")
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+validMigrationFolder,
		dbConfig.Client,
		driver)
	if err != nil {
		logrus.WithError(err).Warning("Error when instantiating migrate")
		return err
	}
	if rollback {
		logrus.Info("About to Rolling back 1 step")
		err = m.Steps(-1)
	} else if versionToForce != -1 {
		logrus.Info(fmt.Sprintf("About to force version %d", versionToForce))
		err = m.Force(versionToForce)
	} else {
		logrus.Info("About to run migration")
		err = m.Up()
	}
	if err != nil {
		if err != migrate.ErrNoChange {
			return err
		}
	}

	return nil
}

func CloseDB(db *gorm.DB) {
	if db == nil {
		return
	}
	err := db.Close()
	if err != nil {
		logrus.Warnf("Error when closing db: %s", err)
	}
}

func TruncateNonRefTables(db *gorm.DB) error {
	models := []interface{}{
		gormrepo.Player{},
	}
	for _, v := range models {
		tableName := db.NewScope(v).TableName()
		err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName)).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func getMysqlUrl(dbName *string) string {
	dbConfig := config.Instance().DB

	dbNameTmp := ""
	if dbName != nil {
		dbNameTmp = *dbName
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true&multiStatements=true", dbConfig.User,
		dbConfig.Password, dbConfig.Host, dbConfig.Port, dbNameTmp)
}

type customDbLogger struct{}

func (customDbLogger) Print(v ...interface{}) {
	tokLen := len(v)
	if tokLen == 0 {
		return
	}

	if tokLen == 3 {
		msg := v[2]
		codeLine := v[1]
		logrus.WithField("code", codeLine).Debug(msg)
		return
	}

	if tokLen == 6 {
		//logType := v[0] // sql|log
		codeLine := v[1]
		elapsed := v[2]
		query := v[3]
		params := v[4]
		if paramList, ok := params.([]interface{}); ok {
			newList := make([]interface{}, 0)
			for _, val := range paramList {
				newList = append(newList, helpers.Val(val))
			}
			params = newList
		}
		logrus.
			WithField("code", codeLine).
			WithField("params", params).
			WithField("elapsed", elapsed).
			Debug(query)

		return
	}

	logrus.Debug(v...)
}
