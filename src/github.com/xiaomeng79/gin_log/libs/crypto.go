/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-26 下午4:10
***********************************************/
package libs

import (
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

/**
MD5加密
 */
func MD5(s string) string {
	r := md5.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

/**
SHA1加密
 */
func SHA1(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

/**
SHA256加密
 */
func SHA256(s string) string {
	r := sha256.Sum256([]byte(s))
	return hex.EncodeToString(r[:])
}

/**
SHA512加密
 */
func SHA512(s string) string {
	r := sha512.Sum512([]byte(s))
	return hex.EncodeToString(r[:])
}
