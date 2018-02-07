/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-24 上午11:00
***********************************************/
package libs

import (
	"os"
	"time"
	"fmt"
	"path"
	"crypto/md5"
	"strconv"
	"github.com/xiaomeng79/gin_log/log"
)

/**
判断文件是否存在 存在 true,nil
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
创建一个时间格式的文件
 */
func CreateTimeFile(pathname,ext string) (*os.File, error) {
	now := time.Now()
	md5.New()

	filename := fmt.Sprintf("%d%02d%02d_%02d_%02d_%02d" + "." +ext,
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second())

	return os.Create(path.Join(pathname, filename))
}

/**
生成毫秒时间戳
 */
func GenMicTime() string {
	return strconv.FormatInt(time.Now().UnixNano()/1e6,10)
}

/**
毫秒时间戳转字符串
 */
func MicTimeToStr(s string) string {
	i64, err := strconv.ParseInt(s,10,64)
	if err != nil {
		log.Release("转换时间出错:%v",err)
		return ""
	}
	tm := time.Unix(i64/1e3,0)
	return tm.Format("2006-01-02 15:04:05")
}

/**
字符串转毫秒时间戳
 */
func StrToMicTime(s string) string {
	//获取本地location   	//待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, s, loc) //使用模板在对应时区转化为time.time类型
	return strconv.FormatInt(theTime.Unix() * 1e3, 10)
}


func ErrorLog(err error) {
	if err != nil {
		log.Release("%v",err)
	}
}
