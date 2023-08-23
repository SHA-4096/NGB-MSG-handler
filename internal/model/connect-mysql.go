package model

import (
	config "NGB-MSG-handler/internal/conf"
	"NGB-MSG-handler/internal/util"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//
//迁移
//
func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Message{})
	if err != nil {
		return err
	}
	return nil

}

var db *gorm.DB
var err error

func MySqlInit() {
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		config.Config.Database.Mysql.Username,
		config.Config.Database.Mysql.Password,
		config.Config.Database.Mysql.Host,
		config.Config.Database.Mysql.Port,
		config.Config.Database.Mysql.DbName,
		config.Config.Database.Mysql.Timeout)
	//fmt.Println(dsn)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		util.MakeErrorLog("Failed when connecting ot mysql, error=" + err.Error())
		panic("Failed when connecting ot mysql, error=" + err.Error())
	}
	err = migrate(db)
	if err != nil {
		panic(err)
	}
	util.MakeInfoLog("[model]successfully initialized mysql")
}
