package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

func main() {
	Init()
}

func Init() {

	// 打开连接
	MysqlDb, MysqlDbErr := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/jamesjjsun_sun?charset=utf8&parseTime=true")
	//defer MysqlDb.Close();
	if MysqlDbErr != nil {
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	MysqlDb.SetMaxOpenConns(5)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(2)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(10*time.Second)


	query, MysqlDbErr := MysqlDb.Query("select * from export_record")

	for query.Next() {
		var id int
		var staff_name string
		var time time.Time
		var c_id string
		var batch_flag string
		MysqlDbErr := query.Scan(&id,&staff_name,&time,&c_id,&batch_flag)
		if 	MysqlDbErr != nil {
			panic(MysqlDbErr)
		}
		fmt.Printf("id:%d staff_name:%s time:%v c_id:%s batch_flag:%s",id,staff_name,time,c_id,batch_flag)
	}

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}
}