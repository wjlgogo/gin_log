/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-24 下午1:54
***********************************************/
package libs

import (
	"testing"
	"github.com/xiaomeng79/gin_log/libs"
)

func TestPathExists(t *testing.T) {
	var tests = []struct {
		in string
		expected bool
		err error
	}{
		{"example.file",true,nil},
		{"example.json",false,nil},
		{"example/example.file",false,nil},
	}


	for _, tt := range tests {
		_expected, _err := libs.PathExists(tt.in)

		if _expected != tt.expected || _err != tt.err {
			t.Errorf("libs.PathExists(%s) = %t,%v expected %t,%v",tt.in,_expected,_err,tt.expected,tt.err)
		}
	}
}

func TestMD5(t *testing.T) {
	var tests = []struct {
		in string
		expected string
	}{
		{"root_","cf9ae78b63bc0687d57fa51456af2184"},
		{"123%_ttt&","6aa8e711a97499e7407fe989ec2fde58"},
		{"--ddddd333%#&/","a5169b77d7f30e61b51b8d58bf6c0df2"},
	}

	for _, tt := range tests {
		_expected := libs.MD5(tt.in)

		if _expected != tt.expected {
			t.Errorf("libs.MD5(%s) = %v expected %s",tt.in,_expected,tt.expected)
		}
	}
}

func TestSHA1(t *testing.T) {
	var tests = []struct {
		in string
		expected string
	}{
		{"root_","c484a1cb5f9827891169204fd8fb13ce865a77b6"},
		{"123%_ttt&","6f028cd2abd3e36779c2f1336a86e5f360a2dbd2"},
		{"--ddddd333%#&/","e849cc0334c3b9eb242197a53f96bc7579056b3e"},
	}

	for _, tt := range tests {
		_expected := libs.SHA1(tt.in)

		if _expected != tt.expected {
			t.Errorf("libs.SHA1(%s) = %v expected %s",tt.in,_expected,tt.expected)
		}
	}
}

func TestSHA256(t *testing.T) {
	var tests = []struct {
		in string
		expected string
	}{
		{"root_","fb2d43e7f4e396cb24ececc8c582d79948a2b1f3a864a596ce2381fc4f215b24"},
		{"123%_ttt&","5b8fdf5c1fc837034feb96a4ae23f59535abd34bb6c202e1fcfe96279419a8e2"},
		{"--ddddd333%#&/","2ecfec1f1beac8c7e1678e7c568b0d3796029ca0554478f3dba5969a9d479c0c"},
	}

	for _, tt := range tests {
		_expected := libs.SHA256(tt.in)

		if _expected != tt.expected {
			t.Errorf("libs.SHA256(%s) = %v expected %s",tt.in,_expected,tt.expected)
		}
	}
}

func TestSHA512(t *testing.T) {
	var tests = []struct {
		in string
		expected string
	}{
		{"root_","76daa5da50d40ab38a937c9ca0a8d30f6a6adb4cc861d8a5238f5bc2e2478db2aeaa178e30dc9718e0d2b45f8d2ee0fa8fa55f0d5ad13526bd2e3daae703a5cc"},
		{"123%_ttt&","3411cd76712d030330c7ef9d4b2a9b0f34f9cb30a1815f647a849ecb1732c7ccea4b66e379744978d98a23c6d4354e1e95fd7149ba599f8980ab4c92e237b133"},
		{"--ddddd333%#&/","b4a77076831790d59b6e17d7b967bc34c6d39d296d7e5a22a0c9a3558d41e82c2ef30b0f389d1800a04c6660299395fe1390cf3c6280630d96e1815c50bbf614"},
	}

	for _, tt := range tests {
		_expected := libs.SHA512(tt.in)

		if _expected != tt.expected {
			t.Errorf("libs.SHA512(%s) = %v expected %s",tt.in,_expected,tt.expected)
		}
	}
}