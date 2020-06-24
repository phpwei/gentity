package Commands

import (
	"flag"
	"fmt"
	"github.com/phpwei/go-Gentity/Helper"
	"github.com/phpwei/go-Gentity/TplParser"
	"log"
	"os"
)

func init()  {
	vc:=NewConfigCmd()
	vc.Register("ConfigCommand",vc) //注册命令
}

type ConfigCmd struct {
	FlagValue *bool
	CommandSet
	ServiceCommandSet *flag.FlagSet
}
func NewConfigCmd() *ConfigCmd{
	var flagValue bool
	flagValue=false
	return &ConfigCmd{FlagValue: &flagValue,
		ServiceCommandSet: flag.NewFlagSet("conf args", flag.ExitOnError)}

}
func(this *ConfigCmd) Init()  {
	if len(os.Args) > 2 && os.Args[1] == "conf" {
		this.FlagValue = this.ServiceCommandSet.Bool("c", false, "create application.yaml")
		err := this.ServiceCommandSet.Parse(os.Args[2:])
		if err != nil {
			log.Println(err)
		}
	}
}
func(this *ConfigCmd) Run()  {
	if *this.FlagValue{
		fmt.Printf("创建config")
		TplParser.NewConfigParser().Parse(Helper.WorkDir+Helper.SYS_CONFIG_PATH)

	}
}