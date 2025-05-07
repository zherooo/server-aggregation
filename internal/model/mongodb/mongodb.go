package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"server-aggregation/config"
	"sync"
)

var once sync.Once

var mongoDBMap map[string]*mongo.Database

func Init() {
	once.Do(func() {
		mongoDBMap = make(map[string]*mongo.Database)
		dbConfigs := config.GetStringMap("mongo")
		for dbName, _ := range dbConfigs {
			client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.GetString(fmt.Sprintf("mongo.%s.dsn", dbName))))
			if err == nil {
				fmt.Println("\033[1;30;42m[info]\033[0m MongoDB [" + dbName + "] connect success")
			} else {
				fmt.Printf("\033[1;30;41m[error]\033[0m MongoDB ["+dbName+"] connect error: %s", err.Error())
				os.Exit(1)
			}
			mongoDBMap[dbName] = client.Database(dbName)
		}
	})
}

// GetMongoFirmwareDB 获取mongodb的链接
func GetMongoFirmwareDB() *mongo.Database {
	return mongoDBMap["firmware"]
}
func GetMongoSimilarityAnalysisDB() *mongo.Database {
	return mongoDBMap["similarity_analysis_db"]
}

func GetMongoFuzzDB() *mongo.Database {
	return mongoDBMap["fuzz_db"]
}
