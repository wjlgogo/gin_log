/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-24 上午11:29
***********************************************/
package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xiaomeng79/gin_log/conf"
	"strconv"
	"github.com/xiaomeng79/gin_log/log"
)

var Mysql *sql.DB

//初始化连接
func MysqlInit() {
	var err error
	dataSourceName := conf.MysqlUserName + ":" + conf.MysqlPassword + "@tcp(" + conf.MysqlAddr + ":" + strconv.Itoa(conf.MysqlPort) +
	")/" + conf.MysqlDb + "?parseTime=true"
	Mysql, err = sql.Open("mysql",dataSourceName)
	if err != nil {
		log.Fatal("mysql open fail %v",err)
	}
	Mysql.SetMaxIdleConns(conf.MysqlIdleConn)
	Mysql.SetMaxOpenConns(conf.MysqlMaxConn)
	err = Mysql.Ping()
	if err != nil {
		log.Fatal("mysql connect fail %v",err)
	}
}