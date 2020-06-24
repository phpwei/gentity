package TplParser

import (
	"fmt"
	"github.com/phpwei/go-Gentity/Helper"
	"github.com/phpwei/go-Gentity/resource"
	"log"
	"os"
	"text/template"
)

type DBModelParser struct {
	TplContent string  //模板内容
}
//初始化 并读取模板内容
func NewDBModelParser() *DBModelParser  {
		return &DBModelParser{TplContent:Helper.UnGzip(resource.DB_TPL)}
}
func(this *DBModelParser) Parse(data interface{},target string){
	config, err := Helper.LoadConfig()
	if err != nil {
		log.Fatal("missing config-file:application.yaml")
	}
	tpl:=template.New("DBModel").Funcs(Helper.NewTplFunction())
	tmpl, err:= tpl.Parse(this.TplContent)
	target = Helper.Ucfirst(Helper.CamelCase(target))
	fmt.Printf(target)
	if err!=nil{
		log.Fatal("db tpl parse-error:",err)
	}
	_,err =os.Stat(Helper.WorkDir+config.DB.Path)
	if err != nil {
		//文件夹不存在则创建文件夹
		_ =createFile(Helper.WorkDir+"config.DB.Path")
	}
	file,err:=os.OpenFile(Helper.WorkDir+config.DB.Path+target+".go",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
	if err!=nil{
		log.Fatal("model target error:",err)
	}
	err=tmpl.Execute(file,data)
	if err!=nil{
		log.Fatal("generate model error:",err)
	}
}


func createFile(filePath string)  error  {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath,os.ModePerm)
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true

}