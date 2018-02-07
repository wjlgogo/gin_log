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
func SetProjectEmails(c *gin.Context) {
	var err error
	var app models.App
	reqpam := c.MustGet("reqdata").(models.ReqParam)
	app.AppKey = reqpam.AppKey
	err = reqpam.DataDecode(&app)
	err = app.SetProjectEmails()
	if err != nil {
		log.Release("设置app信息失败:%v",err)
		reqpam.R(c,500,err.Error(),"")
		return
	}
	err =app.GetApp()
	if err != nil {
		log.Release("获取app信息失败:%v",err)
		reqpam.R(c,500,err.Error(),"")
		return
	}
	reqpam.R(c,200,"",app)
	return

}
