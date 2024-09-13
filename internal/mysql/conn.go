package mysql

import (
	"log"
	"net"
	"time"

	"github.com/chscz/backend-onboard-task/internal/config"
	"github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL(cfg config.MySQL) (*gorm.DB, error) {
	mysqlCfg := &mysql.Config{
		User:   cfg.UserName,
		Passwd: cfg.Password,
		Net:    "tcp",
		Addr:   net.JoinHostPort(cfg.Host, cfg.Port),
		DBName: cfg.DB,
		Params: map[string]string{
			"charset": "utf8mb4",
		},
		// Loc:                  time.UTC,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	var db *gorm.DB
	var err error
	retryCount := 20
	dsn := mysqlCfg.FormatDSN()
	for i := 0; i < retryCount; i++ {

		db, err = gorm.Open(gormmysql.Open(dsn), &gorm.Config{})
		if err == nil {
			return db, nil
		}

		log.Printf("Failed to connect to MySQL. Retrying (%d/%d)...\n", i+1, retryCount)
		time.Sleep(time.Second * 5)
		if i == retryCount-1 {
			panic(err)
		}
	}
	panic(err)
}
