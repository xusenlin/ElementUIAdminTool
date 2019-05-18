package buildCode

import (
	"ElementUIAdminTool/config"
	"ElementUIAdminTool/helper"
	"ElementUIAdminTool/templateStr"
	"text/template"
	"bytes"
	"fmt"
	"strings"
)

func Run() {

	switch config.Type {
	case "1":
		buildTable()

	case "2":
		buildForm()

	default:
		fmt.Println("程序中断，请输入正确的-t参数")

	}
}

func buildTable() {
	buildTableVue()
	buildMixin()
	buildCss()
}

func buildForm() {
	buildFormVue()
	buildMixin()
	buildCss()
}

func buildFormVue()  {

	tmpl, err := template.New("").Delims("{%","%}").Parse(templateStr.FormVue)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, map[string]interface{}{
		"Module": config.Module,
		"Dir":    config.Dir,
		"Fields": strings.Split(config.Fields,"|"),
	})

	if err != nil {
		panic(err)
	}

	if helper.WriteFile(config.PageDir+"/Index.vue", buf.Bytes()) {
		fmt.Println("生成" + config.PageDir + "/Index.vue" + "文件成功！")
	}
}

func buildTableVue()  {
	tmpl, err := template.New("").Delims("{%","%}").Parse(templateStr.TableList)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, map[string]interface{}{
		"Module": config.Module,
		"Dir":    config.Dir,
		"Fields": strings.Split(config.Fields,"|"),
	})

	if err != nil {
		panic(err)
	}

	if helper.WriteFile(config.PageDir+"/Index.vue", buf.Bytes()) {
		fmt.Println("生成" + config.PageDir + "/Index.vue" + "文件成功！")
	}
}


func buildMixin()  {
	if helper.WriteFile(config.MixinDir+"/pagination.js", []byte(templateStr.Pagination)) {
		fmt.Println("生成" + config.MixinDir + "/Index.vue" + "文件成功！")
	}
}

func buildCss()  {
	if helper.WriteFile(config.CssDir+"/style.scss", []byte(templateStr.Css)) {
		fmt.Println("生成" + config.CssDir + "/style.scss" + "文件成功！")
	}
}