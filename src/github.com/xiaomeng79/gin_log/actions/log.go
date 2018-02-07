/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-25 上午11:09
***********************************************/
package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaomeng79/gin_log/models"
	"github.com/xiaomeng79/gin_log/log"
)

/**
添加
 */
func RecordLog(c *gin.Context) {
	var err error
	var loginfo models.LogInfo
	reqpam := c.MustGet("reqdata").(models.ReqParam)
	loginfo.ClientIp = reqpam.ClientIp
	loginfo.AppId = reqpam.AppId
	err = reqpam.DataDecode(&loginfo)
	if err != nil {
		log.Release("记录日志解析参数失败%v",err)
		reqpam.R(c,400,err.Error(),"")
		return
	}
	code,err := loginfo.RecordLog()
	if err != nil {
		log.Release("记录日志失败%v",err)
		reqpam.R(c,code,err.Error(),"")
		return
	}
	reqpam.R(c,code,"","")
	return
}

/**
添加
 */
func GetRecord(c *gin.Context) {
	var err error
	var loginfo models.LogInfo
	reqpam := c.MustGet("reqdata").(models.ReqParam)
	loginfo.AppId = reqpam.AppId
	err = reqpam.DataDecode(&loginfo)
	if err != nil {
		log.Release("获取日志解析参数失败%v",err)
		reqpam.R(c,400,err.Error(),"")
		return
	}
	code,err,logs := loginfo.GetRecord(&reqpam)
	if err != nil {
		log.Release("记录日志失败%v",err)
		reqpam.R(c,code,err.Error(),"")
		return
	}
	reqpam.R(c,code,"",logs)
	return
}

/**
批量添加
 */
func BatRecordLog(c *gin.Context) {
	var err error
	var loginfos []models.LogInfo
	reqpam := c.MustGet("reqdata").(models.ReqParam)
	err = reqpam.DataDecode(&loginfos)
	if err != nil {
		log.Release("批量获取日志解析参数失败%v",err)
		reqpam.R(c,400,err.Error(),"")
		return
	}
	for _,v := range loginfos {
		v.ClientIp = reqpam.ClientIp
		v.AppId = reqpam.AppId
		code,err := v.RecordLog()
		if err != nil {
			log.Release("批量记录日志失败%v",err)
			reqpam.R(c,code,"部分日志记录失败：" + err.Error(),"")

		}
	}
	reqpam.R(c,200,"","")
	return
}
