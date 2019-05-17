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
	ViewDir string
	ModuleDir string
	PageDir string
	CssDir string
	ImgDir string
	MixinDir string
)

func init()  {
	flag.StringVar(&Module, "m", "", "views下的目录,不存在会新建。代表一个功能模块(Module)")
	flag.StringVar(&Dir, "d", "", "参数m目录下的目录,存在不会覆盖。代表一个页面目录(Dir)")
	flag.StringVar(&Type, "t", "1", "默认值(1)。1生成table列表，2生成form(Type)")
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

	cssErr := os.Mkdir(CssDir, os.ModePerm)
	if cssErr != nil {
		fmt.Println("生成"+ CssDir +"目录失败！")
		return false
	}
	fmt.Println("生成"+ CssDir +"目录成功！")

	imgErr := os.Mkdir(ImgDir, os.ModePerm)
	if imgErr != nil {
		fmt.Println("生成"+ ImgDir +"目录失败！")
		return false
	}
	fmt.Println("生成"+ ImgDir +"目录成功！")

	mixinErr := os.Mkdir(MixinDir, os.ModePerm)
	if mixinErr != nil {
		fmt.Println("生成"+ MixinDir +"目录失败！")
		return false
	}
	fmt.Println("生成"+ MixinDir +"目录成功！")

	return true
}