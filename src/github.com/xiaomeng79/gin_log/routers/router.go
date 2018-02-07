/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-24 上午9:57
***********************************************/
package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaomeng79/gin_log/actions"
	"github.com/xiaomeng79/gin_log/middlewares"
	"github.com/xiaomeng79/gin_log/models"
)

func Reg() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.Use(middlewares.VerifyData())
	v1.Use(middlewares.Cors())//解决跨域
	/*************************Person****************************/
	v1.POST("/setProjectEmails",actions.SetProjectEmails)//设置项目信息
	v1.POST("/recordLog",actions.RecordLog)//记录日志
	v1.POST("/batRecordLog",actions.BatRecordLog)//记录日志
	v1.POST("/getRecord",actions.GetRecord)//查询日志



	/*************************不存在**************************/
	r.NoMethod(func(c *gin.Context) {
		reqpam := c.MustGet("reqdata").(models.ReqParam)
		reqpam.R(c,405,"","")
	})

	r.NoRoute(func(c *gin.Context) {
		reqpam := c.MustGet("reqdata").(models.ReqParam)
		reqpam.R(c,404,"","")
	})

	return r
}

