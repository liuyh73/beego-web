package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/Joker/jade"

	"github.com/astaxie/beego"
	"github.com/liuyh73/cloudgo/models"
	_ "github.com/liuyh73/cloudgo/routers"
	"github.com/liuyh73/cloudgo/utils"
)

func main() {
	addJadeTemplate()
	beego.Run()
}

func addJadeTemplate() {
	beego.AddTemplateEngine("jade", func(root, path string, funcs template.FuncMap) (*template.Template, error) {
		jadePath := filepath.Join(root, path)
		content, err := utils.ReadFile(jadePath)
		if err != nil {
			return nil, fmt.Errorf("error loading jade template: %v", err)
		}

		// 把jade文件读取出来的[]byte转化为template string
		fmt.Println(string(content))
		tmplStr, err := jade.Parse("tmplStr", content)
		// tmplStr, err := jade.ParseFile(jadePath)
		fmt.Printf("%s\n", tmplStr)
		if err != nil {
			return nil, fmt.Errorf("error loading jade template: %v", err)
		}

		// 将template string转化为template对象
		tmpl := template.New("jade template")
		tmpl, err = tmpl.Parse(tmplStr)
		fmt.Println(err)
		if err != nil {
			return nil, fmt.Errorf("error loading jade template: %v", err)
		}
		return tmpl, err
	})
}

func useTemplate() {
	tmpl, _ := template.New("template").Parse("hello {{.Username}}!")
	p := models.User{Username: "jjz"}
	tmpl.Execute(os.Stdout, p)
}
