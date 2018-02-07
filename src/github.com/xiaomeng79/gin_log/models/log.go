/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-25 上午11:16
***********************************************/
package models

import (
	"github.com/xiaomeng79/gin_log/db"
	"gopkg.in/mgo.v2/bson"
	"github.com/asaskevich/govalidator"
	"errors"
	"encoding/json"
	"github.com/xiaomeng79/gin_log/libs"
	"github.com/xiaomeng79/gin_log/conf"
	"github.com/xiaomeng79/gin_log/log"
	"strconv"
)

type LogInfo struct {
	Id bson.ObjectId `bson:"_id" json:"recordId" `
	AppId bson.ObjectId `bson:"appid" json:"-" `
	LogType string `bson:"logtype" json:"logType" `
	Issue string `bson:"issue" json:"issue" `
	RecordType string `bson:"recordtype" json:"recordType"`
	RecordPage int `bson:"recordpage" json:"recordPage" `
	RecordTime string `bson:"recordtime" json:"recordTime"`
	ClientIp string `bson:"clientip" json:"clientIp"`
	Remark string `bson:"remark" json:"remark"`
	DetailParams string `bson:"detailparams" json:"detailParams"`
	RecordStartTime string `bson:"-" json:"recordStartTime,omitempty"`
	RecordEndTime string `bson:"-" json:"recordEndTime,omitempty"`
}

/**
接口日志
 */
type InterfaceLog struct {
	Request string
	Response string
	Url string
	Status int
	Delay string
	ContentType string
	Method string
}

/**
内部日志
 */
type InsideLog struct {
	Path string
	Msg string
	ExceptionType string
	OsInfo string
	OsVer string
	HardInfo string
	ProjectVer string
}

/**
监控日志
 */
type MonitorLog struct {
	Subject string `json:"subject"`
	Body string `json:"body"`
	To string `json:"to"`
}

/**
状态日志
 */
type StateLog struct {
	Msg string
}


/**
设定存放的集合
 */
func (this *LogInfo) getC() (string) {
	//设定集合名称
	var c string
	if this.LogType == "normal" {
		c = "normallog"
	} else {
		c = "errorlog"
	}
	return c
}

/**
验证数据
 */
func (this *LogInfo) validRecordLog() error {
	//公共
	logTypes := []string{"normal","error","waring","fatal"}

	var b bool
	b = govalidator.IsIn(this.LogType,logTypes...)
	if !b {
		return errors.New(this.LogType + "日志类型不存在")
	}

	recordTypes := []string{"interface","inside","monitor","state"}
	b = govalidator.IsIn(this.RecordType,recordTypes...)
	if !b {
		return errors.New(this.RecordType + "日志记录类型不存在")
	}

	return nil
}

func (this *LogInfo) validGetRecord() error {
	//公共
	logTypes := []string{"normal","error","waring","fatal"}

	var b bool
	b = govalidator.IsIn(this.LogType,logTypes...)
	if !b {
		return errors.New(this.LogType + "日志类型不存在")
	}

	if !govalidator.IsNull(this.RecordType) {
		recordTypes := []string{"interface","inside","monitor","state"}
		b = govalidator.IsIn(this.RecordType,recordTypes...)
		if !b {
			return errors.New(this.RecordType + "日志记录类型不存在")
		}
	}


	return nil
}

/**
生成发送邮件的body
 */
func (this *LogInfo) getBody() (string) {
	var s string
	s = "<p style='color:red;'>日志类型:" + this.LogType + "</p>" +
		"<p>日志记录ID:" + this.Id.Hex() + "</p>" +
		"<p>问题描述:" + this.Issue + "</p>" +
		"<p>记录类型:" + this.RecordType + "</p>" +
		"<p>记录位置:" + strconv.Itoa(this.RecordPage) + "</p>" +
		"<p>发生时间:" + libs.MicTimeToStr(this.RecordTime) + "</p>" +
		"<p>客户端IP：" + this.ClientIp + "</p>" +
		"<p>备注:" + this.Remark + "</p>" +
		"<p>详细信息:" + this.DetailParams + "</p>"
	return s
}

/**
处理日志详情
 */
func (this *LogInfo) handleDetail() (error) {
	var err error
	var monitor MonitorLog
	switch this.RecordType {
		case "interface":
		case "inside":
		case "monitor":
			err = json.Unmarshal([]byte(this.DetailParams),&monitor)
			if err != nil {
				log.Release("解析detailParams参数错误：%v",err)
				return errors.New("解析detailParams参数错误")
			}
		case "state":
	}


	switch this.LogType {
		case "normal":
		case "error":
		//取出发邮件的人员
		var appinfo App
		appinfo.Id = this.AppId
		err = appinfo.GetAppInfoById()
			if err != nil {
				log.Release("获取appinfo失败：%v",err)
				return errors.New("获取appinfo失败")
			}
		if !govalidator.IsNull(appinfo.Emails) {
			//发送邮件
			go this.sendEmail(appinfo.Emails,"项目:" + appinfo.ProjectName + "发生错误",this.getBody())
		}
		case "waring":
		//提醒monitor中的人员
			if !govalidator.IsNull(monitor.To) {
				go this.sendEmail(monitor.To, monitor.Subject, monitor.Body)
			}
		case "fatal":
		//提醒全部开发人员
			//取出发邮件的人员
			var appinfo App
			appinfo.Id = this.AppId
			var fatalemails FatalEmails
			c1 := make(chan error,1)
			c2 := make(chan error,1)
			go func(ch chan error){
				c1 <- appinfo.GetAppInfoById()
			}(c1)
			go func(ch chan error){
				c2 <- fatalemails.GetEmails()
			}(c2)

			if err = <-c1; err!= nil {
				log.Release("获取appinfo失败：%v",err)
				return errors.New("获取appinfo失败")
			}
			if err = <-c2; err!= nil {
				log.Release("获取邮箱列表失败：%v",err)
				return errors.New("获取全部邮箱列表失败")
			}
			//获取全部邮箱
			if !govalidator.IsNull(fatalemails.Emails) {
				go this.sendEmail(fatalemails.Emails,"项目:" + appinfo.ProjectName + "发生严重错误",this.getBody())
			}
	}

	return nil
}

/**
记录日志
 */
func (this *LogInfo) RecordLog() (int,error) {
	//数据库检查是否存在
	s := db.NewSessionStore()
	defer s.Close()
	//记录错误
	var err error
	//验证数据
	err = this.validRecordLog()
	if err != nil {
		log.Release("RecordLog验证数据失败:%v",err)
		return 400,err
	}
	//记录
	this.Id = bson.NewObjectId()
	if err := s.C(this.getC()).Insert(&this); err != nil {
		log.Release("RecordLog记录日志失败:%v",err)
		return  500,errors.New("记录失败")
	}
	//处理日志详情,主要发送邮件
	err = this.handleDetail()
	if err != nil {
		log.Release("RecordLog处理日志详情失败:%v",err)
		return 400,err
	}
	return 200,nil;
}


/**
获取日志记录
 */
func (this *LogInfo) GetRecord(reqparm *ReqParam) (int,error,interface{}) {
	//数据库检查是否存在
	s := db.NewSessionStore()
	defer s.Close()
	//记录错误
	var err error
	//验证数据
	err = this.validGetRecord()
	if err != nil {
		log.Release("GetRecord验证数据失败:%v",err)
		return 400,err,nil
	}
	//分页处理

	//生成查询条件
	m := bson.M{}
	m["logtype"] = this.LogType
	if !govalidator.IsNull(this.Id.Hex()) {
		m["_id"] = this.Id
	}
	if !govalidator.IsNull(this.RecordType) {
		m["recordtype"] = this.RecordType
	}
	if this.RecordPage != 0 {
		m["recordpage"] = this.RecordPage
	}
	log.Release("%+v",this.RecordPage)
	if !govalidator.IsNull(this.RecordStartTime) {
		m["recordtime"] = bson.M{"$gte":libs.StrToMicTime(this.RecordStartTime)}
	}
	if !govalidator.IsNull(this.RecordEndTime) {
		m["recordtime"] = bson.M{"$lte":libs.StrToMicTime(this.RecordEndTime)}
	}
	if !govalidator.IsNull(this.RecordStartTime) && !govalidator.IsNull(this.RecordEndTime) {
		m["recordtime"] = bson.M{"$lte":libs.StrToMicTime(this.RecordEndTime),"$gte":libs.StrToMicTime(this.RecordStartTime)}
	}
	m["appid"] = this.AppId
	log.Release("%+v",m)
	//设定页数
	var total int
	total,err = s.C(this.getC()).Find(m).Count();
	if  err != nil {
		log.Release("获取日志失败:%v",err)
		return  500,errors.New("获取日志记录失败"),nil
	}
	reqparm.InitPage(total)
	//记录
	var logs []LogInfo
	if err := s.C(this.getC()).Find(m).Limit(reqparm.Page.PageSize).Skip((reqparm.Page.PageIndex-1) * reqparm.Page.PageSize).All(&logs); err != nil {
		log.Release("获取日志失败:%v",err)
		return  500,errors.New("获取日志记录失败"),nil
	}
	log.Release("%+v",logs)
	return 200,nil,logs;
}

/**
发送邮件
 */
func (this *LogInfo) sendEmail(sendTo,subject,body string) {
	err := libs.SendEmailSSL(conf.EmailHost,conf.EmailPort,conf.EmailPassword,conf.EmailUserName,conf.EmailNickName,sendTo,subject,body)
	if err != nil {
		log.Release("发送邮件失败：%v",err)
	}
}


