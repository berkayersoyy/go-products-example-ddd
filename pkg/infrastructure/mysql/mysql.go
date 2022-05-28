package mysql

import (
	"fmt"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/application/util/config"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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

	conf, err := config.LoadConfig("./")
	if err != nil {
		panic(err)
	}
	a := conf.MysqlDSN
	fmt.Println(a)
	db, err := gorm.Open("mysql", conf.MysqlDSN)
	//ctx := context.Background()
	//if err := retry.Fibonacci(ctx, 1*time.Second, func(ctx context.Context) error {
	//	if err := db.DB().Ping(); err != nil {
	//		fmt.Println(err)
	//
	//		return retry.RetryableError(err)
	//	}
	//	return nil
	//}); err != nil {
	//	log.Fatal(err)
	//}
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(5)

	db.AutoMigrate(&domain.Product{})
	db.AutoMigrate(&domain.User{})

	return db
}
