package domain

import "github.com/jinzhu/gorm"

type MysqlClient interface {
	GetClient() *gorm.DB
}
