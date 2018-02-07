package conf

import (
	"encoding/json"
	"io/ioutil"
	"github.com/xiaomeng79/gin_log/libs"
	"github.com/xiaomeng79/gin_log/log"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"os"
	"io"
)

var ServerConf struct {
	// log conf
	//LogFlag int
	LogPath string
	LogLevel string


	//server conf
	ServerMode string //debug release error fatal
	ServerAddr string
	ServerPort int


	//mysql conf
	MysqlAddr string
	MysqlPort int
	MysqlDb string
	MysqlUserName string
	MysqlPassword string
	MysqlIdleConn int
	MysqlMaxConn int

	//mongodb conf
	MongoHosts string
	MongoDb string
	MongoUserName string
	MongoPassword string
	MongoPoolLimit int

	//email conf
	EmailHost string
	EmailPort string
	EmailUserName string
	EmailNickName string
	EmailPassword string
}


func Init() {
	//如果文件不存在,就使用默认配置，否则根据是否有此配置来更新配置
	conffile,_ := filepath.Abs("conf/server.json")

	if ok,_ := libs.PathExists(conffile); ok {//有配置文件

		data, err := ioutil.ReadFile(conffile)
		if err != nil {
			log.Fatal("%v", err)
		}
		err = json.Unmarshal(data, &ServerConf)
		if err != nil {
			log.Fatal("%v", err)
		}
		//比较配置
		if ServerConf.LogPath != "" {
			LogPath = ServerConf.LogPath
		}

		if ServerConf.LogLevel != "" {
			LogLevel = ServerConf.LogLevel
		}

		if ServerConf.ServerMode != "" {
			ServerMode = ServerConf.ServerMode
		}

		if ServerConf.ServerAddr != "" {
			ServerAddr = ServerConf.ServerAddr
		}

		if ServerConf.ServerPort > 0 {
			ServerPort = ServerConf.ServerPort
		}

		if ServerConf.MysqlAddr != "" {
			MysqlAddr = ServerConf.MysqlAddr
		}

		if ServerConf.MysqlPort > 0 {
			MysqlPort = ServerConf.MysqlPort
		}

		if ServerConf.MysqlDb != "" {
			MysqlDb = ServerConf.MysqlDb
		}

		if ServerConf.MysqlUserName != "" {
			MysqlUserName = ServerConf.MysqlUserName
		}

		if ServerConf.MysqlPassword != "" {
			MysqlPassword = ServerConf.MysqlPassword
		}

		if ServerConf.MysqlIdleConn > 0 {
			MysqlIdleConn = ServerConf.MysqlIdleConn
		}

		if ServerConf.MysqlMaxConn > 0 {
			MysqlMaxConn = ServerConf.MysqlMaxConn
		}

		//mongo
		if ServerConf.MongoHosts != "" {
			MongoHosts = ServerConf.MongoHosts
		}

		if ServerConf.MongoDb != "" {
			MongoDb = ServerConf.MongoDb
		}
		if ServerConf.MongoUserName != "" {
			MongoUserName = ServerConf.MongoUserName
		}

		if ServerConf.MongoPassword != "" {
			MongoPassword = ServerConf.MongoPassword
		}

		if ServerConf.MongoPoolLimit > 0 {
			MongoPoolLimit = ServerConf.MongoPoolLimit
		}

		//email
		if ServerConf.EmailHost != "" {
			EmailHost = ServerConf.EmailHost
		}
		if ServerConf.EmailPort != "" {
			EmailPort = ServerConf.EmailPort
		}
		if ServerConf.EmailUserName != "" {
			EmailUserName = ServerConf.EmailUserName
		}
		if ServerConf.EmailNickName != "" {
			EmailNickName = ServerConf.EmailNickName
		}
		if ServerConf.EmailPassword != "" {
			EmailPassword = ServerConf.EmailPassword
		}


		//改变gin默认配置
		if ServerConf.ServerMode == "release" {
			os.Setenv(gin.ENV_GIN_MODE,"release")
			//gin.SetMode(gin.ReleaseMode)
			//关闭console颜色显示
			gin.DisableConsoleColor()
			// 写入日志到文件
			f := CreateLogFile()
			gin.DefaultWriter = io.MultiWriter(f)
			//gin.DefaultWriter = io.MultiWriter(f,os.Stdout)
		}


	}
	//没配置文件，使用默认配置

}

func CreateLogFile() *os.File {
	//logPath := "log"
	logPath := "log/reqlog"
	if ok,_ := libs.PathExists(logPath); !ok {
		os.MkdirAll(logPath,0777)
	}
	//f, err := os.Create(path.Join(logPath, "gin.log"))
	f, err := libs.CreateTimeFile(logPath,"log")
	if err != nil {
		log.Fatal("%v\n",err)
	}
		return f

}

