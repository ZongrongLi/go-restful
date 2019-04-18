package model

import (
	"fmt"

	"github.com/lexkong/log"
	"github.com/spf13/viper"

	// MySQL driver.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Write *gorm.DB
	Read  *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")
	log.Infof("======================================", username, " ", password, addr, name)
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

// used for cli
func InitWritefDB() *gorm.DB {
	return openDB(viper.GetString("write_db.username"),
		viper.GetString("write_db.password"),
		viper.GetString("write_db.addr"),
		viper.GetString("write_db.name"))
}

func GetWriteDB() *gorm.DB {
	return InitWritefDB()
}

func InitReadrDB() *gorm.DB {
	return openDB(viper.GetString("read_db.username"),
		viper.GetString("read_db.password"),
		viper.GetString("read_db.addr"),
		viper.GetString("read_db.name"))
}

func GetReadDB() *gorm.DB {
	return InitReadrDB()
}

func (db *Database) Init() {
	DB = &Database{
		Write: GetWriteDB(),
		Read:  GetReadDB(),
	}
}

func (db *Database) Close() {
	DB.Write.Close()
	DB.Read.Close()
}
