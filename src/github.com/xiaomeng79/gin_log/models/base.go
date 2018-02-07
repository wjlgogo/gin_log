/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-25 下午2:24
***********************************************/
package models

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"errors"
	"encoding/base64"
	"strings"
	"github.com/xiaomeng79/gin_log/libs"
	"github.com/asaskevich/govalidator"
	"github.com/xiaomeng79/gin_log/log"
	"gopkg.in/mgo.v2/bson"
	"math"
)


type ReqParam struct {
	AppId bson.ObjectId  `json:"-"`//AppID
	AppKey string  `json:"appKey" binding:"required"`//密钥ID
	AppSecret string `json:"appSecret"` //密钥
	ClientIp string `form:"clientip" json:"clientIp"`//客户端IP
	RequestId string `json:"requestId" binding:"required"` //32位的唯一请求标识，用于问题排查和防止重复提交
	Timestamp string `json:"timestamp" binding:"required"` //毫秒时间戳
	Sign string `json:"sign" binding:"required"` //签名
	SignType string `json:"signType" binding:"required"` //签名类型：MD5 SHA1 SHA256 SHA512
	Encode bool `json:"encode"` //响应数据data是否进行base64编码，默认true
	Data string `json:"data" binding:"required"` //请求的数据
	Page Page `json:"page"` //分页
	IsPage bool `json:"isPage"` //是否分页
}

//分页
type Page struct {
	PageIndex int `json:"pageIndex"` //页面索引
	PageSize int `json:"pageSize"` //每页大小
	PageTotal int `json:"pageTotal"` //总分页数
	Count int `json:"count"` //当页记录数
	Total int `json:"total"` //总记录数
}

/**
初始化分页
 */
func (r *ReqParam) InitPage(total int) {
	r.IsPage = true
	if r.Page.PageIndex <= 0 {
		r.Page.PageIndex = 1
	}
	if r.Page.PageSize <= 0 {
		r.Page.PageSize = 20
	}
	//最大的索引
	maxIndex := int(math.Ceil(float64(total)/float64(r.Page.PageSize)))
	if maxIndex != 0 && maxIndex < r.Page.PageIndex {
		r.Page.PageIndex = maxIndex
	}
	//总记录数
	r.Page.Total = total
	//总分数页
	if maxIndex <= 0 {//没数据情况
		r.Page.PageTotal = 0
		r.Page.Count = 0
	} else {
		r.Page.PageTotal = maxIndex
	}
	//当页记录数
	if r.Page.PageTotal > r.Page.PageIndex {//总页数大于当前页数
		r.Page.Count = r.Page.PageSize
	}
	if r.Page.PageTotal == r.Page.PageIndex {//总页数等于当前页数
		r.Page.Count = r.Page.Total - (r.Page.PageSize * (r.Page.PageIndex-1))
	}

}



/**
检查RequestID
 */
func (r *ReqParam) CheckRequestId() error {
	if !govalidator.IsUUID(r.RequestId) {
		return errors.New("requestId不是有效的UUID")
	}
	//数据库检查是否存在
	var requestinfo RequestInfo
	requestinfo.RequestId = r.RequestId
	return requestinfo.CheckRequestId()
}

/**
设置客户端IP
 */
func (r * ReqParam) SetClientIp() error {
	if !govalidator.IsUUID(r.RequestId) {
		return errors.New("requestId不是有效的UUID")
	}
	//数据库检查是否存在
	var requestinfo RequestInfo
	requestinfo.RequestId = r.RequestId
	return requestinfo.CheckRequestId()
}

/**
检查appkey和设置appsecret
 */
func (r *ReqParam) SetAppSecret() error {

	if govalidator.IsNull(r.AppKey) {
		return errors.New("appKey不能为空")
	}
	var app App
	app.AppKey = r.AppKey
	err := app.GetApp()
	if err != nil {
		return errors.New("appKey不存在")
	}
	if app.IsUsed == 0 {
		return errors.New("appKey禁止使用")
	}
	r.AppSecret = app.AppSecret
	r.AppId = app.Id

	return nil
}

func (r *ReqParam) SetApp() error {

	if govalidator.IsNull(r.AppKey) {
		return errors.New("appKey不能为空")
	}
	var app App
	app.AppKey = r.AppKey
	app.SetApp()

	return nil
}



/**
**解析data参数
**input: v point
**ouput:  error
 */
func (r *ReqParam) DataDecode (v interface{}) error {
	if r.Data == "" {
		return nil
	}
	//判断参数是否base64编码
	var decoded []byte
	var err error
	if r.Encode {//解码
		decoded, err = base64.StdEncoding.DecodeString(r.Data)
		if err != nil {
			return errors.New("base64解码失败")
		}
	} else {
		decoded = []byte(r.Data)
	}
	err =json.Unmarshal(decoded,v)
	if err != nil {
		return errors.New("data参数json解析失败")
	}

	return nil
}


/**
**编码data参数
**input: v interface
**ouput:  error
 */
func (r *ReqParam) DataEncode (v interface{}) error {
	if v == "" {
		r.Data = ""
		return nil
	}
	var encoded []byte
	var err error
	//json marshal
	encoded,err = json.Marshal(v)
	if err != nil {
		return errors.New("json编码失败")
	}

	//判断参数是否base64编码
	if r.Encode {//解码
		r.Data = base64.StdEncoding.EncodeToString(encoded)
	} else {
		r.Data = string(encoded)
	}

	return nil
}

/**
生成签名 string
 */
func (r *ReqParam) CreateSign() (string,error) {
	_signType := strings.ToUpper(strings.Trim(r.SignType," "))
	var originSign string
	//组合签名字符串
	if r.Data == "" {
		originSign = r.AppKey + r.AppSecret + r.RequestId + r.Timestamp
	} else {
		originSign = r.AppKey + r.AppSecret +r.Data + r.RequestId + r.Timestamp
	}
	var _sign string
	var err error
	switch _signType {
	case "MD5":
		_sign = libs.MD5(originSign)
	case "SHA1":
		_sign = libs.SHA1(originSign)
	case "SHA256":
		_sign = libs.SHA256(originSign)
	case "SHA512":
		_sign = libs.SHA512(originSign)
	default :
		err = errors.New("签名类型不存在")
	}
	if err != nil {
		return "", err
	}
	return _sign, nil
}

/**
根据签名类型比较签名 true:相同 false：不同
 */
func (r *ReqParam) CompareSign() (bool, error){
	_sign,err := r.CreateSign()
	if err != nil {
		return false,err
	}
	if strings.ToLower(r.Sign) == _sign {
		return true, nil
	} else {
		return false,errors.New("签名不匹配")
	}
}


/**
生成返回数据
 */
func (r *ReqParam) genData(errcode int,errmsg string,v interface{}) (interface{}, error) {

	var err error
	err = r.DataEncode(v)
	if err != nil {
		log.Release("%v",err)
		return "",err
	}
	r.Timestamp = libs.GenMicTime()
	r.Sign,err = r.CreateSign()
	if err != nil {
		log.Release("%v",err)
		return "", err
	}
	if r.IsPage {
		return gin.H{
			"code":errcode,
			"message":E[errcode] + "(" +errmsg + ")",
			"appKey":r.AppKey,
			"requestId":r.RequestId,
			"timestamp":r.Timestamp,
			"sign":r.Sign,
			"signType":r.SignType,
			"encode":r.Encode,
			"data":r.Data,
			"page":r.Page,
		}, nil
	} else {
		return gin.H{
			"code":errcode,
			"message":E[errcode] + "(" +errmsg + ")",
			"appKey":r.AppKey,
			"requestId":r.RequestId,
			"timestamp":r.Timestamp,
			"sign":r.Sign,
			"signType":r.SignType,
			"encode":r.Encode,
			"data":r.Data,
		}, nil
	}
}

//返回数据
/*
0.生成data数据和page数据
1.encode data数据
2.生成时间错和签名
3.生成错误吗和错误信息
4.返回数据
 */
func (r *ReqParam) R (c *gin.Context,errcode int,errmsg string,v interface{}) {
	var err error
	var d interface{}
	d,err =r.genData(errcode,errmsg,v)
	if err != nil {
		log.Release("%v",err)
		r.HandleErr(c,500,err.Error())
		return
	}
	c.JSON(errcode,d)
	return
}


/**
内部错误处理
 */
func (r *ReqParam) HandleErr (c *gin.Context,errcode int,errmsg string,) {

	c.AbortWithStatusJSON(errcode,gin.H{
		"code":errcode,
		"message":E[errcode] + "(" +errmsg + ")",
	})

	return
}
