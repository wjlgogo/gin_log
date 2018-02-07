/**********************************************
** @Des: This file ...
** @Author: xiaomeng79
** @Date:   18-1-24 下午1:44
***********************************************/
package libs

import (
	"github.com/xiaomeng79/gin_log/libs"
	"fmt"
)

func ExamplePathExists() {
	fmt.Println(libs.PathExists("example.file"))
	fmt.Println(libs.PathExists("example.json"))

	//Output:
	//true <nil>
	//false <nil>
}

