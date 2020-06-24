package TplParser

import (
	"github.com/phpwei/go-Gentity/Helper"
	"github.com/phpwei/go-Gentity/resource"
	"log"
	"os"
)

type ConfigParser struct {
	TplContent string  //模板内容
}
//初始化 并读取模板内容
func NewConfigParser() *ConfigParser  {
	return &ConfigParser{TplContent:Helper.UnGzip(resource.CONF_TPL)}
}
func(this *ConfigParser) Parse(target string){

	file,err:=os.OpenFile(target,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)


	var d1 = []byte(this.TplContent)
	_,err=file.Write(d1)

	if err!=nil{
		log.Fatal("generate model error:",err)
	}
}


