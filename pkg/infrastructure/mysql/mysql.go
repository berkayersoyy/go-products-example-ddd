package mysql

import (
	"context"
	"fmt"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-retry"
	"log"
	"os"
	"time"
)

type mysqlClient struct {
	SingletonMysql *gorm.DB
}

func ProvideMysqlClient() domain.MysqlClient {
	return &mysqlClient{SingletonMysql: InitDb()}
}

func (m *mysqlClient) GetClient() *gorm.DB {
	return m.SingletonMysql
}

func InitDb() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dsn := os.Getenv("MYSQL_DSN")
	ctx := context.Background()
	var db *gorm.DB
	if err := retry.Fibonacci(ctx, 1*time.Second, func(ctx context.Context) error {
		db, err = gorm.Open("mysql", dsn)
		if err != nil {
			fmt.Println(err)
			return retry.RetryableError(err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(5)

	db.AutoMigrate(&domain.Product{})
	db.AutoMigrate(&domain.User{})

	return db
}
