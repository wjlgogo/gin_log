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

type App struct {
	Id bson.ObjectId `bson:"_id" json:"-" `
	AppKey string `bson:"appkey" json:"-" `
	AppSecret string `bson:"appsecret" json:"-" `
	ProjectName string `bson:"projectname" json:"projectname"`
	IsUsed int `bson:"issued" json:"-" `
	Emails string `bson:"emails" json:"emails"`
	Tels string `bson:"tels" json:"tels"`
	Remark string `bson:"remark" json:"remark"`
}


/**
获取APP信息
 */
func (this *App) GetApp() error {
	//数据库检查是否存在
	s := db.NewSessionStore()
	defer s.Close()
	if err := s.C("appinfo").Find(bson.M{"appkey":this.AppKey}).One(&this); err != nil {
		if err.Error() != s.GetErrNotFound().Error() {
			return  err
		}
	}

	return nil;

}

/**
获取App信息
 */
func (this *App) GetAppInfoById() error {
	//数据库检查是否存在
	s := db.NewSessionStore()
	defer s.Close()
	return s.C("appinfo").FindId(this.Id).One(&this)

}


/**
改变app信息
 */
func (this *App) SetProjectEmails() error {
	//数据库检查是否存在
	s := db.NewSessionStore()
	defer s.Close()
	return s.C("appinfo").Update(bson.M{"appkey":this.AppKey},bson.M{
		"$set":bson.M{
			"emails":this.Emails,
			"tels":this.Tels,
			"remark":this.Remark,
		}});
}

/**
创建APP
 */
func (this *App) SetApp() error {
	this.Id = bson.NewObjectId()
	this.AppSecret = "3456"
	this.ProjectName = "testproject"
	this.IsUsed = 1
	//数据库检查是否存在
	s := db.NewSessionStore()
	defer s.Close()
	if err := s.C("appinfo").Insert(&this); err != nil {
		return  err
	}

	return nil;

}