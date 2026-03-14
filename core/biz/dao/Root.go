package dao

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var Dao *DB

type DB struct {
	pgdb *gorm.DB
	rdb  *redis.Client
}
