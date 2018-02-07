/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-25 上午11:16
***********************************************/
package models

import (
	"github.com/xiaomeng79/gin_log/db"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"github.com/xiaomeng79/gin_log/libs"
)

type RequestInfo struct {
	Id bson.ObjectId `bson:"_id"`
	RequestId string `bson:"requestid"`
	Data string `bson:"data"`
	CreatedAt string `bson:"createdat"`
}

/**
检查RequestID
 */
func (this *RequestInfo) CheckRequestId() error {
	//数据库检查是否存在
	s := db.NewSessionStore()
	defer s.Close()
	n, err := s.C("requestinfo").Find(bson.M{"requestid":this.RequestId}).Count()
	if err != nil {
		return errors.New("查询requestId失败")
	}

	if n > 0 {
		return errors.New("requestId已经存在，数据重复")
	}
	return nil
}

/**
添加
 */

func (this *RequestInfo) Add() error {
	this.Id = bson.NewObjectId()
	this.CreatedAt = libs.GenMicTime()
	//数据库检查是否存在
	s := db.NewSessionStore()
	defer s.Close()
	if err := s.C("requestinfo").Insert(&this); err != nil {
		return  err
	}
	return nil;

}