/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-31 下午5:17
***********************************************/
package db
import (
	"gopkg.in/mgo.v2"
	"time"
	"github.com/xiaomeng79/gin_log/conf"
	"errors"
	"github.com/xiaomeng79/gin_log/log"
)

var session *mgo.Session

func MgoInit() {
	var err error
	dialInfo := &mgo.DialInfo{
	Addrs:     []string{conf.MongoHosts},
	Direct:    false,
	Timeout:   time.Second * 5,
	Username: conf.MongoUserName,
	Password: conf.MongoPassword,
	PoolLimit: conf.MongoPoolLimit, // Session.SetPoolLimit
	}
	session, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Fatal("mongo connect fail %v",err)
	}
	session.SetMode(mgo.Monotonic, true)

}

type SessionStore struct {
	session *mgo.Session
	dbname string
}

//获取数据库的collection
func (d * SessionStore) C(name string) *mgo.Collection {
	return d.session.DB(d.dbname).C(name)
}

//为每一HTTP请求创建新的DataStore对象
func NewSessionStore() * SessionStore {
	ds := &SessionStore{
		session: session.Copy(),
		dbname:conf.MongoDb,
	}
	return ds
}

//设置数据库名称,不设置使用默认的
func (d *SessionStore) SetDbName(dbname string) error{
	if dbname != "" {
		d.dbname = dbname
		return nil
	}
	return errors.New("数据库名称不能为空")
}

//关闭复制的连接
func (d * SessionStore) Close() {
	d.session.Close()
}

//未知错误
func (d * SessionStore) GetErrNotFound() error {
	return mgo.ErrNotFound
}

//关闭原始连接
func MgoClose() {
	session.Close()
}
