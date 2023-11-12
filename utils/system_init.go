package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConfig() bool {
	viper.SetConfigName("app")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("init config err:", err)
	}

	fmt.Println("init config success")

	return true
}

func InitDB() *mongo.Client {
	fmt.Println("db uri:", viper.Get("mongo"))

	clientOptions := options.Client().ApplyURI(viper.GetString("mongo.uri"))
	var debug = viper.GetBool("mongo.debug")
	if debug {
		cmdMonitor := &event.CommandMonitor{
			Started: func(_ context.Context, evt *event.CommandStartedEvent) {
				log.Print(evt.Command)
			},
		}
		clientOptions.SetMonitor(cmdMonitor)
	}

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	// DB = client

	return client
}

func InitRDB() *redis.Client {
	fmt.Println("redis config:", viper.Get("redis"))
	rdb := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.db"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})

	// 测试可用性
	pong, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatal("Redis connect err: ", err)
	}

	fmt.Println("Connected to Redis successfully!", pong)
	return rdb
}

var _ = InitConfig()
var DB *mongo.Client = InitDB()
var RDB *redis.Client = InitRDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}
