package main

import (
	"ElementUIAdminTool/buildCode"
	"ElementUIAdminTool/config"
	"fmt"
)

func main()  {

	if ! config.InitConfigAndDir(){
		fmt.Println("初始化配置和目录失败！")
		return
	}

	buildCode.Run()
}
