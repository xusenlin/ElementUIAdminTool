package config

import (
	"ElementUIAdminTool/helper"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	Module string
	Dir string
	Type string
	Fields string
	ViewDir string
	ModuleDir string
	PageDir string
	CssDir string
	ImgDir string
	MixinDir string
)

func init()  {
	flag.StringVar(&Module, "m", "", "Module必填，views下的目录,不存在会新建。代表一个功能模块")
	flag.StringVar(&Dir, "d", "", "Dir必填，参数m目录下的目录,存在不会覆盖。代表一个页面目录")
	flag.StringVar(&Type, "t", "1", "Type默认值(\"1\")，1生成table列表，2生成form")
	flag.StringVar(&Fields, "f", "ID|标题", "Field默认值(\"ID|标题\")，指定table或者form的字段，用|隔开,表格自动添加操作列，如-f=\"ID|标题|状态|时间\"")
}

func InitConfigAndDir() bool {
	flag.Parse()

	currentDir, pwdErr := os.Getwd()
	if pwdErr != nil {
		fmt.Println("获取当前目录出错：",pwdErr)
		return false
	}

	if ! strings.Contains(currentDir, "utils") {
		fmt.Println("请将此工具放置在你的项目->src->utils目录下")
		return false
	}

	if len(Module) == 0 || len(Dir) == 0{
		fmt.Println("缺少参数，请使用 -h查看帮助！")
		return false
	}

	ViewDir = "../views"
	if ! helper.IsDir(ViewDir){
		fmt.Println("缺少views目录")
		return false
	}

	ModuleDir = ViewDir + "/" + Module

	if ! helper.IsDir(ModuleDir){

		if ! helper.Mkdir(ModuleDir){
			return false
		}
		fmt.Println("生成"+ ModuleDir +"目录！")
	}

	PageDir = ModuleDir + "/" + Dir

	if  helper.IsDir(PageDir){
		fmt.Println( PageDir + "目录已经存在")
		return false
	}else {
		if ! helper.Mkdir(PageDir){
			return false
		}
		fmt.Println("生成"+ PageDir +"目录！")
	}

	if ! buildPageSubDir(){
		return false
	}

	return true
}



func buildPageSubDir() bool {

	CssDir = PageDir + "/css"
	ImgDir = PageDir + "/img"
	MixinDir = PageDir + "/mixin"


	if helper.Mkdir(CssDir) && helper.Mkdir(ImgDir) && helper.Mkdir(MixinDir){
		return true
	}

	fmt.Println("生成" + CssDir + "、" + ImgDir + "、" + MixinDir + "、" + "目录失败！")

	return false

}