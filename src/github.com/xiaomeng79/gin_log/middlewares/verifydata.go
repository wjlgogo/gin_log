/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-26 下午2:38
***********************************************/
package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaomeng79/gin_log/models"
	"github.com/xiaomeng79/gin_log/log"
	"encoding/json"
)

/**
解析参数
 */
func VerifyData() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		//1.解析参数
		var reqpam models.ReqParam
		err = c.ShouldBindJSON(&reqpam)
		if err != nil {
			reqpam.HandleErr(c,400,"解析数据错误")
			log.Release("解析数据错误:%v",err)
			return
		}
		//绑定客户端IP
		reqpam.ClientIp = c.Request.RemoteAddr
		//记录数据
		_d,_ := json.Marshal(reqpam)
		//reqpam.SetApp()
		//验证数据
		c1 := make(chan error,1)
		c2 := make(chan error,1)
		go func(ch chan error){
			c1 <- reqpam.CheckRequestId()//查看RequestId是否存在
		}(c1)

		go func(ch chan error){
			c2 <- reqpam.SetAppSecret() //获取密钥
		}(c2)

		if err = <-c1; err!= nil {
			reqpam.HandleErr(c,403,err.Error())
			log.Release("检查requestId错误:%v",err)
			return
		}
		if err = <-c2; err != nil {
			reqpam.HandleErr(c,401,err.Error())
			log.Release("设置appSecret错误:%v",err)
			return
		}

		//2.验证签名
		b,err := reqpam.CompareSign()
		if err != nil {
			reqpam.HandleErr(c,400,err.Error())
			log.Release("验证签名错误:%v",err)
			return
		}

		if !b {
			reqpam.HandleErr(c,400,"签名不正确")
			log.Release("验证签名错误:%v",err)
			return
		}

		c.Set("reqdata", reqpam)
		//c.Abort() //拒绝
		c.Next() //通过

		//记录请求参数原始数据
		var rinfo models.RequestInfo
		rinfo.RequestId = reqpam.RequestId
		rinfo.Data = string(_d)
		err = rinfo.Add()
		if err != nil {
			log.Release("%v",err)
		}
	}
}




