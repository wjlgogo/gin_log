package conf

import (
	"log"
)

var (
	// log conf
	LogFlag = log.LstdFlags
	LogPath = ""
	LogLevel = "debug"


	//server conf
	ServerMode = "debug" //debug release error fatal
	ServerAddr = "0.0.0.0"
	ServerPort = 8080


	//mysql conf
	MysqlAddr = "127.0.0.1"
	MysqlPort = 3306
	MysqlDb = "test"
	MysqlUserName = "root"
	MysqlPassword = "root"
	MysqlIdleConn = 4
	MysqlMaxConn = 20

	//mongodb conf
	MongoHosts = "127.0.0.1:27017"
	MongoDb = "test"
	MongoUserName = ""
	MongoPassword = ""
	MongoPoolLimit = 4096

	//email conf
	EmailHost = "smtp.mxhichina.com"
	EmailPort = "465"
	EmailUserName = ""
	EmailNickName = ""
	EmailPassword = ""


)
