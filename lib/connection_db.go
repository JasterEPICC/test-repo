package lib

import (
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func ConnectionDB() *sqlx.DB {
	InitConfig()

	db, err := sqlx.Open(getDriverString(), getConnectionString())

	if err != nil {
		LogError(fmt.Sprintf("Cannot connect to database: %v", err), true)
	}

	db.SetConnMaxLifetime(time.Duration(viper.GetInt("db.maxLifeTime")) * time.Minute)
	db.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
	db.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))

	return db
}

func getConnectionString() string {
	mssqlHost := viper.Get("db.host")
	mssqlUsername := viper.Get("db.username")
	mssqlPassword := viper.Get("db.password")
	mssqlDatabase := viper.Get("db.database")

	return fmt.Sprintf("sqlserver://%v:%v@%v?database=%s", mssqlUsername, mssqlPassword, mssqlHost, mssqlDatabase)
}

func getDriverString() string {
	return fmt.Sprint(viper.Get("db.driver"))
}

func ConnectionRedis() *redis.Client {
	redisHost := viper.Get("redis.host")
	redisPort := viper.Get("redis.port")
	redisPassword := viper.Get("redis.password")

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", redisHost, redisPort),
		Password: fmt.Sprint(redisPassword),
	})
}
