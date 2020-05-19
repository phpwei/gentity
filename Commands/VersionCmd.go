package Commands

import (
	"flag"
	"fmt"
)

func init()  {
	vc:=NewVersionCmd()
	vc.Register("VersionCommand",vc) //注册命令
}

type VersionCmd struct {
	FlagValue *bool
	CommandSet
}
func NewVersionCmd() *VersionCmd{
	return &VersionCmd{}
}
func(this *VersionCmd) Init()  {
	this.FlagValue=flag.Bool("v",false,"show version")
}
func(this *VersionCmd) Run()  {
	if *this.FlagValue{
		fmt.Printf("Version: %s\n", "0.0.1")
	}
}