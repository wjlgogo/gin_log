/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-25 上午11:16
***********************************************/
package models

import (
	"github.com/xiaomeng79/gin_log/db"
	"gopkg.in/mgo.v2/bson"
)


type FatalEmails struct {
	Id bson.ObjectId `bson:"_id" `
	Emails string `bson:"emails" `
}

/**
获取APP信息
 */
func (this *FatalEmails) GetEmails() error {
	//数据库检查是否存在
	s := db.NewSessionStore()
	defer s.Close()
	if err := s.C("fatalemails").Find(nil).One(&this); err != nil {
		if err.Error() != s.GetErrNotFound().Error() {
			return  err
		}
	}
	return nil;
}
