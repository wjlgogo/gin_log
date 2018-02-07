/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-22 下午4:41
***********************************************/
package main

import (
	"github.com/xiaomeng79/gin_log/routers"
	"github.com/xiaomeng79/gin_log/conf"
	"strconv"
	"github.com/xiaomeng79/gin_log/db"
	"github.com/xiaomeng79/gin_log/log"
)

var logger *log.Logger
//初始化一些配置
func init() {
	//初始化配置
	conf.Init()
	//初始化日志
	logger, err := log.New(conf.LogLevel, conf.LogPath, conf.LogFlag)
	if err != nil {
		panic(err)
	}
	log.Export(logger)
}

func main() {

	/*************************初始化*************************/
	//初始化mysql数据库连接
	//db.MysqlInit()
	//初始化mongo
	db.MgoInit()

	/***********************启动服务***********************/
	//注册路由
	r := routers.Reg()
	//启动服务
	r.Run(conf.ServerAddr + ":" + strconv.Itoa(conf.ServerPort))

	/**************************关闭*****************************/
	defer logger.Close() //关闭日志连接
	//defer db.Mysql.Close()//关闭mysql数据库连接
	defer db.MgoClose() //关闭mongo连接
}
