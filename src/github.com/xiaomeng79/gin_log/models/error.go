/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-25 下午12:10
***********************************************/
package models


var E = map[int]string{
	200 : "ok",

	400 : "参数错误或格式错误",
	401 : "未授权",
	403 : "服务器拒绝该请求",//不是正确的逻辑
	404 : "接口不存在",
	405 : "方法不允许",
	406 : "未定义接收数据的格式",
	408 : "请求超时",
	410 : "资源不可用",
	413 : "请求体过大",
	429 : "服务不可用(接口被限流)",

	500 : "服务器内部错误",
	501 : "尚未实施",//服务器无法识别请求方法时可能会返回此代码
	502 : "错误网关",
	503 : "服务器维护中",
	504 : "网关超时",



}


