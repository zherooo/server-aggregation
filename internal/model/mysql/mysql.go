package mysql

import (
	"fmt"
	"os"
	"server-aggregation/config"
	"server-aggregation/pkg/log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var once sync.Once

var dbMap map[string]*gorm.DB

func Init() {
	var err error
	once.Do(func() {
		dbMap = make(map[string]*gorm.DB)
		dbConfigs := config.GetStringMap("mysql")
		for dbName, _ := range dbConfigs {
			var tmpDB *gorm.DB
			tmpDB, err = gorm.Open("mysql", config.GetString(fmt.Sprintf("mysql.%s.dsn", dbName)))
			if err == nil {
				fmt.Println("\033[1;30;42m[info]\033[0m db [" + dbName + "] connect success")
			} else {
				fmt.Printf("\033[1;30;41m[error]\033[0m db ["+dbName+"] connect error: %s", err.Error())
				os.Exit(1)
			}
			tmpDB.SetLogger(log.New())
			tmpDB.LogMode(config.GetBool(fmt.Sprintf("mysql.%s.log_model", dbName)))
			tmpDB.DB().SetConnMaxLifetime(time.Minute * time.Duration(config.GetInt(fmt.Sprintf("mysql.%s.conn_max_lifetime", dbName))))
			tmpDB.DB().SetMaxIdleConns(config.GetInt(fmt.Sprintf("mysql.%s.max_idle_conns", dbName)))
			tmpDB.DB().SetMaxOpenConns(config.GetInt(fmt.Sprintf("mysql.%s.max_open_conns", dbName)))
			tmpDB.DB().SetConnMaxLifetime(time.Second * 10)
			dbMap[dbName] = tmpDB
		}
	})
}

// GetUserDB 获取用户相关的db配置
func GetUserDB() *gorm.DB {
	return dbMap["user"]
}
