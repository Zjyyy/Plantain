package main

import (
	"fmt"
	"os"
	"plantain/base"
	bSqlite "plantain/base/sqlite"
	"plantain/collector"
	"plantain/core"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	/**************初始化校验***********************/
	fmt.Println("开始加载配置文件...")
	config, err := base.LoadConfigFromIni("config.ini")
	if err != nil {
		fmt.Printf("加载配置文件失败：%v \n", err)
		return
	}
	fmt.Println("开始项目初始化校验...")
	databaseName := config.System.Database

	_, err = os.Stat(databaseName)
	var db *gorm.DB

	_, err = os.Stat(databaseName)

	db, dbErr := gorm.Open(sqlite.Open(databaseName))
	if dbErr != nil {
		fmt.Printf("打开SQLite连接错误：%v \n", err)
		return
	}

	if err != nil || os.IsNotExist(err) {
		fmt.Println("第一次部署项目，需要创建SQLite")
		db.AutoMigrate(
			&base.RtTable{},
			&base.PDriverInDatabase{},
		)
		bSqlite.CreateMockData(db)
	}
	/**************加载配置库***********************/
	pDriverArr, err := bSqlite.LoadAllDriver(db)
	if err != nil {
		fmt.Printf("加载配置库失败:%v\n", err)
	}

	//temp test
	// for _, item := range pDriverArr {
	// 	fmt.Printf("pDriverList: %v \n", item)
	// }

	/**************加载驱动插件***********************/
	collector.InitCollector(pDriverArr, core.New())
	/**************为配置库实时表建立内存结构***********************/
	/**************启动HttpServer***********************/
}
