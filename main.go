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
		tmplStr, err := jade.Parse("tmplStr", content)
		// file, err := os.OpenFile("tmpl.html", os.O_APPEND, 0644)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// defer file.Close()
		// len, err := file.WriteString(tmplStr)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Printf("****%s*****%d\n", tmplStr, len)
		if err != nil {
			return nil, fmt.Errorf("error loading jade template: %v", err)
		}

		// 将template string转化为template对象
		tmpl := template.New("jade template")
		tmpl, err = tmpl.Parse(tmplStr)
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
