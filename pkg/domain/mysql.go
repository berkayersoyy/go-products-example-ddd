package domain

import "github.com/jinzhu/gorm"

//MysqlClient Mysql client
type MysqlClient interface {
	GetClient() *gorm.DB
}
